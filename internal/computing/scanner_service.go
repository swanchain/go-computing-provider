package computing

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/filswan/go-mcs-sdk/mcs/api/common/logs"
	"github.com/swanchain/go-computing-provider/conf"
	"github.com/swanchain/go-computing-provider/internal/contract"
	"github.com/swanchain/go-computing-provider/internal/contract/ecp"
	"github.com/swanchain/go-computing-provider/internal/contract/fcp"
	"github.com/swanchain/go-computing-provider/internal/models"
	"strings"
	"time"
)

type TaskManagerContract struct {
	cpAccount  string
	contract   *fcp.FcpTaskManager
	job        *models.JobEntity
	count      int64
	ethClient  *ethclient.Client
	sigMethods map[string]abi.Method
}

func NewTaskManagerContract() (*TaskManagerContract, error) {
	var tmc TaskManagerContract
	cpAccountAddress, err := contract.GetCpAccountAddress()
	if err != nil {
		logs.GetLogger().Errorf("get cp account contract address failed, error: %v", err)
		return nil, err
	}
	tmc.cpAccount = cpAccountAddress

	chainUrl, err := conf.GetRpcByNetWorkName()
	if err != nil {
		logs.GetLogger().Errorf("failed to get rpc url, error: %v", err)
		return nil, err
	}
	client, err := contract.GetEthClient(chainUrl)
	if err != nil {
		logs.GetLogger().Errorf("failed to dial rpc, error: %v", err)
		return nil, err
	}
	tmc.ethClient = client
	defer tmc.ethClient.Close()

	tmc.contract, err = fcp.NewFcpTaskManager(common.HexToAddress(conf.GetConfig().CONTRACT.JobManager), client)
	if err != nil {
		logs.GetLogger().Errorf("failed to create job manager contract client, error: %+v", err)
		return nil, err
	}

	tmc.sigMethods, err = ContractSigMethods(fcp.FcpTaskManagerABI)
	if err != nil {
		return nil, err
	}
	return &tmc, nil
}

func (taskManager *TaskManagerContract) Scan() error {
	var endBlockNumber uint64
	var end uint64
	var err error

	scannedBlock := loadLastProcessedBlock(models.ScannerFcpTaskManagerId)
	start := uint64(scannedBlock)
	if scannedBlock != 0 {
		start = uint64(scannedBlock) + 1
	}

	endBlockNumber, err = taskManager.ethClient.BlockNumber(context.Background())
	if err != nil {
		return err
	}

	var step uint64 = 1000
	var count int
	for i := start; i <= endBlockNumber; i = i + step {
		count++
		taskManager.count += 1
		end = i + step - 1
		if end > endBlockNumber {
			end = endBlockNumber
		}
		filterOps := &bind.FilterOpts{
			Start: i,
			End:   &end,
		}
		time.Sleep(3 * time.Second)
		if err := taskManager.scanTaskRewards(filterOps); err != nil {
			logs.GetLogger().Errorf("debug_rpc_chain: count: %d, start: %d, end: %d, error: %s", count, i, end, ecp.ParseTooManyError(err))
			return err
		}
		logs.GetLogger().Infof("debug_rpc_chain:count: %d, start: %d, end: %d", count, i, end)
		saveLastProcessedBlock(int64(end), models.ScannerFcpTaskManagerId)
	}
	return nil
}

func (taskManager *TaskManagerContract) scanTaskRewards(opts *bind.FilterOpts) (err error) {
	iterator, err := taskManager.contract.FilterRewardReleased(opts, nil, []common.Address{common.HexToAddress(taskManager.cpAccount)})
	if err != nil {
		return err
	}
	defer iterator.Close()
	for iterator.Next() {
		event := iterator.Event
		if event != nil {
			if err = taskManager.parseRewardReleased(event); err != nil {
				logs.GetLogger().Error(err)
			}
		}
	}
	return nil
}

func (taskManager *TaskManagerContract) parseRewardReleased(event *fcp.FcpTaskManagerRewardReleased) error {
	raw := event.Raw
	amount := contract.BalanceToStr(event.RewardAmount)

	var taskUUIDs []string
	paras, err := TransactionInputParas(taskManager.ethClient, taskManager.sigMethods, raw.TxHash)
	if err != nil {
		logs.GetLogger().Error(err)
		// force parse tx data for old versions
		tx, _, err := taskManager.ethClient.TransactionByHash(context.Background(), event.Raw.TxHash)
		if err != nil {
			return err
		}

		data := tx.Data()
		if len(data) < 36 {
			logs.GetLogger().Errorf("invalid transaction data size: %d", len(data))
			return fmt.Errorf("invalid transaction data size")
		}
		index := -1
		for i := len(data) - 1; i >= 0; i-- {
			if data[i] != 0 {
				index = i
				break
			}
		}
		if index < 0 {
			logs.GetLogger().Errorf("invalid transaction data: %s", string(data))
			return fmt.Errorf("invalid transaction data")
		}

		txData := data[:index+1]
		taskUUIDs = append(taskUUIDs, string(txData[len(txData)-36:]))
	} else {
		logs.GetLogger().Infof("tx %s parsed input paras %v", raw.TxHash.Hex(), paras)

		if len(paras) == 0 {
			return errors.New("invalid parameters")
		}

		switch paras[0].(type) {
		case string:
			taskUUIDs = append(taskUUIDs, paras[0].(string))
		case []string:
			taskUUIDs = append(taskUUIDs, paras[0].([]string)...)
		default:
			return fmt.Errorf("invalid task uuid %s", paras[0])
		}
	}

	var taskUUID string
	taskUUIDHash := event.TaskUid.Hex()
	for _, taskUUID = range taskUUIDs {
		if crypto.Keccak256Hash([]byte(taskUUID)).Hex() == taskUUIDHash {
			break
		}
	}

	if taskUUID == "" {
		return fmt.Errorf("tx not found task equal to log task uuid hash, taskuuid_event: %s, taskuuid_raw: %s", taskUUIDHash, raw.TxHash.Hex())
	}
	return NewJobService().UpdateJobReward(taskUUID, amount)
}

func TransactionInputParas(ethClient *ethclient.Client, sigMethods map[string]abi.Method, txHash common.Hash) ([]interface{}, error) {
	tx, _, err := ethClient.TransactionByHash(context.Background(), txHash)
	if err != nil {
		return nil, err
	}
	data := tx.Data()
	if len(data) < 4 {
		return nil, errors.New("invalid tx data length")
	}
	sig := hex.EncodeToString(data[:4])
	method, ok := sigMethods[sig]
	if !ok {
		return nil, fmt.Errorf("not found method %s for tx %s", sig, txHash.Hex())
	}

	return method.Inputs.Unpack(data[4:])
}

func ContractSigMethods(contractABI string) (methods map[string]abi.Method, err error) {
	ca, err := ContractABI(contractABI)
	if err != nil {
		return
	}
	methods = make(map[string]abi.Method)
	for _, method := range ca.Methods {
		methods[hex.EncodeToString(method.ID)] = method
	}
	return
}

func ContractABI(contractABI string) (abi.ABI, error) {
	return abi.JSON(strings.NewReader(contractABI))
}

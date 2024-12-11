package ecp

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/filswan/go-swan-lib/logs"
	"github.com/swanchain/go-computing-provider/conf"
	"github.com/swanchain/go-computing-provider/internal/contract"
	"github.com/swanchain/go-computing-provider/internal/models"
	"math/big"
	"strings"
	"sync"
	"time"
)

type TaskStub struct {
	client          *ethclient.Client
	task            *Task
	privateK        string
	publicK         string
	ContractAddress string
	nonceX          uint64
	taskL           sync.Mutex
}

type TaskOption func(*TaskStub)

func WithTaskPrivateKey(pk string) TaskOption {
	return func(obj *TaskStub) {
		obj.privateK = pk
	}
}

func WithTaskContractAddress(contractAddress string) TaskOption {
	return func(obj *TaskStub) {
		obj.ContractAddress = contractAddress
	}
}

func NewTaskStub(client *ethclient.Client, options ...TaskOption) (*TaskStub, error) {
	stub := &TaskStub{}
	for _, option := range options {
		option(stub)
	}
	stub.client = client
	return stub, nil
}

func (s *TaskStub) CreateTaskContract(proof string, task *models.TaskEntity, timeOut int64) (string, error) {
	var err error
	var taskContractAddress string
	var flag bool

	cpAccountAddress, err := contract.GetCpAccountAddress()
	if err != nil {
		return "", fmt.Errorf("get cp account contract address failed, error: %v", err)
	}

	timeOutCh := time.After(time.Second * time.Duration(timeOut))
outerLoop:
	for {
		select {
		case <-timeOutCh:
			err = fmt.Errorf("create contract timed out")
			break outerLoop
		default:
			time.Sleep(3 * time.Second)
			if !flag {
				err = s.getNonce()
				if err != nil {
					logs.GetLogger().Warnf("taskId: %d, get nonce: %s, retrying", task.Id, ParseError(err))
					continue
				}
			}

			txOptions, err := s.createTransactOpts(int64(s.nonceX))
			if err != nil {
				logs.GetLogger().Warnf("taskId: %d, create transaction opts failed, error: %s", task.Id, ParseError(err))
				continue
			}

			contractAddress, transaction, _, err := DeployTask(txOptions, s.client, new(big.Int).SetInt64(task.Id), new(big.Int).SetInt64(int64(task.Type)),
				new(big.Int).SetInt64(int64(task.ResourceType)), task.InputParam, task.VerifyParam, common.HexToAddress(cpAccountAddress),
				proof, new(big.Int).SetInt64(task.Deadline), common.HexToAddress(conf.GetConfig().CONTRACT.TaskRegister), task.CheckCode)
			if err != nil {
				if err.Error() == "replacement transaction underpriced" {
					s.IncrementNonce()
					flag = true
				} else if strings.Contains(err.Error(), "next nonce") {
					err = s.getNonce()
					if err != nil {
						logs.GetLogger().Warnf("taskId: %d, get nonce: %s, retrying", task.Id, ParseError(err))
						flag = false
						continue
					}
				} else {
					logs.GetLogger().Warnf("taskId: %d create task contract failed, error: %s", task.Id, ParseError(err))
					continue
				}
			}
			if transaction != nil {
				if err = checkTransaction(s.client, transaction); err != nil {
					logs.GetLogger().Warnf("taskId: %d checked create task contract failed, error: %s", task.Id, ParseError(err))
					continue
				}
				taskContractAddress = contractAddress.Hex()
				break outerLoop
			} else {
				logs.GetLogger().Warnf("taskId: %d create task contract is nil, retrying", task.Id)
			}
		}
	}
	return taskContractAddress, err
}

func (s *TaskStub) GetTaskInfo() (models.EcpTaskInfo, error) {
	if s.ContractAddress == "" {
		return models.EcpTaskInfo{}, errors.New("missing task contract address")
	}

	taskClient, err := NewTask(common.HexToAddress(s.ContractAddress), s.client)
	if err != nil {
		return models.EcpTaskInfo{}, fmt.Errorf("create task contract client, error: %+v", err)
	}

	return taskClient.TaskInfo(&bind.CallOpts{})
}

func (s *TaskStub) privateKeyToPublicKey() (common.Address, error) {
	if len(strings.TrimSpace(s.privateK)) == 0 {
		return common.Address{}, fmt.Errorf("wallet address private key must be not empty")
	}

	privateKey, err := crypto.HexToECDSA(s.privateK)
	if err != nil {
		return common.Address{}, fmt.Errorf("parses private key error: %+v", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return common.Address{}, fmt.Errorf("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	return crypto.PubkeyToAddress(*publicKeyECDSA), nil
}

func (s *TaskStub) createTransactOpts(nonce int64) (*bind.TransactOpts, error) {
	suggestGasPrice, err := s.client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, fmt.Errorf("task client retrieves the currently suggested gas price, error: %+v", err)
	}

	chainId, err := s.client.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("task client get networkId, error: %+v", err)
	}

	privateKey, err := crypto.HexToECDSA(s.privateK)
	if err != nil {
		return nil, fmt.Errorf("parses private key error: %+v", err)
	}

	txOptions, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		return nil, fmt.Errorf("collateral client create transaction, error: %+v", err)
	}
	txOptions.Nonce = big.NewInt(nonce)
	suggestGasPrice = suggestGasPrice.Mul(suggestGasPrice, big.NewInt(3))
	suggestGasPrice = suggestGasPrice.Div(suggestGasPrice, big.NewInt(2))
	txOptions.GasFeeCap = suggestGasPrice
	txOptions.Context = context.Background()
	return txOptions, nil
}

func (s *TaskStub) getNonce() error {
	publicAddress, err := s.privateKeyToPublicKey()
	if err != nil {
		return err
	}
	nonce, err := s.client.NonceAt(context.Background(), publicAddress, nil)
	if err != nil {
		return fmt.Errorf("address: %s, collateral client get nonce error: %s", publicAddress, ParseError(err))
	}
	s.taskL.Lock()
	defer s.taskL.Unlock()
	s.nonceX = nonce
	return nil
}

func (s *TaskStub) IncrementNonce() {
	s.taskL.Lock()
	defer s.taskL.Unlock()
	s.nonceX++
}

func checkTransaction(client *ethclient.Client, tx *types.Transaction) error {
	timeout := time.After(10 * time.Second)
	ticker := time.Tick(2 * time.Second)
	for {
		select {
		case <-timeout:
			return fmt.Errorf("timeout waiting for transaction confirmation, tx: %s", tx.Hash().Hex())
		case <-ticker:
			receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
			if err != nil {
				if errors.Is(err, ethereum.NotFound) {
					continue
				}
				return err
			}

			if receipt != nil && receipt.Status == types.ReceiptStatusSuccessful {
				return nil
			}
		}
	}
}

func ParseError(err error) string {
	if strings.Contains(err.Error(), "503") {
		return "503 Service Temporarily Unavailable"
	}
	return err.Error()
}

func ParseTooManyError(err error) string {
	if strings.Contains(err.Error(), "Too Many Requests") {
		return "429 Too Many Requests"
	}
	return err.Error()
}

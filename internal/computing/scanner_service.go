package computing

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/filswan/go-mcs-sdk/mcs/api/common/logs"
	"github.com/swanchain/go-computing-provider/conf"
	"github.com/swanchain/go-computing-provider/internal/contract"
	"github.com/swanchain/go-computing-provider/internal/contract/fcp"
	"github.com/swanchain/go-computing-provider/internal/models"
)

type TaskManagerContract struct {
	cpAccount         string
	taskManagerClient *fcp.FcpTaskManager
}

func NewTaskManagerContract() *TaskManagerContract {
	cpAccountAddress, err := contract.GetCpAccountAddress()
	if err != nil {
		logs.GetLogger().Errorf("get cp account contract address failed, error: %v", err)
		return nil
	}
	return &TaskManagerContract{
		cpAccount: cpAccountAddress,
	}
}

func (taskManager *TaskManagerContract) Scan(job *models.JobEntity) {
	scannedBlock := taskManager.ScannedBlock(job)
	start := scannedBlock
	if scannedBlock != 0 {
		start = scannedBlock + 1
	}
	var endBlockNumber uint64
	var err error
	defer func() {
		if err != nil {
			//logs.GetLogger().Errorf("task manager contract task_uuid: %s  start at: %d, end scan at %d, failed: %v", job.TaskUuid, start, endBlockNumber, ecp.ParseError(err))
			return
		}
	}()

	chainUrl, err := conf.GetRpcByNetWorkName()
	if err != nil {
		logs.GetLogger().Errorf("failed to get rpc url, error: %v", err)
		return
	}
	client, err := ethclient.Dial(chainUrl)
	if err != nil {
		logs.GetLogger().Errorf("failed to dial rpc, error: %v", err)
		return
	}
	defer client.Close()

	taskManager.taskManagerClient, err = fcp.NewFcpTaskManager(common.HexToAddress(conf.GetConfig().CONTRACT.JobManager), client)
	if err != nil {
		logs.GetLogger().Errorf("failed to create job manager contract client, error: %+v", err)
		return
	}

	endBlockNumber, err = client.BlockNumber(context.Background())
	if err != nil {
		return
	}

	var step uint64 = 1000
	for i := start; i <= endBlockNumber; i = i + step {
		end := i + step - 1
		if end > endBlockNumber {
			end = endBlockNumber
		}
		filterOps := &bind.FilterOpts{
			Start: i,
			End:   &end,
		}

		if err := taskManager.scanTaskRewards(job, filterOps); err != nil {
			logs.GetLogger().Error(err)
			return
		}
		if err := NewJobService().UpdateJobScannedBlock(job.TaskUuid, end); err != nil {
			logs.GetLogger().Errorf("save job %s scanned block %d err: %v", job.TaskUuid, end, err)
		}
	}
}

func (taskManager *TaskManagerContract) scanTaskRewards(job *models.JobEntity, opts *bind.FilterOpts) (err error) {
	iterator, err := taskManager.taskManagerClient.FilterRewardReleased(opts, []string{job.TaskUuid},
		[]common.Address{common.HexToAddress(taskManager.cpAccount)})
	if err != nil {
		return err
	}
	defer iterator.Close()
	for iterator.Next() {
		event := iterator.Event
		if event != nil {
			amount := contract.BalanceToStr(event.RewardAmount)
			NewJobService().UpdateJobReward(job.TaskUuid, amount)
		}
	}
	return nil
}

func (taskManager *TaskManagerContract) ScannedBlock(job *models.JobEntity) uint64 {
	var block uint64
	if job.ScannedBlock != 0 {
		block = job.ScannedBlock
	} else {
		block = job.StartedBlock
	}
	return block
}

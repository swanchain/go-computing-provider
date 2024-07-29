package computing

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/filswan/go-mcs-sdk/mcs/api/common/logs"
	"github.com/swanchain/go-computing-provider/conf"
	"github.com/swanchain/go-computing-provider/internal/contract"
	"github.com/swanchain/go-computing-provider/internal/contract/ecp"
	"github.com/swanchain/go-computing-provider/internal/contract/fcp"
	"github.com/swanchain/go-computing-provider/internal/models"
	"time"
)

type TaskManagerContract struct {
	cpAccount string
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

	logs.GetLogger().Infof("task manager contract task_uuid: %s start scan from %d", job.TaskUuid, start)
	var endBlockNumber uint64
	var err error
	defer func() {
		if err != nil {
			logs.GetLogger().Errorf("task manager contract task_uuid: %s end scan at %d, failed: %v", job.TaskUuid, endBlockNumber, ecp.ParseError(err))
			return
		}
		logs.GetLogger().Infof("task manager contract task_uuid: %s end scan at %d", job.TaskUuid, endBlockNumber)
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

	taskManagerClient, err := fcp.NewFcpTaskManager(common.HexToAddress(conf.GetConfig().CONTRACT.JobManager), client)
	if err != nil {
		logs.GetLogger().Errorf("failed to create job manager contract client, error: %+v", err)
		return
	}

	endBlockNumber, err = client.BlockNumber(context.Background())
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	var step uint64 = 5000
	for i := start; i <= endBlockNumber; i = i + step {
		end := i + step - 1
		if end > endBlockNumber {
			end = endBlockNumber
		}
		filterOps := &bind.FilterOpts{
			Start: i,
			End:   &end,
		}

		filterRewardReleased, err := taskManagerClient.FilterRewardReleased(filterOps, []string{job.TaskUuid},
			[]common.Address{common.HexToAddress(taskManager.cpAccount)})
		if err != nil {
			logs.GetLogger().Errorf("task manager contract scan task_uuid %s from %d to %d, failed: %v", job.TaskUuid, filterOps.Start, *filterOps.End, ecp.ParseError(err))
			return
		} else {
			if filterRewardReleased.Event != nil {
				amount := contract.BalanceToStr(filterRewardReleased.Event.RewardAmount)
				NewJobService().UpdateJobReward(job.TaskUuid, amount)
			}
			if err := NewJobService().UpdateJobScannedBlock(job.TaskUuid, end); err != nil {
				logs.GetLogger().Errorf("save job %s scanned block %d err: %v", job.TaskUuid, end, err)
			}
			time.Sleep(time.Second)
		}

	}
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

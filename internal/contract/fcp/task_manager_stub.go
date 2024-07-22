package fcp

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/swanchain/go-computing-provider/conf"
	"github.com/swanchain/go-computing-provider/internal/models"
	"math/big"
)

type TaskManagerStub struct {
	client      *ethclient.Client
	taskManager *FcpTaskManager
	privateK    string
	publicK     string
}

func NewTaskManagerStub(client *ethclient.Client) (*TaskManagerStub, error) {
	stub := &TaskManagerStub{}

	jobManagerAddress := common.HexToAddress(conf.GetConfig().CONTRACT.JobManager)
	taskManagerClient, err := NewFcpTaskManager(jobManagerAddress, client)
	if err != nil {
		return nil, fmt.Errorf("failed to create job manager contract client, error: %+v", err)
	}
	stub.taskManager = taskManagerClient
	stub.client = client
	return stub, nil
}

func (s *TaskManagerStub) GetTaskInfo(taskUuid string) (models.TaskInfoOnChain, error) {
	var taskInfo models.TaskInfoOnChain
	fullTaskInfo, err := s.taskManager.GetTaskInfo(&bind.CallOpts{}, taskUuid)
	if err != nil {
		return taskInfo, fmt.Errorf("failed to get task info, task_uuid: %s, error: %+v", taskUuid, err)
	}

	taskInfo.TaskUuid = fullTaskInfo.TaskUid
	taskInfo.OwnerAddress = fullTaskInfo.User.Hex()
	taskInfo.Reward = fullTaskInfo.Reward
	taskInfo.Collateral = fullTaskInfo.Collateral
	taskInfo.StartTimestamp = fullTaskInfo.StartTimestamp.Int64()
	taskInfo.TerminateTimestamp = fullTaskInfo.TerminateTimestamp.Int64()
	taskInfo.Duration = fullTaskInfo.Duration.Int64()
	taskInfo.TaskStatus = int(fullTaskInfo.TaskStatus)
	taskInfo.CollateralStatus = int(fullTaskInfo.CollateralStatus)

	var cpAccount []string
	for _, address := range fullTaskInfo.CpList {
		cpAccount = append(cpAccount, address.Hex())
	}
	taskInfo.CpList = cpAccount
	return taskInfo, nil
}

type TaskInfoOnChain struct {
	TaskUuid           string
	CpList             []string
	OwnerAddress       string
	Reward             *big.Int
	Collateral         *big.Int
	StartTimestamp     int64
	TerminateTimestamp int64
	Duration           int64
	TaskStatus         uint8
	CollateralStatus   uint8
}

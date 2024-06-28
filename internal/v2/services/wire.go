//go:build wireinject
// +build wireinject

package services

import (
	"github.com/google/wire"
	"github.com/swanchain/go-computing-provider/internal/common"
)

func NewTaskService() common.TaskService {
	wire.Build(common.taskSet)
	return common.TaskService{}
}

func NewJobService() common.JobService {
	wire.Build(common.provideRepository)
	return common.JobService{}
}

func NewCpInfoService() common.CpInfoService {
	wire.Build(common.cpInfoSet)
	return common.CpInfoService{}
}

//go:build wireinject
// +build wireinject

package computing

import (
	"github.com/google/wire"
)

func NewTaskService() TaskService {
	wire.Build(taskSet)
	return TaskService{}
}

func NewJobService() JobService {
	wire.Build(jobSet)
	return JobService{}
}

func NewCpInfoService() CpInfoService {
	wire.Build(cpInfoSet)
	return CpInfoService{}
}

func NewEcpJobService() EcpJobService {
	wire.Build(ecpJobSet)
	return EcpJobService{}
}

func NewCpBalanceService() CpBalanceService {
	wire.Build(cpBalanceSet)
	return CpBalanceService{}
}

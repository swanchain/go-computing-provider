package service

import (
	"github.com/google/wire"
	"github.com/swanchain/go-computing-provider/internal/computing"
	"github.com/swanchain/go-computing-provider/internal/db"
)

func NewJobService() *JobService {
	wire.Build(computing.NewK8sService, db.NewDbService)
	return &JobService{}
}

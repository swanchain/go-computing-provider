package service

import (
	"github.com/google/wire"
	"github.com/swanchain/go-computing-provider/internal/db"
	"github.com/swanchain/go-computing-provider/internal/pkg"
)

func NewJobService() *JobService {
	wire.Build(pkg.NewK8sService, db.NewDbService)
	return &JobService{}
}

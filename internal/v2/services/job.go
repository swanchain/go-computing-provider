package services

import (
	"context"
	"github.com/swanchain/go-computing-provider/internal/common"
	"github.com/swanchain/go-computing-provider/internal/v2/models"
	"gorm.io/gorm"
)

type JobService struct {
	*gorm.DB
	k8sClient common.K8sService
}

func (job *JobService) CreateJob(ctx context.Context, createJobReq models.CreateJobReq) {
	//requestId := ctx.Value("request_id")

}

func (job *JobService) ExtendJob(ctx context.Context) {

}

func (job *JobService) TerminateJob(ctx context.Context) {

}

func (job *JobService) GetJobInfo(ctx context.Context) {

}

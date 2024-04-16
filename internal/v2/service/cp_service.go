package service

import (
	"github.com/swanchain/go-computing-provider/internal/computing"
	"github.com/swanchain/go-computing-provider/internal/v2/models"
	"gorm.io/gorm"
)

type JobService struct {
	*gorm.DB
	k8sClient computing.K8sService
}

func (js *JobService) doJob(job models.Job) {

}

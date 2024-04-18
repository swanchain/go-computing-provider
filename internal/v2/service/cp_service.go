package service

import (
	"github.com/swanchain/go-computing-provider/internal/pkg"
	"github.com/swanchain/go-computing-provider/internal/v2/models"
	"gorm.io/gorm"
)

type JobService struct {
	*gorm.DB
	k8sClient pkg.K8sService
}

func (js *JobService) doJob(job models.CreateJobReq) {

}

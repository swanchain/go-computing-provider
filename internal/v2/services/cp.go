package services

import (
	"context"
	"github.com/swanchain/go-computing-provider/internal/common"
	"gorm.io/gorm"
)

type CpService struct {
	*gorm.DB
	k8sClient common.K8sService
}

func (job *CpService) GetCpResources(ctx context.Context) {

}

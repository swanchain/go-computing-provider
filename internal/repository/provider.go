package repository

import (
	"github.com/google/wire"
	"github.com/swanchain/go-computing-provider/internal/db"
	"github.com/swanchain/go-computing-provider/internal/models"
	"gorm.io/gorm"
)

type Provider struct {
	*gorm.DB
}

func (p Provider) GetCpInfoEntityByAccountAddress(accountAddress string) (*models.CpInfoEntity, error) {
	var cp models.CpInfoEntity
	err := p.Where("contract_address=?", accountAddress).Find(&cp).Error
	return &cp, err
}

func (p Provider) SaveCpInfoEntity(cp *models.CpInfoEntity) (err error) {
	p.Where("contract_address =?", cp.ContractAddress).Delete(&models.CpInfoEntity{})
	return p.Save(cp).Error
}

func (p Provider) UpdateCpInfoByNodeId(cp *models.CpInfoEntity) (err error) {
	return p.Where("node_id =?", cp.NodeId).Updates(cp).Error
}

var provideRepository = wire.NewSet(db.NewDbService, wire.Struct(new(Provider), "*"))

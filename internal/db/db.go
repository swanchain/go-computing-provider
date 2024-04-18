package db

import (
	_ "embed"
	"github.com/filswan/go-swan-lib/logs"
	models2 "github.com/swanchain/go-computing-provider/internal/v1/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"path"
)

const DealsDBName = "provider.db"

var DB *gorm.DB

func init() {
	var err error
	cpRepo, ok := os.LookupEnv("CP_PATH")
	if !ok {
		logs.GetLogger().Panicf("not found cp-repo, need to required")
		return
	}
	DB, err = gorm.Open(sqlite.Open(path.Join(cpRepo, DealsDBName)), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(
		&models2.ZkTaskEntity{},
		&models2.JobEntity{})
}

func NewDbService() *gorm.DB {
	return DB
}

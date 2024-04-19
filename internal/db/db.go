package db

import (
	_ "embed"
	logging "github.com/ipfs/go-log/v2"
	models2 "github.com/swanchain/go-computing-provider/internal/v1/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"path"
)

const DealsDBName = "provider.db"

var DB *gorm.DB
var dblog = logging.Logger("db")

func init() {
	var err error
	cpRepo, ok := os.LookupEnv("CP_PATH")
	if !ok {
		dblog.Panicf("not found cp-repo, need to required")
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

package db

import (
	_ "embed"
	logging "github.com/ipfs/go-log/v2"
	"github.com/swanchain/go-computing-provider/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"path"
)

const cpDBName = "provider.db"

var DB *gorm.DB
var dblog = logging.Logger("db")

func InitDb(cpRepoPath string) {
	var err error
	DB, err = gorm.Open(sqlite.Open(path.Join(cpRepoPath, cpDBName)), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(
		&models.TaskEntity{},
		&models.JobEntity{},
		&models.CpInfoEntity{})
}

func NewDbService() *gorm.DB {
	return DB
}

package db

import (
	_ "embed"
	"github.com/swanchain/go-computing-provider/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"path"
)

const cpDBName = "provider.db"

var DB *gorm.DB

func InitDb(cpRepoPath string) {
	var err error

	DB, err = gorm.Open(sqlite.Open(path.Join(cpRepoPath, cpDBName)), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if err = DB.AutoMigrate(
		&models.TaskEntity{},
		&models.JobEntity{},
		&models.CpInfoEntity{}); err != nil {
		panic("failed to auto migrate for provider db")
	}
}

func NewDbService() *gorm.DB {
	return DB
}

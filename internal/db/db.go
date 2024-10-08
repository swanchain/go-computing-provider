package db

import (
	_ "embed"
	"fmt"
	"github.com/swanchain/go-computing-provider/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"path"
	"time"
)

const cpDBName = "provider.db"

var DB *gorm.DB

func InitDb(cpRepoPath string) {
	var err error

	DB, err = gorm.Open(sqlite.Open(path.Join(cpRepoPath, cpDBName)), &gorm.Config{
		Logger: logger.New(myLog{}, logger.Config{
			SlowThreshold: 5 * time.Second, // Slow SQL threshold
			LogLevel:      logger.Error,    // Log level
			Colorful:      false,           // Disable color
		}),
	})
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

type myLog struct {
}

func (c myLog) Printf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stdout, format, args...)
}

package initializer

import (
	"github.com/filswan/go-swan-lib/logs"
	"github.com/swanchain/go-computing-provider/conf"
	"github.com/swanchain/go-computing-provider/internal/computing"
)

func ProjectInit(cpRepoPath string) {
	if err := conf.InitConfig(cpRepoPath, false); err != nil {
		logs.GetLogger().Fatal(err)
	}
	nodeID := computing.InitComputingProvider(cpRepoPath)

	computing.NewCronTask(nodeID).RunTask()

}

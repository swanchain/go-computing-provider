package initializer

import (
	"github.com/filswan/go-swan-lib/logs"
	"github.com/swanchain/go-computing-provider/conf"
	"github.com/swanchain/go-computing-provider/internal/common"
)

func ProjectInit(cpRepoPath string) {
	if err := conf.InitConfig(cpRepoPath, false); err != nil {
		logs.GetLogger().Fatal(err)
	}
	nodeID := common.InitComputingProvider(cpRepoPath)

	common.NewCronTask(nodeID).RunTask()

}

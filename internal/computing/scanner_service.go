package computing

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/swanchain/go-computing-provider/internal/contract/fcp"
	"github.com/swanchain/go-computing-provider/internal/models"
)

type TaskManagerContract struct {
	taskUuid string

	addr      string
	ethClient *ethclient.Client
	contract  *fcp.FcpTaskManager
}

func (c *TaskManagerContract) ScannedBlock(contract *models.JobEntity) (uint64, error) {
	//block, err := redis.RDB.HGet(context.Background(), c.scannedBlockKey(), contract.Addr).Uint64()
	//if err == redis.ErrNotFound {
	//	if contract.ScannedBlock != 0 {
	//		block = contract.ScannedBlock
	//	} else {
	//		block = contract.StartedBlock
	//	}
	//	return block, nil
	//}
	//return block, err
	return 0, nil
}

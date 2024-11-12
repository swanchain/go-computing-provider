package ecp

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/swanchain/go-computing-provider/conf"
	"github.com/swanchain/go-computing-provider/internal/computing"
	"github.com/swanchain/go-computing-provider/internal/db"
	"log"
	"testing"
)

func TestTaskStub_GetTaskInfo(t *testing.T) {
	client, err := ethclient.Dial("https://rpc-proxima.swanchain.io")
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer client.Close()

	taskClient, err := NewTaskStub(client, WithTaskContractAddress("0x739Dd2e1c4Ed2D56B9c1DFcB8114b17E1A73E5B5"))
	if err != nil {
		log.Fatalln(err)
		return
	}

	info, err := taskClient.GetTaskInfo()
	if err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Printf("info: %+v", info)
}

func TestTaskPayment_ScannerTaskPaymentEvents(t *testing.T) {
	cpRepoPath := ""
	if err := conf.InitConfig(cpRepoPath, true); err != nil {
		t.Errorf("failed to load config file, error: %v", err)
		return
	}
	db.InitDb(cpRepoPath)

	computing.NewTaskPaymentService().ScannerChainGetTaskPayment()
}

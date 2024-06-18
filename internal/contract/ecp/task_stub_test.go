package ecp

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
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

package test

import (
	"fmt"
	"github.com/swanchain/go-computing-provider/conf"
	"github.com/swanchain/go-computing-provider/internal/contract"
	"github.com/swanchain/go-computing-provider/internal/contract/fcp"
	"testing"
)

func TestJobOnChain(t *testing.T) {
	if err := conf.InitConfig("/Users/sonic/cp_repo", true); err != nil {
		t.Error(err)
		return
	}

	chainRpc, err := conf.GetRpcByNetWorkName()
	if err != nil {
		t.Error(err)
		return
	}
	client, err := contract.GetEthClient(chainRpc)
	if err != nil {
		t.Error(err)
		return
	}
	defer client.Close()
	taskManagerStub, err := fcp.NewTaskManagerStub(client)
	if err != nil {
		t.Error(err)
		return
	}

	taskInfo, err := taskManagerStub.GetTaskInfo("326ffafb-9b21-4075-b457-e21011a1dc2f")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("%+v", taskInfo)
}

//{TaskUuid:326ffafb-9b21-4075-b457-e21011a1dc2f CpList:[0x6996703E0eF41efE2fd1127cF927194e983b5033 0x2bd6a6f41b37152677F8b4946490580F63494abD] OwnerAddress:0xFbc1d38a2127D81BFe3EA347bec7310a1cfa2373 Reward:+6950000000000000000 Collateral:+10000000000000000 StartTimestamp:1727082984 TerminateTimestamp:0 Duration:25200 TaskStatus:1 CollateralStatus:0}--- PASS: TestJobOnChain (0.73s)
//{TaskUuid:326ffafb-9b21-4075-b457-e21011a1dc2f CpList:[0x6996703E0eF41efE2fd1127cF927194e983b5033 0x2bd6a6f41b37152677F8b4946490580F63494abD] OwnerAddress:0xFbc1d38a2127D81BFe3EA347bec7310a1cfa2373 Reward:+7950000000000000000 Collateral:+10000000000000000 StartTimestamp:1727082984 TerminateTimestamp:0 Duration:28800 TaskStatus:1 CollateralStatus:0}--- PASS: TestJobOnChain (0.73s)

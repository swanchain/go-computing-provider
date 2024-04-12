package ubi

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/filswan/go-swan-lib/logs"
	"github.com/swanchain/go-computing-provider/account"
	"github.com/swanchain/go-computing-provider/conf"
	"github.com/swanchain/go-computing-provider/internal/models"
	"github.com/swanchain/go-computing-provider/wallet"
	coreV1 "k8s.io/api/core/v1"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
)

type TaskType struct {
}

var TaskHandleMap = make(map[string]TaskHandle)

func RegisterUbiTaskHandles(ubiTaskHandle ...TaskHandle) {
	for _, h := range ubiTaskHandle {
		TaskHandleMap[h.TaskType()] = h
	}
}

func GetUbiTaskHandle(zkType string) TaskHandle {
	if strings.Contains(zkType, TASK_ZK_TYPE_FIL_C2) {
		return TaskHandleMap[TASK_ZK_TYPE_FIL_C2]
	}

	h, ok := TaskHandleMap[zkType]
	if ok {
		return h
	}
	return nil
}

type TaskHandle interface {
	TaskType() string

	/*
		The logic code that actually does the task
		params:
			ubiTask: Task request params
			resourceRequirements: Computing resources required by the task, used to create k8s job objects. No operation required, use it directly
			imageName: Image name
			namespace：K8s namespace
			nodeName: The node name selected according to the resource conditions of the task, used to specify the node
			useGpu: Use gpu resources, true: used; false: unused
	*/
	DoTask(ubiTask models.UBITaskReq, resourceRequirements coreV1.ResourceRequirements, imageName, namespace, nodeName string, useGpu bool)
	ReceiveUbiProof(models.ReceiveProof)
	GetConfGpuName() string

	// GetImageName architecture: AMD or INTEL
	GetImageName(architecture string, useGpu bool) string
}

func submitUBIProof(c2Proof models.ReceiveProof) (string, error) {
	chainUrl, err := conf.GetRpcByName(conf.DefaultRpc)
	if err != nil {
		logs.GetLogger().Errorf("get rpc url failed, error: %v,", err)
		return "", err
	}
	client, err := ethclient.Dial(chainUrl)
	if err != nil {
		logs.GetLogger().Errorf("dial rpc connect failed, error: %v,", err)
		return "", err
	}
	client.Close()

	cpStub, err := account.NewAccountStub(client)
	if err != nil {
		logs.GetLogger().Errorf("create ubi task client failed, error: %v,", err)
		return "", err
	}
	cpAccount, err := cpStub.GetCpAccountInfo()
	if err != nil {
		logs.GetLogger().Errorf("get account info failed, error: %v,", err)
		return "", err
	}

	localWallet, err := wallet.SetupWallet(wallet.WalletRepo)
	if err != nil {
		logs.GetLogger().Errorf("setup wallet failed, error: %v,", err)
		return "", err
	}

	ki, err := localWallet.FindKey(cpAccount.OwnerAddress)
	if err != nil || ki == nil {
		logs.GetLogger().Errorf("the address: %s, private key %v,", cpAccount.OwnerAddress, wallet.ErrKeyInfoNotFound)
		return "", err
	}

	accountStub, err := account.NewAccountStub(client, account.WithCpPrivateKey(ki.PrivateKey))
	if err != nil {
		logs.GetLogger().Errorf("create ubi task client failed, error: %v,", err)
		return "", err
	}

	taskType, err := strconv.ParseUint(c2Proof.TaskType, 10, 8)
	if err != nil {
		logs.GetLogger().Errorf("conversion to uint8 error: %v", err)
		return "", err
	}

	submitUBIProofTx, err := accountStub.SubmitUBIProof(c2Proof.TaskId, uint8(taskType), c2Proof.ZkType, c2Proof.Proof)
	if err != nil {
		logs.GetLogger().Errorf("submit ubi proof tx failed, error: %v,", err)
		return "", err
	}
	return submitUBIProofTx, nil
}

func generateString(length int) string {
	characters := "abcdefghijklmnopqrstuvwxyz"
	numbers := "0123456789"
	source := characters + numbers
	result := make([]byte, length)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		result[i] = source[rand.Intn(len(source))]
	}
	return string(result)
}

func getLocalIp() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for _, iface := range interfaces {
		if iface.Flags&net.FlagLoopback == 0 && iface.Flags&net.FlagUp != 0 {
			addrs, err := iface.Addrs()
			if err != nil {
				return "", err
			}

			for _, addr := range addrs {
				ipNet, ok := addr.(*net.IPNet)
				if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
					return ipNet.IP.String(), nil
				}
			}
		}
	}
	return "", fmt.Errorf("not found local ip")
}

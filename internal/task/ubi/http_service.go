package ubi

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/filswan/go-swan-lib/logs"
	"github.com/gin-gonic/gin"
	"github.com/swanchain/go-computing-provider/conf"
	"github.com/swanchain/go-computing-provider/constants"
	"github.com/swanchain/go-computing-provider/internal/computing"
	"github.com/swanchain/go-computing-provider/internal/models"
	"github.com/swanchain/go-computing-provider/util"
	coreV1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	TASK_TYPE_CPU = 0
	TASK_TYPE_GPU = 1
)

const (
	TASK_ZK_TYPE_FIL_C2 = "fil-c2"
	TASK_ZK_TYPE_ALEO   = "aleo-proof"
)

func DoUbiTask(c *gin.Context) {
	var ubiTask models.UBITaskReq
	if err := c.ShouldBindJSON(&ubiTask); err != nil {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.JsonError))
		return
	}
	logs.GetLogger().Infof("receive ubi task received: %+v", ubiTask)

	if ubiTask.ID == 0 {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.UbiTaskParamError, "missing required field: id"))
		return
	}
	if strings.TrimSpace(ubiTask.Name) == "" {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.UbiTaskParamError, "missing required field: name"))
		return
	}
	if ubiTask.Type != TASK_TYPE_CPU && ubiTask.Type != TASK_TYPE_GPU {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.UbiTaskParamError, "the value of task_type is 0 or 1"))
		return
	}
	if strings.TrimSpace(ubiTask.ZkType) == "" {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.UbiTaskParamError, "missing required field: zk_type"))
		return
	}
	if strings.TrimSpace(ubiTask.InputParam) == "" {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.UbiTaskParamError, "missing required field: input_param"))
		return
	}
	if strings.TrimSpace(ubiTask.Signature) == "" {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.UbiTaskParamError, "missing required field: signature"))
		return
	}

	cpRepoPath, _ := os.LookupEnv("CP_PATH")
	nodeID := GetNodeId(cpRepoPath)
	signature, err := verifySignature(conf.GetConfig().UBI.UbiEnginePk, fmt.Sprintf("%s%d", nodeID, ubiTask.ID), ubiTask.Signature)
	if err != nil {
		logs.GetLogger().Errorf("verifySignature for ubi task failed, error: %+v", err)
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.UbiTaskParamError, "sign data failed"))
		return
	}
	logs.GetLogger().Infof("ubi task sign verifing, task_id: %d, type: %s, verify: %v", ubiTask.ID, ubiTask.ZkType, signature)
	if !signature {
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.UbiTaskParamError, "signature verify failed"))
		return
	}

	var ubiTaskToRedis = new(models.CacheUbiTaskDetail)
	ubiTaskToRedis.TaskId = strconv.Itoa(ubiTask.ID)
	ubiTaskToRedis.TaskType = GetTaskTypeStr(ubiTask.Type)
	var useGpu bool
	if ubiTask.Type == TASK_TYPE_GPU {
		useGpu = true
	}

	ubiTaskToRedis.Status = constants.UBI_TASK_RECEIVED_STATUS
	ubiTaskToRedis.ZkType = ubiTask.ZkType
	ubiTaskToRedis.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	computing.SaveUbiTaskMetadata(ubiTaskToRedis)

	ubiTaskHandle := GetUbiTaskHandle(ubiTask.ZkType)
	if ubiTaskHandle != nil {
		confGpuName := ubiTaskHandle.GetConfGpuName()
		nodeName, architecture, needCpu, needMemory, needStorage, err := checkResourceAvailableForUbi(ubiTask.Type, confGpuName, ubiTask.Resource)
		if err != nil {
			ubiTaskToRedis.Status = constants.UBI_TASK_FAILED_STATUS
			computing.SaveUbiTaskMetadata(ubiTaskToRedis)
			logs.GetLogger().Errorf("check resource failed, error: %v", err)
			c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.CheckResourcesError))
			return
		}

		if nodeName == "" {
			ubiTaskToRedis.Status = constants.UBI_TASK_FAILED_STATUS
			computing.SaveUbiTaskMetadata(ubiTaskToRedis)
			logs.GetLogger().Warnf("ubi task id: %d, type: %s, not found a resources available", ubiTask.ID, ubiTaskToRedis.TaskType)
			c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.CheckAvailableResources))
			return
		}

		resourceRequirements, err := generateResource(ubiTask.Resource, needCpu, needMemory, needStorage, useGpu)
		if err != nil {
			ubiTaskToRedis.Status = constants.UBI_TASK_FAILED_STATUS
			computing.SaveUbiTaskMetadata(ubiTaskToRedis)
			return
		}

		imageName := ubiTaskHandle.GetImageName(architecture, useGpu)
		if imageName == "" {
			logs.GetLogger().Warnf("ubi task id: %d, type: %s, not found a resources available", ubiTask.ID, ubiTaskToRedis.TaskType)
			return
		}

		var namespace = strings.ToLower(fmt.Sprintf("ubi-%s-%d", ubiTask.ZkType, ubiTask.ID))
		k8sService := computing.NewK8sService()
		if _, err = k8sService.GetNameSpace(context.TODO(), namespace, metaV1.GetOptions{}); err != nil {
			if errors.IsNotFound(err) {
				k8sNamespace := &coreV1.Namespace{
					ObjectMeta: metaV1.ObjectMeta{
						Name: namespace,
					},
				}
				_, err = k8sService.CreateNameSpace(context.TODO(), k8sNamespace, metaV1.CreateOptions{})
				if err != nil {
					logs.GetLogger().Errorf("create namespace failed, error: %v", err)
					return
				}
			}
		}

		ubiTaskHandle.DoTask(ubiTask, resourceRequirements, imageName, namespace, nodeName, useGpu)
		c.JSON(http.StatusOK, util.CreateSuccessResponse("success"))
	} else {
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.UbiTaskError, "not found task handle"))
		return
	}
}

func ReceiveUbiProof(c *gin.Context) {
	var c2Proof models.ReceiveProof
	if err := c.ShouldBindJSON(&c2Proof); err != nil {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.JsonError))
		return
	}
	logs.GetLogger().Infof("task_id: %s, proof of ubi task out received: %+v", c2Proof.TaskId, c2Proof)

	ubiTaskHandle := GetUbiTaskHandle(c2Proof.ZkType)
	if ubiTaskHandle != nil {
		ubiTaskHandle.ReceiveUbiProof(c2Proof)
	}

	c.JSON(http.StatusOK, util.CreateSuccessResponse("success"))
}

func generateResource(ubiResource *models.TaskResource, needCpu, needMemory, needStorage int64, useGpu bool) (coreV1.ResourceRequirements, error) {
	mem := strings.Split(strings.TrimSpace(ubiResource.Memory), " ")[1]
	memUnit := strings.ReplaceAll(mem, "B", "")
	disk := strings.Split(strings.TrimSpace(ubiResource.Storage), " ")[1]
	diskUnit := strings.ReplaceAll(disk, "B", "")
	memQuantity, err := resource.ParseQuantity(fmt.Sprintf("%d%s", needMemory, memUnit))
	if err != nil {

		logs.GetLogger().Error("get memory failed, error: %+v", err)
		return coreV1.ResourceRequirements{}, err
	}

	storageQuantity, err := resource.ParseQuantity(fmt.Sprintf("%d%s", needStorage, diskUnit))
	if err != nil {
		logs.GetLogger().Error("get storage failed, error: %+v", err)
		return coreV1.ResourceRequirements{}, err
	}

	maxMemQuantity, err := resource.ParseQuantity(fmt.Sprintf("%d%s", needMemory*2, memUnit))
	if err != nil {
		logs.GetLogger().Error("get memory failed, error: %+v", err)
		return coreV1.ResourceRequirements{}, err
	}

	maxStorageQuantity, err := resource.ParseQuantity(fmt.Sprintf("%d%s", needStorage*2, diskUnit))
	if err != nil {
		logs.GetLogger().Error("get storage failed, error: %+v", err)
		return coreV1.ResourceRequirements{}, err
	}

	var gpuNum = "0"
	if useGpu {
		gpuNum = "1"
	}
	return coreV1.ResourceRequirements{
		Limits: coreV1.ResourceList{
			coreV1.ResourceCPU:              *resource.NewQuantity(needCpu*2, resource.DecimalSI),
			coreV1.ResourceMemory:           maxMemQuantity,
			coreV1.ResourceEphemeralStorage: maxStorageQuantity,
			"nvidia.com/gpu":                resource.MustParse(gpuNum),
		},
		Requests: coreV1.ResourceList{
			coreV1.ResourceCPU:              *resource.NewQuantity(needCpu, resource.DecimalSI),
			coreV1.ResourceMemory:           memQuantity,
			coreV1.ResourceEphemeralStorage: storageQuantity,
			"nvidia.com/gpu":                resource.MustParse(gpuNum),
		},
	}, nil
}

func checkResourceAvailableForUbi(taskType int, gpuName string, resource *models.TaskResource) (string, string, int64, int64, int64, error) {
	k8sService := computing.NewK8sService()
	activePods, err := k8sService.GetAllActivePod(context.TODO())
	if err != nil {
		return "", "", 0, 0, 0, err
	}

	nodes, err := k8sService.Client.CoreV1().Nodes().List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return "", "", 0, 0, 0, err
	}

	nodeGpuSummary, err := k8sService.GetNodeGpuSummary(context.TODO())
	if err != nil {
		logs.GetLogger().Errorf("Failed collect k8s gpu, error: %+v", err)
		return "", "", 0, 0, 0, err
	}

	needCpu, _ := strconv.ParseInt(resource.CPU, 10, 64)
	var needMemory, needStorage float64
	if len(strings.Split(strings.TrimSpace(resource.Memory), " ")) > 0 {
		needMemory, err = strconv.ParseFloat(strings.Split(strings.TrimSpace(resource.Memory), " ")[0], 64)

	}
	if len(strings.Split(strings.TrimSpace(resource.Storage), " ")) > 0 {
		needStorage, err = strconv.ParseFloat(strings.Split(strings.TrimSpace(resource.Storage), " ")[0], 64)
	}

	var nodeName, architecture string
	for _, node := range nodes.Items {
		if _, ok := node.Labels[constants.CPU_INTEL]; ok {
			architecture = constants.CPU_INTEL
		}
		if _, ok := node.Labels[constants.CPU_AMD]; ok {
			architecture = constants.CPU_AMD
		}

		nodeGpu, remainderResource, _ := computing.GetNodeResource(activePods, &node)
		remainderCpu := remainderResource[computing.ResourceCpu]
		remainderMemory := float64(remainderResource[computing.ResourceMem] / 1024 / 1024 / 1024)
		remainderStorage := float64(remainderResource[computing.ResourceStorage] / 1024 / 1024 / 1024)

		logs.GetLogger().Infof("checkResourceAvailableForUbi: needCpu: %d, needMemory: %.2f, needStorage: %.2f", needCpu, needMemory, needStorage)
		logs.GetLogger().Infof("checkResourceAvailableForUbi: remainingCpu: %d, remainingMemory: %.2f, remainingStorage: %.2f", remainderCpu, remainderMemory, remainderStorage)
		if needCpu <= remainderCpu && needMemory <= remainderMemory && needStorage <= remainderStorage {
			nodeName = node.Name
			if taskType == 0 {
				return nodeName, architecture, needCpu, int64(needMemory), int64(needStorage), nil
			} else if taskType == 1 {
				if gpuName == "" {
					nodeName = ""
					continue
				}
				gpuName = strings.ReplaceAll(gpuName, " ", "-")
				logs.GetLogger().Infof("gpuName: %s, nodeGpu: %+v, nodeGpuSummary: %+v", gpuName, nodeGpu, nodeGpuSummary)
				usedCount, ok := nodeGpu[gpuName]
				if !ok {
					usedCount = 0
				}

				if usedCount+1 <= nodeGpuSummary[node.Name][gpuName] {
					return nodeName, architecture, needCpu, int64(needMemory), int64(needStorage), nil
				}
				nodeName = ""
				continue
			}
		}
	}
	return nodeName, architecture, needCpu, int64(needMemory), int64(needStorage), nil
}

func GetNodeId(cpRepoPath string) string {
	nodeID, _, _ := computing.GenerateNodeID(cpRepoPath)
	return nodeID
}

func convertGpuName(name string) string {
	if strings.TrimSpace(name) == "" {
		return ""
	} else {
		name = strings.Split(name, ":")[0]
	}
	if strings.Contains(name, "NVIDIA") {
		if strings.Contains(name, "Tesla") {
			return strings.Replace(name, "Tesla ", "", 1)
		}

		if strings.Contains(name, "GeForce") {
			name = strings.Replace(name, "GeForce ", "", 1)
		}
		return strings.Replace(name, "RTX ", "", 1)
	} else {
		if strings.Contains(name, "GeForce") {
			cpName := strings.Replace(name, "GeForce ", "NVIDIA", 1)
			return strings.Replace(cpName, "RTX", "", 1)
		}
	}
	return name
}

func verifySignature(pubKStr, data, signature string) (bool, error) {
	sb, err := hexutil.Decode(signature)
	if err != nil {
		return false, err
	}
	hash := crypto.Keccak256Hash([]byte(data))
	sigPublicKeyECDSA, err := crypto.SigToPub(hash.Bytes(), sb)
	if err != nil {
		return false, err
	}
	pub := crypto.PubkeyToAddress(*sigPublicKeyECDSA).Hex()
	if pubKStr != pub {
		return false, err
	}
	return true, nil
}

func GetTaskTypeStr(taskType int) string {
	switch taskType {
	case TASK_TYPE_CPU:
		return "CPU"
	case TASK_TYPE_GPU:
		return "GPU"
	}
	return ""
}

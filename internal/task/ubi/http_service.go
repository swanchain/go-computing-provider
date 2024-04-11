package ubi

import (
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
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func DoUbiTask(c *gin.Context, ubiTask models.UBITaskReq) {
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

	var gpuFlag = "0"
	var ubiTaskToRedis = new(models.CacheUbiTaskDetail)
	ubiTaskToRedis.TaskId = strconv.Itoa(ubiTask.ID)
	ubiTaskToRedis.TaskType = "CPU"
	if ubiTask.Type == 1 {
		ubiTaskToRedis.TaskType = "GPU"
		gpuFlag = "1"
	}
	ubiTaskToRedis.Status = constants.UBI_TASK_RECEIVED_STATUS
	ubiTaskToRedis.ZkType = ubiTask.ZkType
	ubiTaskToRedis.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	computing.SaveUbiTaskMetadata(ubiTaskToRedis)
	println(gpuFlag)

}

func ReceiveUbiProof(c *gin.Context, c2Proof models.ReceiveProof) {

}

func GetNodeId(cpRepoPath string) string {
	nodeID, _, _ := computing.GenerateNodeID(cpRepoPath)
	return nodeID
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

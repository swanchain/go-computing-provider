package computing

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/filswan/go-swan-lib/logs"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/swanchain/go-computing-provider/conf"
	"github.com/swanchain/go-computing-provider/constants"
	"github.com/swanchain/go-computing-provider/internal/models"
	"github.com/swanchain/go-computing-provider/util"
	"io"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

func ReceivePrivateJob(c *gin.Context) {
	if conf.GetConfig().HUB.BidMode == constants.BidMode_Auto || conf.GetConfig().HUB.BidMode == constants.BidMode_None {
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.ServerError, "the provider does not accept private task"))
		return
	}

	var jobData models.PrivateJobReq
	if err := c.ShouldBindJSON(&jobData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	logs.GetLogger().Infof("private job received Data: %+v", jobData)

	if len(jobData.Signature) == 0 {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.SignatureError, "missing signature field"))
		return
	}

	if jobData.Config.Vcpu == 0 {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.BadParamError, "missing config.Vcpu field"))
		return
	}

	if jobData.Config.Memory == 0 {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.BadParamError, "missing config.Memory field"))
		return
	}

	if jobData.Config.Storage == 0 {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.BadParamError, "missing config.Storage field"))
		return
	}

	if jobData.Config.GPU != 0 && len(strings.TrimSpace(jobData.Config.GPUModel)) == 0 {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.BadParamError, "missing config.gpu_model field"))
	}

	if jobData.Config.GPU == 0 && len(strings.TrimSpace(jobData.Config.GPUModel)) != 0 {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.BadParamError, "missing config.gpu field"))
	}

	cpRepoPath, _ := os.LookupEnv("CP_PATH")
	nodeID := GetNodeId(cpRepoPath)

	signature, err := verifySignatureForHub(conf.GetConfig().UBI.UbiEnginePk, fmt.Sprintf("%s%s", nodeID, jobData.UUID), jobData.Signature)
	if err != nil {
		logs.GetLogger().Errorf("verifySignature for private job failed, error: %+v", err)
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.ServerError, "verify sign data failed"))
		return
	}

	logs.GetLogger().Infof("private job sign verifing, task_id: %s,  verify: %v", jobData.UUID, signature)
	if !signature {
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.SignatureError, "signature verify failed"))
		return
	}

	available, gpuProductName, err := checkResourceAvailableForPrivate(jobData.Config.Vcpu, jobData.Config.Memory, jobData.Config.Storage, jobData.Config.GPUModel, jobData.Config.GPU)
	if err != nil {
		logs.GetLogger().Errorf("check job resource failed, error: %+v", err)
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.CheckResourcesError))
		return
	}

	if !available {
		logs.GetLogger().Warnf(" task id: %s, name: %s, not found a resources available", jobData.UUID, jobData.Name)
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.CheckResourcesError))
		return
	}

	var logHost string
	if strings.HasPrefix(conf.GetConfig().API.Domain, ".") {
		logHost = "log" + conf.GetConfig().API.Domain
	} else {
		logHost = "log." + conf.GetConfig().API.Domain
	}

	go func() {
		DeployPrivateTask(jobData.Name, logHost, jobData.Duration, jobData.UUID, gpuProductName, jobData.User,
			jobData.Config.Vcpu, jobData.Config.Memory, jobData.Config.Storage, jobData.Config.GPUModel, jobData.Config.SshKey, jobData.Config.Image, jobData.Config.GPU)
	}()

	var privateJob models.PrivateJobResp
	privateJob.UUID = jobData.UUID
	multiAddressSplit := strings.Split(conf.GetConfig().API.MultiAddress, "/")
	wsUrl := fmt.Sprintf("wss://%s:%s/api/v1/computing/lagrange/spaces/log?space_id=%s", logHost, multiAddressSplit[4], jobData.UUID)
	privateJob.ContainerLog = wsUrl + "&type=container&order=private"
	privateJob.UpdatedAt = time.Now().Unix()
	logs.GetLogger().Infof("submit private job detail: %+v", jobData)
	c.JSON(http.StatusOK, util.CreateSuccessResponse(privateJob))
}

func ExtendJob(c *gin.Context) {
	var jobData struct {
		Uuid     string `json:"uuid"`
		Duration int    `json:"duration"`
	}

	if err := c.ShouldBindJSON(&jobData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	logs.GetLogger().Infof("extend private Job received: %+v", jobData)

	if strings.TrimSpace(jobData.Uuid) == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing required field: uuid"})
		return
	}

	if jobData.Duration == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing required field: duration"})
		return
	}

	jobEntity, err := NewJobService().GetJobEntityByTaskUuid(jobData.Uuid)
	if err != nil {
		logs.GetLogger().Errorf("Failed get job from db, taskUuid: %s, error: %+v", jobData.Uuid, err)
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.FoundJobEntityError))
		return
	}

	leftTime := jobEntity.ExpireTime - time.Now().Unix()
	if leftTime < 0 {
		c.JSON(http.StatusOK, util.CreateErrorResponse(util.BadParamError, "The job was terminated due to its expiration date"))
		return
	} else {
		if getJobExpiredTime(jobEntity) <= 0 {
			jobEntity.ExpireTime = time.Now().Unix() + leftTime + int64(jobData.Duration)
		} else {
			jobEntity.ExpireTime = getJobExpiredTime(jobEntity)
		}

		err = NewJobService().SaveJobEntity(&jobEntity)
		if err != nil {
			logs.GetLogger().Errorf("update job expireTime failed, taskUuid: %s, error: %+v", jobData.Uuid, err)
			c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.SaveJobEntityError))
			return
		}
	}
	c.JSON(http.StatusOK, util.CreateSuccessResponse("success"))
}

func TerminateJob(c *gin.Context) {
	taskUuid := c.Param("task_uuid")
	if taskUuid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "task_uuid is required"})
		return
	}

	jobEntity, err := NewJobService().GetJobEntityByTaskUuid(taskUuid)
	if err != nil {
		logs.GetLogger().Errorf("Failed get job from db, taskUuid: %s, error: %+v", taskUuid, err)
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.FoundJobEntityError))
		return
	}

	if jobEntity.WalletAddress == "" {
		c.JSON(http.StatusOK, util.CreateSuccessResponse("deleted success"))
		return
	}

	go func() {
		defer func() {
			if err := recover(); err != nil {
				logs.GetLogger().Errorf("task_uuid: %s, delete space request failed, error: %+v", taskUuid, err)
				return
			}
		}()
		k8sNameSpace := constants.K8S_NAMESPACE_NAME_PREFIX + strings.ToLower(jobEntity.WalletAddress)
		deleteJob(k8sNameSpace, jobEntity.SpaceUuid, "")
		NewJobService().DeleteJobEntityBySpaceUuId(jobEntity.SpaceUuid, models.JOB_TERMINATED_STATUS)
	}()

	c.JSON(http.StatusOK, util.CreateSuccessResponse("deleted success"))
}

func submitPrivateJob(jobData *models.PrivateJobResp) error {
	logs.GetLogger().Printf("submitting private job...")
	oldMask := syscall.Umask(0)
	defer syscall.Umask(oldMask)

	jobDetailFile := filepath.Join("jobs", uuid.NewString()+".json")
	os.MkdirAll(filepath.Join("mcs_cache", "jobs"), os.ModePerm)
	taskDetailFilePath := filepath.Join("mcs_cache", jobDetailFile)

	jobBytes, err := json.Marshal(jobData)
	if err != nil {
		logs.GetLogger().Errorf("Failed Marshal JobData, error: %v", err)
		return err
	}
	if err = os.WriteFile(taskDetailFilePath, jobBytes, os.ModePerm); err != nil {
		logs.GetLogger().Errorf("Failed jobData write to file, error: %v", err)
		return err
	}

	for i := 0; i < 5; i++ {
		storageService, err := NewStorageService()
		if err != nil {
			return err
		}
		mcsOssFile, err := storageService.UploadFileToBucket(jobDetailFile, taskDetailFilePath, true)
		if err != nil {
			logs.GetLogger().Errorf("upload file to bucket failed, error: %v", err)
			continue
		}

		if mcsOssFile == nil || mcsOssFile.PayloadCid == "" {
			continue
		}

		gatewayUrl, err := storageService.GetGatewayUrl()
		if err != nil {
			logs.GetLogger().Errorf("get mcs ipfs gatewayUrl failed, error: %v", err)
			continue
		}
		jobData.ResultURI = *gatewayUrl + "/ipfs/" + mcsOssFile.PayloadCid
		break
	}
	return nil
}

func DeployPrivateTask(name string, logHost string, duration int, taskUuid string, gpuProductName string, walletAddress string, cpu, memory, storage int, gpModel, sshKey, image string, gpuNum int) {
	updatePrivateStatus(taskUuid, JobStatusDeploying, "", "")
	var success bool
	var spaceUuid string

	defer func() {
		if !success {
			k8sNameSpace := constants.K8S_NAMESPACE_NAME_PREFIX + strings.ToLower(walletAddress)
			deleteJob(k8sNameSpace, spaceUuid, "failed to deploy private space")
			NewJobService().DeleteJobEntityBySpaceUuId(spaceUuid, models.JOB_TERMINATED_STATUS)
		}

		if err := recover(); err != nil {
			logs.GetLogger().Errorf("deploy private task painc, error: %+v", err)
			return
		}
	}()

	spaceUuid = strings.ToLower(taskUuid)

	var job = new(models.JobEntity)
	job.WalletAddress = walletAddress
	job.Name = name
	job.ExpireTime = time.Now().Unix() + int64(duration)
	job.TaskUuid = taskUuid
	if err := NewJobService().UpdateJobEntityBySpaceUuid(job); err != nil {
		logs.GetLogger().Errorf("update job info failed, error: %v", err)
		return
	}

	deploy := NewDeploy(spaceUuid, "", walletAddress, "", int64(duration), taskUuid, constants.SPACE_TYPE_PRIVATE)
	deploy.WithSpaceInfo(spaceUuid, name)
	deploy.WithGpuProductName(gpuProductName)
	deploy.WithHardware(cpu, memory, storage, gpModel, gpuNum)
	deploy.WithImage(image)

	sshUrl, err := deploy.WithSshKey(sshKey).DeploySshTaskToK8s()
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	var privateJob models.PrivateJobResp
	privateJob.UUID = taskUuid
	privateJob.RealURI = sshUrl

	multiAddressSplit := strings.Split(conf.GetConfig().API.MultiAddress, "/")
	privateJob.ContainerLog = fmt.Sprintf("wss://%s:%s/api/v1/computing/lagrange/spaces/log?space_id=%s&type=container&order=private",
		logHost, multiAddressSplit[4], taskUuid)
	privateJob.UpdatedAt = time.Now().Unix()
	privateJob.RealURI = sshUrl
	if err = submitPrivateJob(&privateJob); err != nil {
		privateJob.ResultURI = ""
	}
	updatePrivateStatus(taskUuid, JobStatusRunning, sshUrl, privateJob.ResultURI)

	success = true
}

type PrivateJobStatus int

const (
	JobStatusDeploying PrivateJobStatus = 3
	JobStatusRunning   PrivateJobStatus = 4
)

func updatePrivateStatus(jobUuid string, jobStatus PrivateJobStatus, result, resultUri string) {
	reqParam := make(map[string]interface{})
	reqParam["job_uuid"] = jobUuid
	reqParam["status"] = jobStatus
	if result != "" {
		reqParam["real_uri"] = result
	}
	if resultUri != "" {
		reqParam["result_uri"] = resultUri
	}

	payload, err := json.Marshal(reqParam)
	if err != nil {
		logs.GetLogger().Errorf("Failed convert to json, error: %+v", err)
		return
	}

	client := &http.Client{}
	url := conf.GetConfig().UBI.UbiUrl + "/jobs/" + jobUuid
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(payload))
	if err != nil {
		logs.GetLogger().Errorf("Error creating request: %v", err)
		return
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		logs.GetLogger().Errorf("Failed send a request, error: %+v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyData, _ := io.ReadAll(resp.Body)
		logs.GetLogger().Errorf("report private job status failed, %s", string(bodyData))
		return
	}
	return
}

func checkResourceAvailableForPrivate(cpu, memory, storage int, gpuModel string, gpuNum int) (bool, string, error) {
	taskType, hardwareDetail := getHardwareDetailForPrivate(cpu, memory, storage, gpuModel, gpuNum)
	k8sService := NewK8sService()

	activePods, err := k8sService.GetAllActivePod(context.TODO())
	if err != nil {
		return false, "", err
	}

	nodes, err := k8sService.k8sClient.CoreV1().Nodes().List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return false, "", err
	}

	nodeGpuSummary, err := k8sService.GetNodeGpuSummary(context.TODO())
	if err != nil {
		logs.GetLogger().Errorf("Failed collect k8s gpu, error: %+v", err)
		return false, "", err
	}

	for _, node := range nodes.Items {
		nodeGpu, remainderResource, _ := GetNodeResource(activePods, &node)
		remainderCpu := remainderResource[ResourceCpu]
		remainderMemory := float64(remainderResource[ResourceMem] / 1024 / 1024 / 1024)
		remainderStorage := float64(remainderResource[ResourceStorage] / 1024 / 1024 / 1024)

		needCpu := hardwareDetail.Cpu.Quantity
		needMemory := float64(hardwareDetail.Memory.Quantity)
		needStorage := float64(hardwareDetail.Storage.Quantity)
		logs.GetLogger().Infof("checkResourceAvailableForPrivate: needCpu: %d, needMemory: %.2f, needStorage: %.2f, gpuName: %s, gpuNum: %d", needCpu, needMemory, needStorage, gpuModel, gpuNum)
		logs.GetLogger().Infof("checkResourceAvailableForPrivate: remainingCpu: %d, remainingMemory: %.2f, remainingStorage: %.2f", remainderCpu, remainderMemory, remainderStorage)
		if needCpu <= remainderCpu && needMemory <= remainderMemory && needStorage <= remainderStorage {
			if taskType == "CPU" {
				return true, "", nil
			} else if taskType == "GPU" {
				var usedCount int64 = 0
				gpuName := strings.ToUpper(strings.ReplaceAll(hardwareDetail.Gpu.Unit, " ", "-"))
				logs.GetLogger().Infof("gpuName: %s, nodeGpu: %+v, nodeGpuSummary: %+v", gpuName, nodeGpu, nodeGpuSummary)
				var gpuProductName = ""
				for name, count := range nodeGpu {
					if strings.Contains(strings.ToUpper(name), gpuName) {
						usedCount = count
						gpuProductName = strings.ReplaceAll(strings.ToUpper(name), " ", "-")
						break
					}
				}

				for gName, gCount := range nodeGpuSummary[node.Name] {
					if strings.Contains(strings.ToUpper(gName), gpuName) {
						gpuProductName = strings.ReplaceAll(strings.ToUpper(gName), " ", "-")
						if usedCount+hardwareDetail.Gpu.Quantity <= gCount {
							return true, gpuProductName, nil
						}
					}
				}
				continue
			}
		}
	}
	return false, "", nil
}

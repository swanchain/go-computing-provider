package computing

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/filswan/go-mcs-sdk/mcs/api/common/logs"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/swanchain/go-computing-provider/build"
	"github.com/swanchain/go-computing-provider/conf"
	"github.com/swanchain/go-computing-provider/constants"
	"github.com/swanchain/go-computing-provider/internal/contract"
	"github.com/swanchain/go-computing-provider/internal/models"
	"github.com/swanchain/go-computing-provider/util"
	"io"
	batchv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func GetServiceProviderInfo(c *gin.Context) {
	info := new(models.HostInfo)
	info.SwanProviderVersion = build.UserVersion()
	info.OperatingSystem = runtime.GOOS
	info.Architecture = runtime.GOARCH
	info.CPUCores = runtime.NumCPU()
	c.JSON(http.StatusOK, util.CreateSuccessResponse(info))
}

func ReceiveJob(c *gin.Context) {
	var jobData models.JobData
	if err := c.ShouldBindJSON(&jobData); err != nil {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.JsonError))
		return
	}
	logs.GetLogger().Infof("Job received Data: %+v", jobData)

	if !CheckWalletWhiteList(jobData.JobSourceURI) {
		logs.GetLogger().Errorf("This cp does not accept tasks from wallet addresses outside the whitelist")
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.SpaceCheckWhiteListError))
		return
	}

	if CheckWalletBlackList(jobData.JobSourceURI) {
		logs.GetLogger().Errorf("This cp does not accept tasks from wallet addresses inside the blacklist")
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.SpaceCheckBlackListError))
		return
	}

	if conf.GetConfig().HUB.VerifySign {
		if len(jobData.NodeIdJobSourceUriSignature) == 0 {
			c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.BadParamError, "missing node_id_job_source_uri_signature field"))
			return
		}
		cpRepoPath, _ := os.LookupEnv("CP_PATH")
		nodeID := GetNodeId(cpRepoPath)

		cpAccountAddress, err := contract.GetCpAccountAddress()
		if err != nil {
			logs.GetLogger().Errorf("get cp account contract address failed, error: %v", err)
			c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.GetCpAccountError))
			return
		}

		signature, err := verifySignatureForHub(conf.GetConfig().HUB.OrchestratorPk, fmt.Sprintf("%s%s%s", cpAccountAddress, nodeID, jobData.JobSourceURI), jobData.NodeIdJobSourceUriSignature)
		if err != nil {
			logs.GetLogger().Errorf("verifySignature for space job failed, error: %+v", err)
			c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.SignatureError, "verify sign data occur error"))
			return
		}

		if !signature {
			logs.GetLogger().Errorf("space job sign verifing, task_id: %s, verify: %v", jobData.TaskUUID, signature)
			c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.SignatureError, "signature verify failed"))
			return
		}
	}

	spaceDetail, err := getSpaceDetail(jobData.JobSourceURI)
	if err != nil {
		logs.GetLogger().Errorln(err)
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.SpaceParseResourceUriError))
		return
	}

	available, gpuProductName, err := checkResourceAvailableForSpace(spaceDetail.Data.Space.ActiveOrder.Config.Description)
	if err != nil {
		logs.GetLogger().Errorf("check job resource failed, error: %+v", err)
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.CheckResourcesError))
		return
	}

	if !available {
		logs.GetLogger().Warnf(" task id: %s, name: %s, not found a resources available", jobData.TaskUUID, jobData.Name)
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.NoAvailableResourcesError))
		return
	}

	var hostName string
	var logHost string
	prefixStr := generateString(10)
	if strings.HasPrefix(conf.GetConfig().API.Domain, ".") {
		hostName = prefixStr + conf.GetConfig().API.Domain
		logHost = "log" + conf.GetConfig().API.Domain
	} else {
		hostName = strings.Join([]string{prefixStr, conf.GetConfig().API.Domain}, ".")
		logHost = "log." + conf.GetConfig().API.Domain
	}

	multiAddressSplit := strings.Split(conf.GetConfig().API.MultiAddress, "/")
	jobSourceUri := jobData.JobSourceURI
	spaceUuid := jobSourceUri[strings.LastIndex(jobSourceUri, "/")+1:]
	wsUrl := fmt.Sprintf("wss://%s:%s/api/v1/computing/lagrange/spaces/log?space_id=%s", logHost, multiAddressSplit[4], spaceUuid)
	jobData.BuildLog = wsUrl + "&type=build"
	jobData.ContainerLog = wsUrl + "&type=container"
	jobData.JobRealUri = fmt.Sprintf("https://%s", hostName)
	jobData.NodeIdJobSourceUriSignature = ""
	go func() {
		job, err := NewJobService().GetJobEntityBySpaceUuid(spaceUuid)
		if err != nil {
			logs.GetLogger().Errorf("get job failed, error: %+v", err)
			return
		}
		if job.SpaceUuid != "" {
			NewJobService().DeleteJobEntityBySpaceUuId(spaceUuid)
		}

		var jobEntity = new(models.JobEntity)
		jobEntity.Source = jobData.StorageSource
		jobEntity.SpaceUuid = spaceUuid
		jobEntity.TaskUuid = jobData.TaskUUID
		jobEntity.SourceUrl = jobSourceUri
		jobEntity.RealUrl = jobData.JobRealUri
		jobEntity.BuildLog = jobData.BuildLog
		jobEntity.ContainerLog = jobData.ContainerLog
		jobEntity.Duration = jobData.Duration
		jobEntity.JobUuid = jobData.UUID
		jobEntity.DeployStatus = models.DEPLOY_RECEIVE_JOB
		jobEntity.CreateTime = time.Now().Unix()
		NewJobService().SaveJobEntity(jobEntity)

		go func() {
			if err = submitJob(&jobData); err != nil {
				logs.GetLogger().Errorf("upload job data to MCS failed, jobUuid: %s, spaceUuid: %s, error: %v", jobData.UUID, spaceUuid, err)
				return
			}
			logs.GetLogger().Infof("jobuuid: %s successfully uploaded to MCS", jobData.UUID)
		}()

		DeploySpaceTask(jobData.JobSourceURI, hostName, jobData.Duration, jobData.UUID, jobData.TaskUUID, gpuProductName)
	}()

	c.JSON(http.StatusOK, util.CreateSuccessResponse(jobData))
}

func submitJob(jobData *models.JobData) error {
	cpRepoPath, ok := os.LookupEnv("CP_PATH")
	if !ok {
		return fmt.Errorf("missing CP_PATH env, please set export CP_PATH=<YOUR CP_PATH>")
	}
	folderPath := "mcs_cache"
	jobDetailFile := filepath.Join(folderPath, uuid.NewString()+".json")
	os.MkdirAll(filepath.Join(cpRepoPath, folderPath), os.ModePerm)
	taskDetailFilePath := filepath.Join(cpRepoPath, jobDetailFile)

	jobData.JobResultURI = jobData.JobRealUri
	bytes, err := json.Marshal(jobData)
	if err != nil {
		return fmt.Errorf(" parse to json failed, error: %v", err)
	}

	f, err := os.OpenFile(taskDetailFilePath, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		logs.GetLogger().Errorf("Failed to open file, error: %v", err)
		return err
	}
	defer f.Close()

	if _, err = f.Write(bytes); err != nil {
		return fmt.Errorf("write jobData to file failed, error: %v", err)
	}

	var resultMcsUrl string
	for i := 0; i < 5; i++ {
		storageService := NewStorageService()
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
		resultMcsUrl = *gatewayUrl + "/ipfs/" + mcsOssFile.PayloadCid
		break
	}
	return NewJobService().UpdateJobResultUrlByJobUuid(jobData.UUID, resultMcsUrl)
}

func RedeployJob(c *gin.Context) {
	var jobData models.JobData

	if err := c.ShouldBindJSON(&jobData); err != nil {
		logs.GetLogger().Errorln(err)
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.JsonError))
		return
	}
	logs.GetLogger().Infof("redeploy Job received: %+v", jobData)

	spaceDetail, err := getSpaceDetail(jobData.JobSourceURI)
	if err != nil {
		logs.GetLogger().Errorln(err)
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.SpaceParseResourceUriError))
		return
	}

	available, gpuProductName, err := checkResourceAvailableForSpace(spaceDetail.Data.Space.ActiveOrder.Config.Description)
	if err != nil {
		logs.GetLogger().Errorf("check job resource failed, error: %+v", err)
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.CheckResourcesError))
		return
	}

	if !available {
		logs.GetLogger().Warnf(" task id: %s, name: %s, not found a resources available", jobData.TaskUUID, jobData.Name)
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.NoAvailableResourcesError))
		return
	}

	var hostName string
	if jobData.JobResultURI != "" {
		resp, err := http.Get(jobData.JobResultURI)
		if err != nil {
			logs.GetLogger().Errorf("error making request to Space API: %+v", err)
			c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.ServerError, err.Error()))
			return
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				logs.GetLogger().Errorf("error closed resp Space API: %+v", err)
			}
		}(resp.Body)
		logs.GetLogger().Infof("Space API response received. Response: %d", resp.StatusCode)
		if resp.StatusCode != http.StatusOK {
			logs.GetLogger().Errorf("space API response not OK. Status Code: %d", resp.StatusCode)
			c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.ServerError, err.Error()))
		}

		var hostInfo struct {
			JobResultUri string `json:"job_result_uri"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&hostInfo); err != nil {
			logs.GetLogger().Errorf("error decoding Space API response JSON: %v", err)
			c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.ServerError, err.Error()))
			return
		}
		hostName = strings.ReplaceAll(hostInfo.JobResultUri, "https://", "")
	} else {
		hostName = generateString(10) + conf.GetConfig().API.Domain
	}
	jobData.JobRealUri = fmt.Sprintf("https://%s", hostName)

	go func() {
		spaceUuid := jobData.JobSourceURI[strings.LastIndex(jobData.JobSourceURI, "/")+1:]
		job, err := NewJobService().GetJobEntityBySpaceUuid(spaceUuid)
		if err != nil {
			logs.GetLogger().Errorf("get job failed, error: %+v", err)
			return
		}
		if job.SpaceUuid != "" {
			NewJobService().DeleteJobEntityBySpaceUuId(spaceUuid)
		}

		var jobEntity = new(models.JobEntity)
		jobEntity.Source = jobData.StorageSource
		jobEntity.SpaceUuid = spaceUuid
		jobEntity.TaskUuid = jobData.TaskUUID
		jobEntity.SourceUrl = jobData.JobSourceURI
		jobEntity.RealUrl = jobData.JobRealUri
		jobEntity.BuildLog = jobData.BuildLog
		jobEntity.ContainerLog = jobData.ContainerLog
		jobEntity.Duration = jobData.Duration
		jobEntity.DeployStatus = models.DEPLOY_RECEIVE_JOB
		jobEntity.CreateTime = time.Now().Unix()
		NewJobService().SaveJobEntity(jobEntity)

		go func() {
			if err = submitJob(&jobData); err != nil {
				logs.GetLogger().Errorf("upload job data to MCS failed, jobUuid: %s, spaceUuid: %s, error: %v",
					jobData.UUID, spaceUuid, err)
				return
			}
			logs.GetLogger().Infof("jobuuid: %s successfully uploaded to MCS", jobData.UUID)
		}()

		DeploySpaceTask(jobData.JobSourceURI, hostName, jobData.Duration, jobData.UUID, jobData.TaskUUID, gpuProductName)
	}()

	c.JSON(http.StatusOK, util.CreateSuccessResponse(jobData))
}

func ReNewJob(c *gin.Context) {
	var jobData struct {
		TaskUuid string `json:"task_uuid"`
		Duration int    `json:"duration"`
	}

	if err := c.ShouldBindJSON(&jobData); err != nil {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.JsonError))
		return
	}
	logs.GetLogger().Infof("renew Job received: %+v", jobData)

	if strings.TrimSpace(jobData.TaskUuid) == "" {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.BadParamError, "missing required field: task_uuid"))
		return
	}

	if jobData.Duration == 0 {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.BadParamError, "missing required field: duration"))
		return
	}

	jobEntity, err := NewJobService().GetJobEntityByTaskUuid(jobData.TaskUuid)
	if err != nil {
		logs.GetLogger().Errorf("Failed get job from db, taskUuid: %s, error: %+v", jobData.TaskUuid, err)
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.FoundJobEntityError))
		return
	}

	leftTime := jobEntity.ExpireTime - time.Now().Unix()
	if leftTime < 0 {
		c.JSON(http.StatusOK, util.CreateErrorResponse(util.BadParamError, "The job was terminated due to its expiration date"))
		return
	} else {
		jobEntity.ExpireTime = time.Now().Unix() + leftTime + int64(jobData.Duration)
		err = NewJobService().SaveJobEntity(&jobEntity)
		if err != nil {
			logs.GetLogger().Errorf("update job expireTime failed, taskUuid: %s, error: %+v", jobData.TaskUuid, err)
			c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.SaveJobEntityError))
			return
		}
	}
	c.JSON(http.StatusOK, util.CreateSuccessResponse("success"))
}

func CancelJob(c *gin.Context) {
	taskUuid := c.Query("task_uuid")
	if taskUuid == "" {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.BadParamError, "task_uuid is required"))
		return
	}

	nodeIdAndTaskUuidSignature := c.Query("node_id_task_uuid_signature")
	if len(nodeIdAndTaskUuidSignature) == 0 {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.SignatureError, "missing node_id_task_uuid_signature field"))
		return
	}

	if conf.GetConfig().HUB.VerifySign {

		cpRepoPath, _ := os.LookupEnv("CP_PATH")
		nodeID := GetNodeId(cpRepoPath)

		cpAccountAddress, err := contract.GetCpAccountAddress()
		if err != nil {
			logs.GetLogger().Errorf("get cp account contract address failed, error: %v", err)
			c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.GetCpAccountError))
			return
		}

		signature, err := verifySignatureForHub(conf.GetConfig().HUB.OrchestratorPk, fmt.Sprintf("%s%s%s", cpAccountAddress, nodeID, taskUuid), nodeIdAndTaskUuidSignature)
		if err != nil {
			logs.GetLogger().Errorf("verifySignature for space job failed, error: %+v", err)
			c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.ServerError, "verify sign data failed"))
			return
		}

		if !signature {
			logs.GetLogger().Errorf("space job sign verifing, task_id: %s,  verify: %v", taskUuid, signature)
			c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.SignatureError, "signature verify failed"))
			return
		}
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
		deleteJob(k8sNameSpace, jobEntity.SpaceUuid)
		NewJobService().DeleteJobEntityBySpaceUuId(jobEntity.SpaceUuid)
	}()

	c.JSON(http.StatusOK, util.CreateSuccessResponse("deleted success"))
}

func WhiteList(c *gin.Context) {
	walletWhiteListUrl := conf.GetConfig().API.WalletWhiteList
	list, err := getWalletList(walletWhiteListUrl)
	if err != nil {
		logs.GetLogger().Errorf("Failed get whiteList, error: %+v", err)
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.FoundWhiteListError))
		return
	}
	c.JSON(http.StatusOK, util.CreateSuccessResponse(list))
}

func BlackList(c *gin.Context) {
	walletBlackListUrl := conf.GetConfig().API.WalletBlackList
	list, err := getWalletList(walletBlackListUrl)
	if err != nil {
		logs.GetLogger().Errorf("Failed get blackList, error: %+v", err)
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.FoundBlackListError))
		return
	}
	c.JSON(http.StatusOK, util.CreateSuccessResponse(list))
}

func GetJobStatus(c *gin.Context) {
	jobUuId := c.Param("job_uuid")
	if jobUuId == "" {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.BadParamError, "missing required field: job_uuid"))
		return
	}

	signatureMsg := c.Query("signature")
	if signatureMsg == "" {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.BadParamError, "missing required field: signature"))
		return
	}

	cpAccountAddress, err := contract.GetCpAccountAddress()
	if err != nil {
		logs.GetLogger().Errorf("get cp account contract address failed, error: %v", err)
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.GetCpAccountError))
		return
	}

	logs.GetLogger().Infof("job_uuid: %s, signatureMsg: %s", jobUuId, signatureMsg)
	cpRepoPath, _ := os.LookupEnv("CP_PATH")
	nodeID := GetNodeId(cpRepoPath)
	signature, err := verifySignatureForHub(conf.GetConfig().HUB.OrchestratorPk, fmt.Sprintf("%s%s%s", cpAccountAddress, nodeID, jobUuId), signatureMsg)
	if err != nil {
		logs.GetLogger().Errorf("verifySignature for space job failed, error: %+v", err)
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.SignatureError, "verify sign data occur error"))
		return
	}

	if !signature {
		logs.GetLogger().Errorf("get job status sign verifing, jobUuid: %s, verify: %t", jobUuId, signature)
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.SignatureError))
		return
	}

	jobEntity, err := NewJobService().GetJobEntityByJobUuid(jobUuId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.FoundJobEntityError))
		return
	}

	if jobEntity.JobUuid == "" {
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.NotFoundJobEntityError))
		return
	}

	var jobResult struct {
		JobUuid      string `json:"job_uuid"`
		JobStatus    string `json:"job_status"`
		JobResultUrl string `json:"job_result_url"`
	}
	jobResult.JobUuid = jobEntity.JobUuid
	jobResult.JobStatus = models.GetDeployStatusStr(jobEntity.DeployStatus)
	jobResult.JobResultUrl = jobEntity.ResultUrl

	c.JSON(http.StatusOK, util.CreateSuccessResponse(jobResult))
}

func StatisticalSources(c *gin.Context) {
	location, err := getLocation()
	if err != nil {
		logs.GetLogger().Error(err)
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.GetLocationError))
		return
	}

	k8sService := NewK8sService()
	statisticalSources, err := k8sService.StatisticalSources(context.TODO())
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.GeResourceError))
		return
	}

	cpAccountAddress, err := contract.GetCpAccountAddress()
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.GetCpAccountError))
		return
	}

	cpRepo, _ := os.LookupEnv("CP_PATH")
	c.JSON(http.StatusOK, models.ClusterResource{
		Region:           location,
		ClusterInfo:      statisticalSources,
		NodeName:         conf.GetConfig().API.NodeName,
		NodeId:           GetNodeId(cpRepo),
		CpAccountAddress: cpAccountAddress,
	})
}

func GetSpaceLog(c *gin.Context) {
	spaceUuid := c.Query("space_id")
	logType := c.Query("type")
	orderType := c.Query("order")
	if strings.TrimSpace(spaceUuid) == "" {
		logs.GetLogger().Errorf("get space log failed, space_id is empty: %s", spaceUuid)
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.BadParamError, "missing required field: space_id"))
		return
	}

	if strings.TrimSpace(logType) == "" {
		logs.GetLogger().Errorf("get space log failed, type is empty: %s", logType)
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.BadParamError, "missing required field: type"))
		return
	}

	if strings.TrimSpace(logType) != "build" && strings.TrimSpace(logType) != "container" {
		logs.GetLogger().Errorf("get space log failed, type is build or container, type:: %s", logType)
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.BadParamError, "missing required field: type"))
		return
	}

	jobEntity, err := NewJobService().GetJobEntityBySpaceUuid(spaceUuid)
	if err != nil {
		logs.GetLogger().Error(err)
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.FoundJobEntityError))
		return
	}

	conn, err := upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logs.GetLogger().Errorf("upgrading connection failed, error: %+v", err)
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.ServerError, "websocket upgrading connection failed"))
		return
	}

	if orderType == "private" {
		handlePodEvent(conn, jobEntity.SpaceUuid, jobEntity.WalletAddress)
	} else {
		handleConnection(conn, jobEntity, logType)
	}
}

func DoProof(c *gin.Context) {
	var proofTask struct {
		Method    string `json:"method"`
		BlockData string `json:"block_data"`
		Exp       int64  `json:"exp"`
	}
	if err := c.ShouldBindJSON(&proofTask); err != nil {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.JsonError))
		return
	}
	logs.GetLogger().Infof("do proof task received: %+v", proofTask)

	if strings.TrimSpace(proofTask.Method) == "" {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.ProofParamError, "missing required field: method"))
		return
	}
	if proofTask.Method != "mine" {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.ProofParamError, "method must be mine"))
		return
	}
	if proofTask.Exp < 0 || proofTask.Exp > 250 {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.ProofParamError, "exp range is [0~250]"))
		return
	}

	k8sService := NewK8sService()
	job := &batchv1.Job{
		ObjectMeta: metaV1.ObjectMeta{
			Name: "proof-job-" + generateString(5),
		},
		Spec: batchv1.JobSpec{
			Template: v1.PodTemplateSpec{
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:  "worker-container-" + generateString(5),
							Image: "filswan/worker-proof:v1.0",
							Env: []v1.EnvVar{
								{
									Name:  "METHOD",
									Value: proofTask.Method,
								},
								{
									Name:  "BLOCK_DATA",
									Value: proofTask.BlockData,
								},
								{
									Name:  "EXP",
									Value: strconv.Itoa(int(proofTask.Exp)),
								},
							},
						},
					},
					RestartPolicy: "Never",
				},
			},
			BackoffLimit:            new(int32),
			TTLSecondsAfterFinished: new(int32),
		},
	}

	*job.Spec.BackoffLimit = 1
	*job.Spec.TTLSecondsAfterFinished = 30

	createdJob, err := k8sService.k8sClient.BatchV1().Jobs(metaV1.NamespaceDefault).Create(context.TODO(), job, metaV1.CreateOptions{})
	if err != nil {
		logs.GetLogger().Errorf("Failed creating Pod: %v", err)
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.ProofError))
		return
	}

	err = wait.PollImmediate(time.Second*3, time.Minute*5, func() (bool, error) {
		job, err := k8sService.k8sClient.BatchV1().Jobs(metaV1.NamespaceDefault).Get(context.Background(), createdJob.Name, metaV1.GetOptions{})
		if err != nil {
			logs.GetLogger().Errorf("Failed getting Job status: %v\n", err)
			return false, err
		}

		if job.Status.Succeeded > 0 {
			return true, nil
		}

		return false, nil
	})
	if err != nil {
		logs.GetLogger().Errorf("Failed waiting for Job completion: %v", err)
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.ProofError))
		return
	}

	podList, err := k8sService.k8sClient.CoreV1().Pods(metaV1.NamespaceDefault).List(context.Background(), metaV1.ListOptions{
		LabelSelector: fmt.Sprintf("job-name=%s", createdJob.Name),
	})
	if err != nil {
		logs.GetLogger().Errorf("Error getting Pods for Job: %v\n", err)
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.ProofError))
		return
	}

	if len(podList.Items) == 0 {
		logs.GetLogger().Errorf("No Pods found for Job.")
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.ProofError))
		return
	}

	podName := podList.Items[0].Name
	podLog, err := k8sService.k8sClient.CoreV1().Pods(metaV1.NamespaceDefault).GetLogs(podName, &v1.PodLogOptions{}).Stream(context.Background())
	if err != nil {
		logs.GetLogger().Errorf("Failed gettingPod logs: %v", err)
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.ProofReadLogError))
		return
	}
	defer podLog.Close()

	bytes, err := io.ReadAll(podLog)
	if err != nil {
		logs.GetLogger().Errorf("Failed gettingPod logs: %v", err)
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.ProofReadLogError))
		return
	}
	c.JSON(http.StatusOK, util.CreateSuccessResponse(string(bytes)))
}

func handlePodEvent(conn *websocket.Conn, spaceUuid string, walletAddress string) {
	client := NewWsClient(conn)

	k8sNameSpace := constants.K8S_NAMESPACE_NAME_PREFIX + strings.ToLower(walletAddress)
	k8sService := NewK8sService()
	events, err := k8sService.k8sClient.CoreV1().Events(k8sNameSpace).List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		logs.GetLogger().Errorf("get pod events failed, error: %v", err)
		return
	}

	var buffer strings.Builder
	for _, event := range events.Items {
		if strings.Contains(event.InvolvedObject.Name, spaceUuid) {
			buffer.WriteString(event.Message)
			buffer.WriteString("\n")
		}
	}
	client.HandleLogs(strings.NewReader(buffer.String()))

}

func handleConnection(conn *websocket.Conn, jobDetail models.JobEntity, logType string) {
	client := NewWsClient(conn)

	if logType == "build" {
		cpRepoPath, _ := os.LookupEnv("CP_PATH")
		buildLogPath := filepath.Join(cpRepoPath, "build", jobDetail.WalletAddress, "spaces", jobDetail.Name, BuildFileName)
		if _, err := os.Stat(buildLogPath); err != nil {
			client.HandleLogs(strings.NewReader("This space is deployed starting from a image."))
		} else {
			logFile, _ := os.Open(buildLogPath)
			defer logFile.Close()
			client.HandleLogs(logFile)
		}
	} else if logType == "container" {
		k8sNameSpace := constants.K8S_NAMESPACE_NAME_PREFIX + strings.ToLower(jobDetail.WalletAddress)

		k8sService := NewK8sService()
		pods, err := k8sService.k8sClient.CoreV1().Pods(k8sNameSpace).List(context.TODO(), metaV1.ListOptions{
			LabelSelector: fmt.Sprintf("lad_app=%s", jobDetail.SpaceUuid),
		})
		if err != nil {
			logs.GetLogger().Errorf("Error listing Pods: %v", err)
			return
		}

		if len(pods.Items) > 0 {
			line := int64(1000)
			containerStatuses := pods.Items[0].Status.ContainerStatuses
			lastIndex := len(containerStatuses) - 1
			req := k8sService.k8sClient.CoreV1().Pods(k8sNameSpace).GetLogs(pods.Items[0].Name, &v1.PodLogOptions{
				Container:  containerStatuses[lastIndex].Name,
				Follow:     true,
				Timestamps: true,
				TailLines:  &line,
			})

			podLogs, err := req.Stream(context.Background())
			if err != nil {
				logs.GetLogger().Errorf("Error opening log stream: %v", err)
				return
			}
			defer podLogs.Close()

			client.HandleLogs(podLogs)
		}
	}
}

func DeploySpaceTask(jobSourceURI, hostName string, duration int, jobUuid string, taskUuid string, gpuProductName string) string {
	updateJobStatus(jobUuid, models.DEPLOY_UPLOAD_RESULT)

	var success bool
	var spaceUuid string
	var walletAddress string
	defer func() {
		if !success {
			k8sNameSpace := constants.K8S_NAMESPACE_NAME_PREFIX + strings.ToLower(walletAddress)
			deleteJob(k8sNameSpace, spaceUuid)
			NewJobService().DeleteJobEntityBySpaceUuId(spaceUuid)
		}

		if err := recover(); err != nil {
			logs.GetLogger().Errorf("deploy space task painc, error: %+v", err)
			return
		}
	}()

	spaceDetail, err := getSpaceDetail(jobSourceURI)
	if err != nil {
		logs.GetLogger().Errorln(err)
		return ""
	}

	walletAddress = spaceDetail.Data.Owner.PublicAddress
	spaceName := spaceDetail.Data.Space.Name
	spaceUuid = strings.ToLower(spaceDetail.Data.Space.Uuid)
	spaceHardware := spaceDetail.Data.Space.ActiveOrder.Config

	logs.GetLogger().Infof("uuid: %s, spaceName: %s, hardwareName: %s", spaceUuid, spaceName, spaceHardware.Description)
	if len(spaceHardware.Description) == 0 {
		return ""
	}

	var job = new(models.JobEntity)
	job.WalletAddress = walletAddress
	job.Name = spaceName
	job.SpaceUuid = spaceDetail.Data.Space.Uuid
	job.Hardware = spaceHardware.Description
	job.SpaceType = 0
	if err = NewJobService().UpdateJobEntityBySpaceUuid(job); err != nil {
		logs.GetLogger().Errorf("update job info failed, error: %v", err)
		return ""
	}

	deploy := NewDeploy(jobUuid, hostName, walletAddress, spaceHardware.Description, int64(duration), taskUuid, constants.SPACE_TYPE_PUBLIC)
	deploy.WithSpaceInfo(spaceUuid, spaceName)
	deploy.WithGpuProductName(gpuProductName)

	cpRepoPath, _ := os.LookupEnv("CP_PATH")
	spacePath := filepath.Join(cpRepoPath, "build", walletAddress, "spaces", spaceName)
	os.RemoveAll(spacePath)
	updateJobStatus(jobUuid, models.DEPLOY_DOWNLOAD_SOURCE)
	containsYaml, yamlPath, imagePath, modelsSettingFile, _, err := BuildSpaceTaskImage(spaceUuid, spaceDetail.Data.Files)
	if err != nil {
		logs.GetLogger().Error(err)
		return ""
	}

	deploy.WithSpacePath(imagePath)
	if len(modelsSettingFile) > 0 {
		err := deploy.WithModelSettingFile(modelsSettingFile).ModelInferenceToK8s()
		if err != nil {
			logs.GetLogger().Error(err)
			return ""
		}
		return hostName
	}

	if containsYaml {
		deploy.WithYamlInfo(yamlPath).YamlToK8s()
	} else {
		imageName, dockerfilePath := BuildImagesByDockerfile(jobUuid, spaceUuid, spaceName, imagePath)
		deploy.WithDockerfile(imageName, dockerfilePath).DockerfileToK8s()
	}
	success = true

	return hostName
}

func deleteJob(namespace, spaceUuid string) error {
	deployName := constants.K8S_DEPLOY_NAME_PREFIX + spaceUuid
	serviceName := constants.K8S_SERVICE_NAME_PREFIX + spaceUuid
	ingressName := constants.K8S_INGRESS_NAME_PREFIX + spaceUuid

	logs.GetLogger().Infof("deleting space service, space_uuid: %s", spaceUuid)
	k8sService := NewK8sService()
	if err := k8sService.DeleteIngress(context.TODO(), namespace, ingressName); err != nil && !errors.IsNotFound(err) {
		logs.GetLogger().Errorf("Failed delete ingress, ingressName: %s, error: %+v", ingressName, err)
		return err
	}

	if err := k8sService.DeleteService(context.TODO(), namespace, serviceName); err != nil && !errors.IsNotFound(err) {
		logs.GetLogger().Errorf("Failed delete service, serviceName: %s, error: %+v", serviceName, err)
		return err
	}

	dockerService := NewDockerService()
	deployImageIds, err := k8sService.GetDeploymentImages(context.TODO(), namespace, deployName)
	if err != nil && !errors.IsNotFound(err) {
		logs.GetLogger().Errorf("Failed get deploy imageIds, deployName: %s, error: %+v", deployName, err)
		return err
	}
	for _, imageId := range deployImageIds {
		dockerService.RemoveImage(imageId)
	}

	if err := k8sService.DeleteDeployment(context.TODO(), namespace, deployName); err != nil && !errors.IsNotFound(err) {
		logs.GetLogger().Errorf("Failed delete deployment, deployName: %s, error: %+v", deployName, err)
		return err
	}
	time.Sleep(6 * time.Second)

	if err := k8sService.DeleteDeployRs(context.TODO(), namespace, spaceUuid); err != nil && !errors.IsNotFound(err) {
		logs.GetLogger().Errorf("Failed delete ReplicaSetsController, spaceUuid: %s, error: %+v", spaceUuid, err)
		return err
	}

	if err := k8sService.DeletePod(context.TODO(), namespace, spaceUuid); err != nil && !errors.IsNotFound(err) {
		logs.GetLogger().Errorf("Failed delete pods, spaceUuid: %s, error: %+v", spaceUuid, err)
		return err
	}

	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()
	var count = 0
	for {
		<-ticker.C
		count++
		if count >= 20 {
			break
		}
		getPods, err := k8sService.GetPods(namespace, spaceUuid)
		if err != nil && !errors.IsNotFound(err) {
			logs.GetLogger().Errorf("Failed get pods form namespace, namepace: %s, error: %+v", namespace, err)
			continue
		}
		if !getPods {
			break
		}
	}
	return nil
}

func downloadModelUrl(namespace, spaceUuid, serviceIp string, podCmd []string) {
	k8sService := NewK8sService()
	podName, err := k8sService.WaitForPodRunningByHttp(namespace, spaceUuid, serviceIp)
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	if err = k8sService.PodDoCommand(namespace, podName, "", podCmd); err != nil {
		logs.GetLogger().Error(err)
		return
	}
}

func updateJobStatus(jobUuid string, jobStatus int, url ...string) {
	go func() {
		if len(url) > 0 {
			deployingChan <- models.Job{
				Uuid:   jobUuid,
				Status: jobStatus,
				Url:    url[0],
			}
		} else {
			deployingChan <- models.Job{
				Uuid:   jobUuid,
				Status: jobStatus,
				Url:    "",
			}
		}
	}()
}

func getSpaceDetail(jobSourceURI string) (models.SpaceJSON, error) {
	resp, err := http.Get(jobSourceURI)
	if err != nil {
		return models.SpaceJSON{}, fmt.Errorf("error making request to Space API: %+v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return models.SpaceJSON{}, fmt.Errorf("space API response not OK. Status Code: %d", resp.StatusCode)
	}

	var spaceJson models.SpaceJSON
	if err := json.NewDecoder(resp.Body).Decode(&spaceJson); err != nil {
		return models.SpaceJSON{}, fmt.Errorf("error decoding Space API response JSON: %v", err)
	}
	return spaceJson, nil
}

func checkResourceAvailableForSpace(configDescription string) (bool, string, error) {
	taskType, hardwareDetail := getHardwareDetail(configDescription)
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
		logs.GetLogger().Infof("checkResourceAvailableForSpace: needCpu: %d, needMemory: %.2f, needStorage: %.2f", needCpu, needMemory, needStorage)
		logs.GetLogger().Infof("checkResourceAvailableForSpace: remainingCpu: %d, remainingMemory: %.2f, remainingStorage: %.2f", remainderCpu, remainderMemory, remainderStorage)
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

func checkResourceAvailableForUbi(taskType int, gpuName string, resource *models.TaskResource) (string, string, int64, int64, int64, error) {
	k8sService := NewK8sService()
	activePods, err := k8sService.GetAllActivePod(context.TODO())
	if err != nil {
		return "", "", 0, 0, 0, err
	}

	nodes, err := k8sService.k8sClient.CoreV1().Nodes().List(context.TODO(), metaV1.ListOptions{})
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

		nodeGpu, remainderResource, _ := GetNodeResource(activePods, &node)
		remainderCpu := remainderResource[ResourceCpu]
		remainderMemory := float64(remainderResource[ResourceMem] / 1024 / 1024 / 1024)
		remainderStorage := float64(remainderResource[ResourceStorage] / 1024 / 1024 / 1024)

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
				logs.GetLogger().Infof("needGpuName: %s, nodeGpu: %+v, nodeGpuSummary: %+v", gpuName, nodeGpu, nodeGpuSummary)
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

var regionCache string

func getLocation() (string, error) {
	var err error
	if regionCache != "" {
		return regionCache, nil
	}
	regionCache, err = getRegionByIpInfo()
	if err != nil {
		regionCache, err = getRegionByIpApi()
		if err != nil {
			logs.GetLogger().Errorf("get region info failed, error: %+v", err)
			return "", err
		}
	}
	return regionCache, nil
}

func getRegionByIpApi() (string, error) {
	req, err := http.NewRequest("GET", "https://ipapi.co/ip", nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36")

	client := http.DefaultClient
	IpResp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer IpResp.Body.Close()

	ipBytes, err := io.ReadAll(IpResp.Body)
	if err != nil {
		return "", err
	}

	regionResp, err := http.Get("http://ip-api.com/json/" + string(ipBytes))
	if err != nil {
		return "", err
	}
	defer regionResp.Body.Close()

	body, err := io.ReadAll(regionResp.Body)
	if err != nil {
		return "", err
	}

	var ipInfo struct {
		Country     string `json:"country"`
		CountryCode string `json:"countryCode"`
		City        string `json:"city"`
		Region      string `json:"region"`
		RegionName  string `json:"regionName"`
	}
	if err = json.Unmarshal(body, &ipInfo); err != nil {
		return "", err
	}
	region := strings.TrimSpace(ipInfo.RegionName) + "-" + ipInfo.CountryCode
	return region, nil
}

func getRegionByIpInfo() (string, error) {
	req, err := http.NewRequest("GET", "https://ipinfo.io", nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36")

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	ipBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var ipInfo struct {
		Ip      string `json:"ip"`
		City    string `json:"city"`
		Region  string `json:"region"`
		Country string `json:"country"`
	}
	if err = json.Unmarshal(ipBytes, &ipInfo); err != nil {
		return "", err
	}
	region := strings.TrimSpace(ipInfo.Region) + "-" + ipInfo.Country
	return region, nil
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

func verifySignatureForHub(pubKStr string, message string, signedMessage string) (bool, error) {
	hashedMessage := []byte("\x19Ethereum Signed Message:\n" + strconv.Itoa(len(message)) + message)
	hash := crypto.Keccak256Hash(hashedMessage)

	decodedMessage, err := hexutil.Decode(signedMessage)
	if err != nil {
		return false, err
	}

	if decodedMessage[64] == 27 || decodedMessage[64] == 28 {
		decodedMessage[64] -= 27
	}

	sigPublicKeyECDSA, err := crypto.SigToPub(hash.Bytes(), decodedMessage)
	if sigPublicKeyECDSA == nil {
		err = fmt.Errorf("could not get a public get from the message signature")
	}
	if err != nil {
		return false, err
	}

	return pubKStr == crypto.PubkeyToAddress(*sigPublicKeyECDSA).String(), nil
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
	return strings.ToUpper(name)
}

func getWalletList(walletUrl string) ([]string, error) {
	if walletUrl == "" {
		return nil, nil
	}

	var walletMap = make(map[string]struct{})
	resp, err := http.Get(walletUrl)
	if err != nil {
		logs.GetLogger().Errorf("send wallet url: %s, failed, error: %v", walletUrl, err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logs.GetLogger().Errorf("response wallet whitelist failed, error: %v", err)
		return nil, err
	}

	var addressList []string
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "#") {
			continue
		}
		walletAddress := scanner.Text()
		if strings.TrimSpace(walletAddress) != "" {
			walletMap[walletAddress] = struct{}{}
		}

		addressList = append(addressList, walletAddress)
	}

	if err := scanner.Err(); err != nil {
		logs.GetLogger().Errorf("read response of wallet whitelist failed, error: %v", err)
		return nil, err
	}
	return addressList, nil
}

func CheckWalletWhiteList(jobSourceURI string) bool {
	walletWhiteListUrl := conf.GetConfig().API.WalletWhiteList
	if walletWhiteListUrl == "" {
		return true
	}
	whiteList, err := getWalletList(walletWhiteListUrl)
	if err != nil {
		logs.GetLogger().Errorf("get whiteList By url failed, url: %s, error: %v", err)
		return false
	}

	spaceDetail, err := getSpaceDetail(jobSourceURI)
	if err != nil {
		logs.GetLogger().Errorln(err)
		return false
	}
	userWalletAddress := spaceDetail.Data.Owner.PublicAddress

	for _, address := range whiteList {
		if userWalletAddress == address {
			return true
		}
	}
	return false
}

func CheckWalletBlackList(jobSourceURI string) bool {
	walletBlackListUrl := conf.GetConfig().API.WalletBlackList
	if walletBlackListUrl == "" {
		return false
	}
	blackList, err := getWalletList(walletBlackListUrl)
	if err != nil {
		logs.GetLogger().Errorf("get blacklist By url failed, url: %s, error: %v", err)
		return false
	}

	spaceDetail, err := getSpaceDetail(jobSourceURI)
	if err != nil {
		logs.GetLogger().Errorln(err)
		return false
	}
	userWalletAddress := spaceDetail.Data.Owner.PublicAddress

	for _, address := range blackList {
		if userWalletAddress == address {
			return true
		}
	}
	return false
}

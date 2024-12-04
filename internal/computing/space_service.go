package computing

import (
	"bufio"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/swanchain/go-computing-provider/internal/contract/account"
	"github.com/swanchain/go-computing-provider/internal/yaml"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

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
	"github.com/swanchain/go-computing-provider/internal/contract/fcp"
	"github.com/swanchain/go-computing-provider/internal/models"
	"github.com/swanchain/go-computing-provider/util"
	batchv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
)

var gpuResourceCache sync.Map

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
			logs.GetLogger().Errorf("failed to get cp account contract address, error: %v", err)
			c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.GetCpAccountError))
			return
		}

		signature, err := verifySignatureForHub(conf.GetConfig().HUB.OrchestratorPk, fmt.Sprintf("%s%s%s", cpAccountAddress, nodeID, jobData.JobSourceURI), jobData.NodeIdJobSourceUriSignature)
		if err != nil {
			logs.GetLogger().Errorf("failed to verify signature for space job, error: %+v", err)
			c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.SignatureError, "verify sign data occur error"))
			return
		}

		if !signature {
			logs.GetLogger().Errorf("space job sign verifing, job_uuid: %s, verify: %v", jobData.UUID, signature)
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

	if jobData.JobType == 1 {
		if !conf.GetConfig().API.Pricing {
			checkPriceFlag, totalCost, err := checkPrice(jobData.BidPrice, jobData.Duration, spaceDetail.Data.Space.ActiveOrder.Config)
			if err != nil {
				logs.GetLogger().Errorf("failed to check price, job_uuid: %s, error: %v", jobData.UUID, err)
				c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.JsonError))
				return
			}

			if !checkPriceFlag {
				logs.GetLogger().Warnf("the price is too low, job_uuid: %s, paid: %s, required: %0.4f", jobData.UUID, jobData.BidPrice, totalCost)
				c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.BelowPriceError))
				return
			}
		}
	}

	available, gpuProductName, gpuIndex, err := checkResourceAvailableForSpace(jobData.JobType, spaceDetail.Data.Space.ActiveOrder.Config)
	if err != nil {
		logs.GetLogger().Errorf("failed to check job resource, error: %+v", err)
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.CheckResourcesError))
		return
	}

	if !available {
		if gpuProductName != "" {
			logs.GetLogger().Warnf("job_uuid: %s, name: %s, gpu_name: %s, not found a resources available", jobData.UUID, jobData.Name, gpuProductName)
		} else {
			logs.GetLogger().Warnf("job_uuid: %s, name: %s, not found a resources available", jobData.UUID, jobData.Name)
		}
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
	spaceUuid := spaceDetail.Data.Space.Uuid
	wsUrl := fmt.Sprintf("wss://%s:%s/api/v1/computing/lagrange/spaces/log?job_uuid=%s", logHost, multiAddressSplit[4], jobData.UUID)
	jobData.BuildLog = wsUrl + "&type=build"
	jobData.ContainerLog = wsUrl + "&type=container"
	jobData.JobRealUri = fmt.Sprintf("https://%s", hostName)
	jobData.NodeIdJobSourceUriSignature = ""

	deployParam, err := DownloadSpaceResources(jobData.UUID, spaceDetail.Data.Files)
	if err != nil {
		logs.GetLogger().Errorf("failed to download space resource, job_uuid: %s, error: %v", jobData.UUID, err)
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.DownloadResourceError))
		return
	}

	var serviceNodePort int32
	if deployParam.ContainsYaml {
		containerResources, err := yaml.HandlerYaml(deployParam.YamlFilePath)
		if err != nil {
			logs.GetLogger().Errorf("failed to parse yaml, job_uuid: %s, error: %v", jobData.UUID, err)
			c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.DownloadResourceError))
			return
		}

		if len(containerResources) == 1 && containerResources[0].ServiceType == yaml.ServiceTypeNodePort {
			chainRpc, err := conf.GetRpcByNetWorkName()
			if err != nil {
				logs.GetLogger().Errorf("failed to get rpc, job_uuid: %s, error: %v", jobData.UUID, err)
				c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.RpcConnectError))
				return
			}
			client, err := contract.GetEthClient(chainRpc)
			if err != nil {
				logs.GetLogger().Errorf("failed to connect rpc, job_uuid: %s, error: %v", jobData.UUID, err)
				c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.RpcConnectError))
				return
			}
			defer client.Close()

			cpStub, err := account.NewAccountStub(client)
			if err != nil {
				logs.GetLogger().Errorf("job_uuid: %s, error: %v", jobData.UUID, err)
				c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.RpcConnectError))
				return
			}
			cpAccount, err := cpStub.GetCpAccountInfo()
			if err != nil {
				logs.GetLogger().Errorf("job_uuid: %s, error: %v", jobData.UUID, err)
				c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.RpcConnectError))
				return
			}

			var sshTaskFlag bool
			for _, taskType := range cpAccount.TaskTypes {
				if taskType == 5 {
					sshTaskFlag = true
					break
				}
			}

			if !sshTaskFlag {
				c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.NotAcceptNodePortError))
				return
			}

			_, serviceNodePort, err = NewK8sService().CheckServiceNodePort(0)
			if err != nil {
				logs.GetLogger().Errorf("failed to check port, error: %v", err)
				c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.PortNoAvailableError))
				return
			}

			realUrl := fmt.Sprintf("ssh root@%s -p%d", multiAddressSplit[2], serviceNodePort)
			jobData.JobRealUri = realUrl
			jobData.ContainerLog = jobData.ContainerLog + "&order=private"
			logs.GetLogger().Infof("job_uuid: %s, real url: %s", jobData.UUID, realUrl)
		}
	}

	go func() {
		var currentBlockNumber uint64
		for i := 0; i < 5; i++ {
			currentBlockNumber, err = getChainBlockNumber()
			if err != nil {
				logs.GetLogger().Errorf("failed to get blockNumber, error: %v", err)
				time.Sleep(time.Second)
				continue
			}
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
		jobEntity.ExpireTime = time.Now().Unix() + int64(jobData.Duration)
		jobEntity.StartedBlock = conf.GetConfig().CONTRACT.JobManagerCreated
		jobEntity.ScannedBlock = currentBlockNumber
		jobEntity.WalletAddress = spaceDetail.Data.Owner.PublicAddress
		jobEntity.Name = spaceDetail.Data.Space.Name
		jobEntity.Hardware = spaceDetail.Data.Space.ActiveOrder.Config.Description
		jobEntity.SpaceType = 0
		jobEntity.ResourceType = spaceDetail.Data.Space.ActiveOrder.Config.HardwareType
		jobEntity.NameSpace = constants.K8S_NAMESPACE_NAME_PREFIX + strings.ToLower(spaceDetail.Data.Owner.PublicAddress)
		jobEntity.K8sDeployName = constants.K8S_DEPLOY_NAME_PREFIX + strings.ToLower(jobData.UUID)
		jobEntity.Status = models.JOB_RECEIVED_STATUS
		jobEntity.K8sResourceType = "deployment"
		jobEntity.IpWhiteList = strings.Join(jobData.IpWhiteList, ",")
		err = NewJobService().SaveJobEntity(jobEntity)
		if err != nil {
			logs.GetLogger().Errorf("failed to save job to db, job_uuid: %s, error: %+v", jobData.UUID, err)
		}

		go func() {
			if err = submitJob(&jobData); err != nil {
				logs.GetLogger().Errorf("failed to upload job data to MCS, job_uuid: %s, spaceUuid: %s, error: %v", jobData.UUID, spaceUuid, err)
				return
			}
			logs.GetLogger().Infof("successfully uploaded to MCS, jobuuid: %s", jobData.UUID)
		}()

		DeploySpaceTask(jobData, deployParam, hostName, gpuProductName, serviceNodePort, jobData.JobType, jobData.IpWhiteList, gpuIndex)
	}()

	c.JSON(http.StatusOK, util.CreateSuccessResponse(jobData))
}

func getChainBlockNumber() (uint64, error) {
	var currentBlockNumber uint64
	chainUrl, err := conf.GetRpcByNetWorkName()
	if err != nil {
		return 0, err
	}
	client, err := contract.GetEthClient(chainUrl)
	if err != nil {
		return 0, err
	}
	defer client.Close()

	currentBlockNumber, err = client.BlockNumber(context.Background())
	if err != nil {
		return 0, err
	}
	return currentBlockNumber, nil
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
		return fmt.Errorf("failed to parse to json, error: %v", err)
	}

	f, err := os.OpenFile(taskDetailFilePath, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		logs.GetLogger().Errorf("failed to open file, error: %v", err)
		return err
	}
	defer f.Close()

	if _, err = f.Write(bytes); err != nil {
		return fmt.Errorf("write jobData to file failed, error: %v", err)
	}

	var resultMcsUrl string
	for i := 0; i < 5; i++ {
		storageService, err := NewStorageService()
		if err != nil {
			return err
		}
		mcsOssFile, err := storageService.UploadFileToBucket(jobDetailFile, taskDetailFilePath, true)
		if err != nil {
			logs.GetLogger().Errorf("failed to upload file to bucket, error: %v", err)
			continue
		}

		if mcsOssFile == nil || mcsOssFile.PayloadCid == "" {
			continue
		}

		gatewayUrl, err := storageService.GetGatewayUrl()
		if err != nil {
			logs.GetLogger().Errorf("failed to get mcs ipfs gatewayUrl, error: %v", err)
			continue
		}
		resultMcsUrl = *gatewayUrl + "/ipfs/" + mcsOssFile.PayloadCid
		break
	}
	return NewJobService().UpdateJobResultUrlByJobUuid(jobData.UUID, resultMcsUrl)
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
		logs.GetLogger().Errorf("failed get job from db, taskUuid: %s, error: %+v", jobData.TaskUuid, err)
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
			logs.GetLogger().Errorf("failed to update job expireTime, taskUuid: %s, error: %+v", jobData.TaskUuid, err)
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
		//Compatible with old versions
		DeleteJob(k8sNameSpace, jobEntity.SpaceUuid, "compatible with old versions, terminated job form hub")
		DeleteJob(k8sNameSpace, jobEntity.JobUuid, "terminated job form hub")
		NewJobService().DeleteJobEntityByJobUuId(jobEntity.JobUuid, models.JOB_TERMINATED_STATUS)
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

	if conf.GetConfig().HUB.VerifySign {
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

func GetPublicKey(c *gin.Context) {
	cpRepoPath, _ := os.LookupEnv("CP_PATH")
	publicKeyPath := filepath.Join(cpRepoPath, util.RSA_DIR_NAME, util.RSA_PUBLIC_KEY)
	privateKeyPath := filepath.Join(cpRepoPath, util.RSA_DIR_NAME, util.RSA_PRIVATE_KEY)

	var publicKeyData []byte
	_, err := os.Stat(publicKeyPath)
	if err != nil {
		os.Mkdir(filepath.Join(cpRepoPath, util.RSA_DIR_NAME), 0755)
		privateKey, publicKey, err := util.GenerateRSAKeyPair(2048)
		if err != nil {
			logs.GetLogger().Errorf("failed to generate rsa keyPair, error: %v", err)
			c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.GenerateRsaError))
			return
		}
		if err = util.SavePrivateEMKey(privateKeyPath, privateKey); err != nil {
			logs.GetLogger().Errorf("failed to save rsa private key, error: %v", err)
			c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.SaveRsaKeyError))
			return
		}
		if err = util.SavePublicPEMKey(publicKeyPath, publicKey); err != nil {
			logs.GetLogger().Errorf("failed to save rsa public key, error: %v", err)
			c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.SaveRsaKeyError))
			return
		}
	}

	publicKeyData, err = util.ReadPublicKey(publicKeyPath)
	if err != nil {
		logs.GetLogger().Errorf("failed to read rsa public key, error: %v", err)
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.ReadRsaKeyError))
		return
	}

	encodedData := base64.StdEncoding.EncodeToString(publicKeyData)
	c.JSON(http.StatusOK, util.CreateSuccessResponse(map[string]any{
		"public_key": encodedData,
	}))
}

func CheckNodeportServiceEnv(c *gin.Context) {
	CheckClusterNetworkPolicy()
	k8sService := NewK8sService()
	usedPorts, err := k8sService.GetUsedNodePorts()
	if err != nil {
		logs.GetLogger().Errorf("failed to get used node-port, error: %v", err)
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.CheckNodePortError))
		return
	}
	var msg string
	availability := util.CheckPortAvailability(usedPorts)
	if !availability {
		msg = "failed to check node port"
	}

	daemonSet, err := k8sService.k8sClient.AppsV1().DaemonSets("kube-system").Get(context.TODO(), "resource-limit", metaV1.GetOptions{})
	if err != nil {
		logs.GetLogger().Errorf("failed to get resource-limit daemonSet, error: %v", err)
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.CheckResourceLimitError))
		return
	}

	if daemonSet == nil {
		msg = "failed to check resource-limit"
	}

	if !NetworkPolicyFlag {
		msg = "failed to check network policy"
	}

	if msg == "" {
		c.JSON(http.StatusOK, util.CreateSuccessResponse("All checks have passed"))
	} else {
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(http.StatusInternalServerError, msg))
	}
}

func GetPrice(c *gin.Context) {
	readPriceConfig, err := ReadPriceConfig()
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.ReadPriceError))
		return
	}

	var resourcePriceResp models.ResourcePrice
	resourcePriceResp.CpuPrice = readPriceConfig.TARGET_CPU
	resourcePriceResp.MemoryPrice = readPriceConfig.TARGET_MEMORY
	resourcePriceResp.HdEphemeralPrice = readPriceConfig.TARGET_HD_EPHEMERAL
	resourcePriceResp.GpuDefaultPrice = readPriceConfig.TARGET_GPU_DEFAULT
	resourcePriceResp.GpusPrice = readPriceConfig.GpusPrice
	resourcePriceResp.Pricing = bool(conf.GetConfig().API.Pricing)

	c.JSON(http.StatusOK, util.CreateSuccessResponse(resourcePriceResp))
}

func StatisticalSources(c *gin.Context) {
	location, err := getLocation()
	if err != nil {
		logs.GetLogger().Error(err)
		location = "-"
	}

	k8sService := NewK8sService()
	statisticalSources, err := k8sService.StatisticalSources(context.TODO())
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.GeResourceError))
		return
	}

	clusterRuntime, err := k8sService.GetClusterRuntime()
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
		Runtime:          clusterRuntime,
		ClientType:       "FCP",
	})
}

func GetSpaceLog(c *gin.Context) {
	jobUuid := c.Query("job_uuid")
	logType := c.Query("type")
	orderType := c.Query("order")
	if strings.TrimSpace(jobUuid) == "" {
		logs.GetLogger().Errorf("get space log failed, job_uuid is empty: %s", jobUuid)
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.BadParamError, "missing required field: job_uuid"))
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

	jobEntity, err := NewJobService().GetJobEntityByJobUuid(jobUuid)
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
		handlePodEvent(conn, jobEntity.JobUuid, jobEntity.WalletAddress)
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

func handlePodEvent(conn *websocket.Conn, jobUuid string, walletAddress string) {
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
		if strings.Contains(event.InvolvedObject.Name, jobUuid) {
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
			LabelSelector: fmt.Sprintf("lad_app=%s", jobDetail.JobUuid),
		})
		if err != nil {
			logs.GetLogger().Errorf("Error listing Pods: %v", err)
			return
		}

		if len(pods.Items) > 0 {
			line := int64(1000)
			containerStatuses := pods.Items[0].Status.ContainerStatuses
			if len(containerStatuses) == 0 {
				return
			}
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

func DeploySpaceTask(jobData models.JobData, deployParam DeployParam, hostName string, gpuProductName string, nodePort int32, jobType int, ipWhiteList []string, gpuIndex []string) {
	saveGpuCache(gpuProductName)
	updateJobStatus(jobData.UUID, models.DEPLOY_UPLOAD_RESULT)
	var success bool
	var jobUuid string
	var walletAddress string
	defer func() {
		deleteGpuCache(gpuProductName)
		if !success {
			k8sNameSpace := constants.K8S_NAMESPACE_NAME_PREFIX + strings.ToLower(walletAddress)
			DeleteJob(k8sNameSpace, jobUuid, "failed to deploy space")
			NewJobService().DeleteJobEntityByJobUuId(jobData.UUID, models.JOB_TERMINATED_STATUS)
		}

		if err := recover(); err != nil {
			logs.GetLogger().Errorf("deploy space task painc, error: %+v", err)
			return
		}
	}()

	spaceDetail, err := getSpaceDetail(jobData.JobSourceURI)
	if err != nil {
		logs.GetLogger().Errorln(err)
		return
	}

	walletAddress = spaceDetail.Data.Owner.PublicAddress
	spaceName := spaceDetail.Data.Space.Name
	jobUuid = strings.ToLower(jobData.UUID)
	spaceHardware := spaceDetail.Data.Space.ActiveOrder.Config

	var deploy *Deploy
	if jobType == 0 {
		logs.GetLogger().Infof("job_uuid: %s, spaceName: %s, hardwareName: %s", jobData.UUID, spaceName, spaceHardware.Description)
		if len(spaceHardware.Description) == 0 {
			return
		}
		deploy = NewDeploy(jobData.UUID, jobUuid, hostName, walletAddress, spaceHardware.Description, int64(jobData.Duration), constants.SPACE_TYPE_PUBLIC, spaceHardware, jobType)
	} else if jobType == 1 {
		deploy = NewDeploy(jobData.UUID, jobUuid, hostName, walletAddress, "", int64(jobData.Duration), constants.SPACE_TYPE_PUBLIC, spaceHardware, jobType)
	}

	deploy.WithIpWhiteList(ipWhiteList)
	deploy.WithSpaceName(spaceName)
	deploy.WithGpuProductName(gpuProductName)
	deploy.WithGpuIndex(gpuIndex)
	deploy.WithSpacePath(deployParam.BuildImagePath)
	if len(deployParam.ModelsSettingFilePath) > 0 {
		err := deploy.WithModelSettingFile(deployParam.ModelsSettingFilePath).ModelInferenceToK8s()
		if err != nil {
			logs.GetLogger().Error(err)
			return
		}
		success = true
		return
	}

	if deployParam.ContainsYaml {
		err := deploy.WithYamlInfo(deployParam.YamlFilePath).YamlToK8s(nodePort)
		if err != nil {
			logs.GetLogger().Errorf("failed to use yaml to deploy job, error: %v", err)
			return
		}
		if deploy.nodePortUrl != "" {
			jobData.JobRealUri = deploy.nodePortUrl[:len(deploy.nodePortUrl)-2]
			if err = submitJob(&jobData); err != nil {
				logs.GetLogger().Errorf("failed to upload job result to MCS, jobUuid: %s, error: %v", jobData.UUID, err)
				return
			}
			updateJobStatus(jobData.UUID, models.DEPLOY_TO_K8S)
		}
		success = true
	} else {
		imageName, dockerfilePath := BuildImagesByDockerfile(jobData.UUID, spaceName, deployParam.BuildImagePath)

		clusterRuntime, err := NewK8sService().GetClusterRuntime()
		if err != nil {
			logs.GetLogger().Errorf("failed to get cluster runtime, error: %v", err)
		} else {
			logs.GetLogger().Infof("cluster runtime: %s", clusterRuntime)
			if strings.Contains(strings.ToLower(clusterRuntime), "containerd") {
				imageTar, err := NewDockerService().SaveDockerImage(imageName)
				if err != nil {
					logs.GetLogger().Errorf("failed to save image, imageName: %s error: %v", imageName, err)
					return
				}
				if err = ImportImageToContainerd(imageTar); err != nil {
					logs.GetLogger().Errorf("failed to load image into containerd, imageName: %s error: %v", imageName, err)
					return
				}
			}
			deploy.WithDockerfile(imageName, dockerfilePath).DockerfileToK8s()
			success = true
		}
	}
	return
}

func DeleteJob(namespace, jobUuid string, msg string) error {
	jobUuid = strings.ToLower(jobUuid)
	deployName := constants.K8S_DEPLOY_NAME_PREFIX + jobUuid
	serviceName := constants.K8S_SERVICE_NAME_PREFIX + jobUuid
	ingressName := constants.K8S_INGRESS_NAME_PREFIX + jobUuid

	k8sService := NewK8sService()

	if namespace != "" {
		if err := k8sService.DeleteIngress(context.TODO(), namespace, ingressName); err != nil && !errors.IsNotFound(err) {
			logs.GetLogger().Errorf("Failed delete ingress, ingressName: %s, error: %+v", ingressName, err)
			return err
		}
		logs.GetLogger().Infof(" deleted ingress, job_uuid: %s, ingressName: %s", jobUuid, ingressName)

		if err := k8sService.DeleteService(context.TODO(), namespace, serviceName); err != nil && !errors.IsNotFound(err) {
			logs.GetLogger().Errorf("Failed delete service, serviceName: %s, error: %+v", serviceName, err)
			return err
		}
		logs.GetLogger().Infof(" deleted service, job_uuid: %s, serviceName: %s", jobUuid, serviceName)

		dockerService := NewDockerService()
		deployImageIds, err := k8sService.GetDeploymentImages(context.TODO(), namespace, deployName)
		if err != nil && !errors.IsNotFound(err) {
			logs.GetLogger().Errorf("Failed get deploy imageIds, deployName: %s, error: %+v", deployName, err)
			return err
		}
		for _, imageId := range deployImageIds {
			dockerService.RemoveImage(imageId)
			logs.GetLogger().Infof(" deleted images, job_uuid: %s, image: %s", jobUuid, imageId)
		}

		if err := k8sService.DeleteDeployment(context.TODO(), namespace, deployName); err != nil && !errors.IsNotFound(err) {
			logs.GetLogger().Errorf("Failed delete deployment, deployName: %s, error: %+v", deployName, err)
			return err
		}
		logs.GetLogger().Infof(" deleted deployment, job_uuid: %s, deployName: %s", jobUuid, deployName)
		time.Sleep(6 * time.Second)

		if err := k8sService.DeleteDeployRs(context.TODO(), namespace, jobUuid); err != nil && !errors.IsNotFound(err) {
			logs.GetLogger().Errorf("Failed delete ReplicaSetsController, job_uuid: %s, error: %+v", jobUuid, err)
			return err
		}
		logs.GetLogger().Infof(" deleted ReplicaSetsController, job_uuid: %s", jobUuid)

		if err := k8sService.DeletePod(context.TODO(), namespace, jobUuid); err != nil && !errors.IsNotFound(err) {
			logs.GetLogger().Errorf("Failed delete pods, job_uuid: %s, error: %+v", jobUuid, err)
			return err
		}
		logs.GetLogger().Infof(" deleted pod, jobUuid: %s", jobUuid)
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
		getPods, err := k8sService.GetPods(namespace, jobUuid)
		if err != nil && !errors.IsNotFound(err) {
			logs.GetLogger().Errorf("Failed get pods form namespace, namepace: %s, error: %+v", namespace, err)
			continue
		}
		if !getPods {
			break
		}
	}

	if msg != "" {
		logs.GetLogger().Infof("%s, job_uuid: %s", msg, jobUuid)
	} else {
		logs.GetLogger().Infof("delete job service finished, job_uuid: %s", jobUuid)
	}

	return nil
}

func downloadModelUrl(namespace, jobUuid, serviceIp string, podCmd []string) {
	k8sService := NewK8sService()
	podName, err := k8sService.WaitForPodRunningByHttp(namespace, jobUuid, serviceIp)
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

	readAll, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.SpaceJSON{}, fmt.Errorf("failed to read job_source_url body, error: %v", err)
	}

	var spaceJson models.SpaceJSON
	if err := json.Unmarshal(readAll, &spaceJson); err != nil {
		return models.SpaceJSON{}, fmt.Errorf("error decoding Space API response JSON: %v", err)
	}

	if spaceJson.Data.Files == nil {
		var spaceJsonWithNoData models.SpaceJsonWithNoData
		if err := json.Unmarshal(readAll, &spaceJsonWithNoData); err != nil {
			return models.SpaceJSON{}, fmt.Errorf("error decoding Space API response JSON: %v", err)
		}
		spaceJson.Data = spaceJsonWithNoData
	}

	return spaceJson, nil
}

func checkResourceAvailableForSpace(jobType int, resourceConfig models.SpaceHardware) (bool, string, []string, error) {
	var taskType string
	var hardwareDetail models.Resource
	if jobType == 1 {
		taskType, hardwareDetail = getHardwareDetailByByte(resourceConfig)
	} else {
		taskType, hardwareDetail = getHardwareDetail(resourceConfig.Description)
	}

	k8sService := NewK8sService()

	activePods, err := k8sService.GetAllActivePod(context.TODO())
	if err != nil {
		return false, "", nil, err
	}

	nodes, err := k8sService.k8sClient.CoreV1().Nodes().List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return false, "", nil, err
	}

	nodeGpuSummary, err := k8sService.GetNodeGpuSummary(context.TODO())
	if err != nil {
		logs.GetLogger().Errorf("Failed collect k8s gpu, error: %+v", err)
		return false, "", nil, err
	}

	gpuName := strings.ToUpper(strings.ReplaceAll(hardwareDetail.Gpu.Unit, " ", "-"))
	gpuNum := hardwareDetail.Gpu.Quantity
	needCpu := hardwareDetail.Cpu.Quantity
	needMemory := float64(hardwareDetail.Memory.Quantity)
	needStorage := float64(hardwareDetail.Storage.Quantity)
	logs.GetLogger().Infof("checkResourceForSpace: needCpu: %d, needMemory: %.2f, needStorage: %.2f, needGpu: %s, gpuNum: %d", needCpu, needMemory, needStorage, gpuName, gpuNum)

	type gpuData struct {
		Total     int
		Free      int
		FreeIndex []string
	}

	for _, node := range nodes.Items {
		var nodeName = node.Name
		var nodeGpuInfo = nodeGpuSummary[nodeName]
		nodeGpu, remainderResource, _ := GetNodeResource(activePods, &node)
		remainderCpu := remainderResource[ResourceCpu]
		remainderMemory := float64(remainderResource[ResourceMem] / 1024 / 1024 / 1024)
		remainderStorage := float64(remainderResource[ResourceStorage] / 1024 / 1024 / 1024)

		freeGpuMap := make(map[string]gpuData)
		for gn, gd := range nodeGpuInfo {
			freeGpuMap[gn] = gpuData{
				Total:     gd.Total,
				Free:      gd.Free,
				FreeIndex: gd.FreeIndex,
			}
		}

		logs.GetLogger().Infof("checkResourceAvailableForSpace: nodeName: %s,remainingCpu: %d, remainingMemory: %.2f, remainingStorage: %.2f, remainingGpu: %+v", node.Name, remainderCpu, remainderMemory, remainderStorage, freeGpuMap)
		if needCpu <= remainderCpu && needMemory <= remainderMemory && needStorage <= remainderStorage {
			if taskType == "CPU" {
				return true, "", nil, nil
			} else if taskType == "GPU" {
				for gname, gData := range nodeGpuInfo {
					logs.GetLogger().Infof("nodeGpuInfo: %+v, used gpu: %+v; gname: %s, gpuName: %s", nodeGpuInfo, nodeGpu, gname, gpuName)
					if strings.Contains(gname, gpuName) {
						gpuName = gname
						remainingGpu := difference(gData.FreeIndex, nodeGpu[gpuName].UsedIndex)
						if gpuNum <= int64(len(remainingGpu)) {
							return true, gpuName, remainingGpu, nil
						}
					}
				}
				continue
			}
		}
	}
	return false, "", nil, nil
}

func checkResourceAvailableForUbi(taskType int, gpuName string, resource *models.TaskResource) (string, string, int64, int64, int64, []string, error) {
	k8sService := NewK8sService()
	activePods, err := k8sService.GetAllActivePod(context.TODO())
	if err != nil {
		return "", "", 0, 0, 0, nil, err
	}

	nodes, err := k8sService.k8sClient.CoreV1().Nodes().List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return "", "", 0, 0, 0, nil, err
	}

	nodeGpuSummary, err := k8sService.GetNodeGpuSummary(context.TODO())
	if err != nil {
		logs.GetLogger().Errorf("Failed collect k8s gpu, error: %+v", err)
		return "", "", 0, 0, 0, nil, err
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
				return nodeName, architecture, needCpu, int64(needMemory), int64(needStorage), nil, nil
			} else if taskType == 1 {
				if gpuName == "" {
					nodeName = ""
					continue
				}
				gpuName = strings.ReplaceAll(gpuName, " ", "-")
				logs.GetLogger().Infof("needGpuName: %s, nodeGpu: %+v, nodeGpuSummary: %+v", gpuName, nodeGpu, nodeGpuSummary)

				if gData, ok := nodeGpuSummary[node.Name][gpuName]; ok {
					remainingGpu := difference(gData.FreeIndex, nodeGpu[gpuName].UsedIndex)
					if len(remainingGpu) > 0 {
						return nodeName, architecture, needCpu, int64(needMemory), int64(needStorage), remainingGpu, nil
					}
				}
				nodeName = ""
				continue
			}
		}
	}
	return nodeName, architecture, needCpu, int64(needMemory), int64(needStorage), nil, nil
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

func generateStringNoNum(length int) string {
	characters := "abcdefghijklmnopqrstuvwxyz"
	source := characters
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
		return true
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
		return true
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

func getJobExpiredTime(jobEntity models.JobEntity) int64 {
	var expiredTime = jobEntity.ExpireTime
	var taskInfoOnChain models.TaskInfoOnChain
	var err error

	for i := 0; i < 5; i++ {
		taskInfoOnChain, err = getJobOnChain(jobEntity.TaskUuid)
		if err != nil {
			time.Sleep(5 * time.Second)
			continue
		}
		break
	}
	expiredTime = taskInfoOnChain.StartTimestamp + taskInfoOnChain.Duration
	return expiredTime
}

func getJobOnChain(taskUuid string) (models.TaskInfoOnChain, error) {
	chainRpc, err := conf.GetRpcByNetWorkName()
	if err != nil {
		return models.TaskInfoOnChain{}, err
	}
	client, err := contract.GetEthClient(chainRpc)
	if err != nil {
		return models.TaskInfoOnChain{}, fmt.Errorf("failed to rpc, error: %v", err)
	}
	defer client.Close()
	taskManagerStub, err := fcp.NewTaskManagerStub(client)
	if err != nil {
		return models.TaskInfoOnChain{}, err
	}
	return taskManagerStub.GetTaskInfo(taskUuid)
}

type DeployParam struct {
	ContainsYaml          bool
	YamlFilePath          string
	BuildImagePath        string
	ModelsSettingFilePath string
}

func saveGpuCache(gpuProductName string) {
	value, ok := gpuResourceCache.Load(gpuProductName)
	if ok {
		value = value.(int) + 1
	} else {
		value = 1
	}
	gpuResourceCache.Store(gpuProductName, value)

}

func deleteGpuCache(gpuProductName string) {
	value, ok := gpuResourceCache.Load(gpuProductName)
	if ok {
		value = value.(int) - 1
	}
	gpuResourceCache.Store(gpuProductName, value)

}

func checkPrice(userPrice string, duration int, resource models.SpaceHardware) (bool, float64, error) {
	priceConfig, err := ReadPriceConfig()
	if err != nil {
		return false, 0, err
	}

	userPayPrice, err := parsePrice(userPrice)
	if err != nil {
		return false, 0, fmt.Errorf("failed to converting user price: %v", err)
	}

	// Convert price strings to float64
	cpuPrice, err := parsePrice(priceConfig.TARGET_CPU)
	if err != nil {
		return false, 0, fmt.Errorf("failed to converting CPU price: %v", err)
	}
	memoryPrice, err := parsePrice(priceConfig.TARGET_MEMORY)
	if err != nil {
		return false, 0, fmt.Errorf("failed to converting Memory price: %v", err)
	}
	storagePrice, err := parsePrice(priceConfig.TARGET_HD_EPHEMERAL)
	if err != nil {
		return false, 0, fmt.Errorf("failed to converting Storage price: %v", err)
	}

	var gpuPriceStr string
	var gpuPrice float64
	if len(priceConfig.GpusPrice) != 0 {

		gpuName := strings.ReplaceAll(resource.Hardware, "NVIDIA ", "")
		gpuName = strings.ReplaceAll(gpuName, " ", "_")
		key := "TARGET_GPU_" + gpuName
		if price, ok := priceConfig.GpusPrice[key]; ok {
			gpuPriceStr = price
		}
	}
	if gpuPriceStr == "" {
		gpuPriceStr = priceConfig.TARGET_GPU_DEFAULT
	}
	gpuPrice, err = parsePrice(priceConfig.TARGET_GPU_DEFAULT)
	if err != nil {
		return false, 0, fmt.Errorf("failed to converting GPU price: %v", err)
	}

	// Calculate total cost
	cpuCost := float64(resource.Vcpu) * cpuPrice * float64(duration/3600)
	memoryCost := float64(resource.Memory/1024/1024/1024) * memoryPrice * float64(duration/3600)
	storageCost := float64(resource.Storage/1024/1024/1024) * storagePrice * float64(duration/3600)
	gpuCost := float64(1) * gpuPrice * float64(duration/3600)

	totalCost := cpuCost + memoryCost + storageCost + gpuCost

	// Compare user's price with total cost
	return userPayPrice >= totalCost, totalCost, nil
}

func difference(arr1, arr2 []string) []string {
	set := make(map[string]struct{})
	for _, v := range arr2 {
		set[v] = struct{}{}
	}
	var result []string
	for _, v := range arr1 {
		if _, found := set[v]; !found {
			result = append(result, v)
		}
	}
	return result
}

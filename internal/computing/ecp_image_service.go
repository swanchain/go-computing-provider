package computing

import "C"
import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
	"github.com/filswan/go-swan-lib/logs"
	"github.com/gin-gonic/gin"
	"github.com/swanchain/go-computing-provider/conf"
	"github.com/swanchain/go-computing-provider/internal/contract"
	"github.com/swanchain/go-computing-provider/internal/models"
	"github.com/swanchain/go-computing-provider/util"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const traefikListenPortMapHost = 9000

type ImageJobService struct {
}

func NewImageJobService() *ImageJobService {
	return &ImageJobService{}
}

func (*ImageJobService) CheckJobCondition(c *gin.Context) {
	var job models.EcpImageJobReq
	err := c.ShouldBindJSON(&job)
	if err != nil {
		logs.GetLogger().Errorf("failed to parse json, error: %v", err)
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.JsonError))
		return
	}
	logs.GetLogger().Infof("check job condition, received Data: %+v", job.Resource)

	var totalCost float64
	var checkPriceFlag bool
	if !conf.GetConfig().API.Pricing {
		checkPriceFlag, totalCost, err = checkPriceForDocker(job.Price, job.Duration, job.Resource)
		if err != nil {
			logs.GetLogger().Errorf("failed to check price, error: %v", err)
			c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.CheckPriceError))
			return
		}

		if !checkPriceFlag {
			logs.GetLogger().Errorf("bid below the set price, pid: %s, need: %0.4f", job.Price, totalCost)
			c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.BelowPriceError))
			return
		}
	}

	receive, _, _, _, _, err := checkResourceForImage(job.Resource)
	if receive {
		c.JSON(http.StatusOK, util.CreateSuccessResponse(map[string]interface{}{
			"price": totalCost,
		}))
	} else {
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.NoAvailableResourcesError))
	}
	return
}

func (imageJob *ImageJobService) DeployJob(c *gin.Context) {
	var job models.EcpImageJobReq
	err := c.ShouldBindJSON(&job)
	if err != nil {
		logs.GetLogger().Errorf("failed to parse json, error: %v", err)
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.JsonError))
		return
	}
	logs.GetLogger().Infof("Job received Data: %+v", job)

	if !CheckWalletWhiteListForEcp(job.WalletAddress) {
		logs.GetLogger().Errorf("This cp does not accept tasks from wallet addresses outside the whitelist")
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.SpaceCheckWhiteListError))
		return
	}

	if CheckWalletBlackListForEcp(job.WalletAddress) {
		logs.GetLogger().Errorf("This cp does not accept tasks from wallet addresses inside the blacklist")
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.SpaceCheckBlackListError))
		return
	}

	if job.JobType == 0 {
		job.JobType = 1
	}

	if strings.TrimSpace(job.Uuid) == "" {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.UbiTaskParamError, "missing required field: [uuid]"))
		return
	}
	if strings.TrimSpace(job.Name) == "" {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.UbiTaskParamError, "missing required field: [name]"))
		return
	}
	if job.Resource == nil {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.UbiTaskParamError, "missing required field: [resource]"))
		return
	}
	if job.JobType != models.MiningJobType && job.JobType != models.InferenceJobType {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.UbiTaskParamError, "invalidate value: [job_type], support: 1 or 2"))
		return
	}
	if job.Image == "" {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.UbiTaskParamError, "missing required field: [image]"))
	}

	if conf.GetConfig().UBI.VerifySign {
		if len(job.Sign) == 0 {
			c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.BadParamError, "missing required field: [sign]"))
			return
		}
		if len(job.WalletAddress) == 0 {
			c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.BadParamError, "missing required field: [wallet_address]"))
			return
		}

		cpAccountAddress, err := contract.GetCpAccountAddress()
		if err != nil {
			logs.GetLogger().Errorf("get cp account contract address failed, error: %v", err)
			c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.GetCpAccountError))
			return
		}

		signature, err := verifySignature(job.WalletAddress, fmt.Sprintf("%s%s", cpAccountAddress, job.Uuid), job.Sign)
		if err != nil {
			logs.GetLogger().Errorf("failed to verifySignature for ecp job, error: %+v", err)
			c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.SignatureError, "verify sign data occur error"))
			return
		}

		logs.GetLogger().Infof("ubi task sign verifing, task_id: %s, verify: %v", job.Uuid, signature)
		if !signature {
			c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.SignatureError, "signature verify failed"))
			return
		}
	}

	var totalCost float64
	var checkPriceFlag bool
	if !conf.GetConfig().API.Pricing && job.Price != "-1" {
		checkPriceFlag, totalCost, err = checkPriceForDocker(job.Price, job.Duration, job.Resource)
		if err != nil {
			logs.GetLogger().Errorf("failed to check price, job_uuid: %s, error: %v", job.Uuid, err)
			c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.CheckPriceError))
			return
		}

		if !checkPriceFlag {
			logs.GetLogger().Errorf("bid below the set price, job_uuid: %s, pid: %s, need: %0.4f", job.Uuid, job.Price, totalCost)
			c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.BelowPriceError))
			return
		}
	}

	isReceive, _, needCpu, _, indexs, err := checkResourceForImage(job.Resource)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.CheckResourcesError))
		return
	}
	if !isReceive {
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.NoAvailableResourcesError))
		return
	}

	var envs []string
	var needResource container.Resources
	var useIndexs []string
	if job.Resource.GPUModel != "" && job.Resource.GPU > 0 {
		for i := 0; i < job.Resource.GPU; i++ {
			if i >= len(indexs) {
				break
			}
			useIndexs = append(useIndexs, indexs[i])
		}
		envs = append(envs, fmt.Sprintf("CUDA_VISIBLE_DEVICES=%s", strings.Join(useIndexs, ",")))
		needResource = container.Resources{
			CPUQuota: needCpu * 100000,
			Memory:   job.Resource.Memory,
			DeviceRequests: []container.DeviceRequest{
				{
					Driver:       "nvidia",
					DeviceIDs:    useIndexs,
					Capabilities: [][]string{{"compute", "utility"}},
				},
			},
		}
	} else {
		needResource = container.Resources{
			CPUQuota: needCpu * 100000,
			Memory:   job.Resource.Memory,
		}
	}

	var handleCmd []string
	for _, cmd := range job.Cmd {
		handleCmd = append(handleCmd, strings.TrimSpace(cmd))
	}
	job.Cmd = handleCmd

	var deployJob models.DeployJobParam
	deployJob.Uuid = job.Uuid
	deployJob.Name = job.Name
	deployJob.JobType = job.JobType
	deployJob.HealthPath = job.HealthPath
	deployJob.NeedResource = needResource

	if len(job.RunCommands) == 0 {
		deployJob.Image = job.Image
		deployJob.Cmd = job.Cmd
		var ports []int
		for _, p := range job.Ports {
			ports = append(ports, p...)
		}
		deployJob.Ports = ports

		for k, v := range job.Envs {
			envs = append(envs, fmt.Sprintf("%s=%s", k, v))
		}
		deployJob.Envs = envs
	} else {
		buildParams, err := parseDockerfileConfigForEcp(job)
		if err != nil {
			logs.GetLogger().Errorln(err)
			return
		}
		deployJob.Image = buildParams.Image
		deployJob.Cmd = buildParams.Cmd
		deployJob.Ports = buildParams.Ports
		deployJob.Envs = buildParams.Envs
	}

	if job.Price == "-1" {
		if job.JobType == models.MiningJobType {
			var maxID int64
			result := NewTaskService().Model(&models.TaskEntity{}).Select("MAX(id)").Scan(&maxID)
			if result.Error != nil {
				logs.GetLogger().Errorf("failed to fetch max task id, error: %v", err)
				c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.SaveTaskEntityError))
				return
			}

			var taskEntity = new(models.TaskEntity)
			taskEntity.Id = maxID + 1
			taskEntity.Uuid = job.Uuid
			taskEntity.Type = models.Mining
			taskEntity.Name = job.Name
			taskEntity.ResourceType = models.RESOURCE_TYPE_GPU
			taskEntity.Status = models.TASK_RECEIVED_STATUS
			taskEntity.CreateTime = time.Now().Unix()
			taskEntity.Sequencer = 0
			err = NewTaskService().SaveTaskEntity(taskEntity)
			if err != nil {
				logs.GetLogger().Errorf("save task entity failed, error: %v", err)
				c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.SaveTaskEntityError))
				return
			}
		}
	} else {
		if err = NewEcpJobService().SaveEcpJobEntity(&models.EcpJobEntity{
			Uuid:          job.Uuid,
			Name:          job.Name,
			Image:         deployJob.Image,
			Env:           strings.Join(deployJob.Envs, ","),
			JobType:       job.JobType,
			Cmd:           strings.Join(deployJob.Cmd, ","),
			Cpu:           job.Resource.CPU,
			Memory:        job.Resource.Memory,
			Storage:       job.Resource.Storage,
			GpuName:       job.Resource.GPUModel,
			GpuIndex:      strings.Join(useIndexs, ","),
			Status:        models.CreatedStatus,
			HealthUrlPath: job.HealthPath,
			CreateTime:    time.Now().Unix(),
		}); err != nil {
			logs.GetLogger().Errorf("failed to save job to db, error: %v", err)
			c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.SaveTaskEntityError))
			return
		}
	}

	var logUrl string
	multiAddressSplit := strings.Split(conf.GetConfig().API.MultiAddress, "/")
	if len(multiAddressSplit) >= 4 {
		logUrl = fmt.Sprintf("http://%s:%s/api/v1/computing/cp/job/log?job_uuid=%s&expire_time=60", multiAddressSplit[2], multiAddressSplit[4], job.Uuid)
	}

	if job.JobType == models.MiningJobType {
		imageJob.DeployMining(c, deployJob, totalCost, logUrl)
		return
	} else if job.JobType == models.InferenceJobType {
		imageJob.DeployInference(c, deployJob, totalCost, logUrl)
		return
	}
}

func (*ImageJobService) GetJobStatus(c *gin.Context) {
	jobUuid := c.Query("job_uuid")
	ecpJobs, err := NewEcpJobService().GetEcpJobs(jobUuid)
	if err != nil {
		return
	}

	containerStatus, err := NewDockerService().GetContainerStatus()
	if err != nil {
		return
	}

	var result []models.EcpJobStatusResp
	for _, entity := range ecpJobs {
		var statusStr = entity.Status
		if status, ok := containerStatus[entity.ContainerName]; ok {
			fmt.Printf("container name: %s, status: %s \n", entity.ContainerName, status)
			statusStr = status
		}
		result = append(result, models.EcpJobStatusResp{Uuid: entity.Uuid, Status: statusStr, Message: entity.Message})
	}

	c.JSON(http.StatusOK, util.CreateSuccessResponse(result))
}

func (*ImageJobService) DeleteJob(c *gin.Context) {
	jobUuId := c.Param("job_uuid")
	if strings.TrimSpace(jobUuId) == "" {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.BadParamError, "missing required field: [job_uuid]"))
		return
	}

	ecpJobEntity, err := NewEcpJobService().GetEcpJobByUuid(jobUuId)
	if err != nil {
		logs.GetLogger().Errorf("failed to get job, job_uuid: %s, error: %v", jobUuId, err)
		return
	}
	containerName := ecpJobEntity.ContainerName
	if err = NewDockerService().RemoveContainerByName(containerName); err != nil {
		logs.GetLogger().Errorf("failed to remove container, job_uuid: %s, error: %v", jobUuId, err)
		return
	}
	NewEcpJobService().DeleteContainerByUuid(jobUuId)

	c.JSON(http.StatusOK, util.CreateSuccessResponse("success"))
}

func (*ImageJobService) DockerLogsHandler(c *gin.Context) {
	jobUuId := c.Query("job_uuid")
	expireTimeStr := c.Query("expire_time")
	if strings.TrimSpace(jobUuId) == "" {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.BadParamError, "missing required field: [job_uuid]"))
		return
	}

	var expireTime int64
	var err error
	if expireTimeStr != "" {
		expireTime, err = strconv.ParseInt(expireTimeStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.BadParamError, "invalidate field: [expire_time]"))
			return
		}
	} else {
		expireTime = 60
	}

	ecpJobEntity, err := NewEcpJobService().GetEcpJobByUuid(jobUuId)
	if err != nil {
		logs.GetLogger().Errorf("failed to get job, job_uuid: %s, error: %v", jobUuId, err)
		return
	}
	containerName := ecpJobEntity.ContainerName
	c.Writer.Header().Set("Content-Type", "text/plain")
	c.Writer.Header().Set("Transfer-Encoding", "chunked")
	c.Writer.Header().Set("Cache-Control", "no-cache")

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(expireTime)*time.Second)
	defer cancel()

	containerLogStream, err := NewDockerService().GetContainerLogStream(ctx, containerName)
	if err != nil {
		logs.GetLogger().Errorf("get docker container log stream failed, error: %v", err)
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.ServerError, "failed to get logs: "+err.Error()))
		return
	}
	defer containerLogStream.Close()
	_, err = io.Copy(c.Writer, containerLogStream)
	if err != nil && errors.Is(ctx.Err(), context.DeadlineExceeded) {
		c.Writer.Write([]byte("Log streaming stopped: timeout reached.\n"))
		return
	}
}

func (*ImageJobService) DeployMining(c *gin.Context, deployJob models.DeployJobParam, totalCost float64, logUrl string) {
	var err error
	containerName := deployJob.Name + "-" + generateString(5)
	go func() {
		if deployJob.BuildImagePath != "" && deployJob.BuildImageName != "" {
			if err := NewDockerService().BuildImage(deployJob.BuildImagePath, deployJob.BuildImageName); err != nil {
				logs.GetLogger().Errorf("failed to building %s image, job_uuid: %s, error: %v", deployJob.Image, deployJob.Uuid, err)
				NewEcpJobService().UpdateEcpJobEntityMessage(deployJob.Uuid, fmt.Sprintf("failed to build image: %s", deployJob.Image))
			}
		} else {
			if err := NewDockerService().PullImage(deployJob.Image); err != nil {
				logs.GetLogger().Errorf("failed to pull %s image, job_uuid: %s, error: %v", deployJob.Image, deployJob.Uuid, err)
				NewEcpJobService().UpdateEcpJobEntityMessage(deployJob.Uuid, fmt.Sprintf("failed to pull image: %s", deployJob.Image))
				return
			}
		}

		hostConfig := &container.HostConfig{
			Resources:  deployJob.NeedResource,
			Privileged: true,
		}
		containerConfig := &container.Config{
			Image:        deployJob.Image,
			Env:          deployJob.Envs,
			AttachStdout: true,
			AttachStderr: true,
			Tty:          true,
		}
		dockerService := NewDockerService()
		if err := dockerService.ContainerCreateAndStart(containerConfig, hostConfig, nil, containerName); err != nil {
			logs.GetLogger().Errorf("failed to create job container, job_uuid: %s, error: %v", deployJob.Uuid, err)
			NewEcpJobService().UpdateEcpJobEntityMessage(deployJob.Uuid, "failed to create container")
			return
		}
		logs.GetLogger().Warnf("job_uuid: %s, starting container, container name: %s", deployJob.Uuid, containerName)

		time.Sleep(3 * time.Second)
		if !dockerService.IsExistContainer(containerName) {
			logs.GetLogger().Warnf("job_uuid: %s, not found container", deployJob.Uuid)
			NewEcpJobService().UpdateEcpJobEntityMessage(deployJob.Uuid, "failed to start container")
			return
		}
		logs.GetLogger().Warnf("job_uuid: %s, started container, container name: %s", deployJob.Uuid, containerName)

		if err = NewEcpJobService().UpdateEcpJobEntityContainerName(deployJob.Uuid, containerName); err != nil {
			logs.GetLogger().Errorf("failed to save job to db, error: %v", err)
			return
		}
	}()

	c.JSON(http.StatusOK, util.CreateSuccessResponse(map[string]interface{}{
		"uuid":    deployJob.Uuid,
		"price":   totalCost,
		"log_url": logUrl,
	}))
}

func (*ImageJobService) DeployInference(c *gin.Context, deployJob models.DeployJobParam, totalCost float64, logUrl string) {
	var containerName = deployJob.Name + "-" + generateString(5)
	var err error
	var apiUrl string
	var portBinding map[nat.Port][]nat.PortBinding
	var portMaps []models.PortMap
	var labelMap map[string]string
	if len(deployJob.Ports) > 1 {
		portBinding, portMaps, err = handleMultiPort(deployJob.Ports)
		multiAddressSplit := strings.Split(conf.GetConfig().API.MultiAddress, "/")
		apiUrl = multiAddressSplit[2]
	} else {
		prefixStr := generateString(10)
		if strings.HasPrefix(conf.GetConfig().API.Domain, ".") {
			apiUrl = prefixStr + conf.GetConfig().API.Domain
		} else {
			apiUrl = strings.Join([]string{prefixStr, conf.GetConfig().API.Domain}, ".")
		}

		labelMap = map[string]string{
			"traefik.enable": "true",
			fmt.Sprintf("traefik.http.routers.%s.entrypoints", containerName): "web",
			fmt.Sprintf("traefik.http.routers.%s.rule", containerName):        fmt.Sprintf("Host(`%s`)", apiUrl),
		}
		apiUrl += fmt.Sprintf(":%d", traefikListenPortMapHost)
	}

	go func() {
		if deployJob.BuildImagePath != "" && deployJob.BuildImageName != "" {
			if err := NewDockerService().BuildImage(deployJob.BuildImagePath, deployJob.BuildImageName); err != nil {
				logs.GetLogger().Errorf("failed to building %s image, job_uuid: %s, error: %v", deployJob.Image, deployJob.Uuid, err)
				NewEcpJobService().UpdateEcpJobEntityMessage(deployJob.Uuid, fmt.Sprintf("failed to build image: %s", deployJob.Image))
			}
		} else {
			if err := NewDockerService().PullImage(deployJob.Image); err != nil {
				logs.GetLogger().Errorf("failed to pull %s image, job_uuid: %s, error: %v", deployJob.Image, deployJob.Uuid, err)
				NewEcpJobService().UpdateEcpJobEntityMessage(deployJob.Uuid, fmt.Sprintf("failed to pull image: %s", deployJob.Image))
				return
			}
		}

		hostConfig := &container.HostConfig{
			Resources:  deployJob.NeedResource,
			Privileged: true,
		}
		containerConfig := &container.Config{
			Image:        deployJob.Image,
			Env:          deployJob.Envs,
			Cmd:          deployJob.Cmd,
			AttachStdout: true,
			AttachStderr: true,
			Tty:          true,
		}

		var networkConfig *network.NetworkingConfig
		if len(deployJob.Ports) > 1 {
			hostConfig.PortBindings = portBinding
		} else {
			containerConfig.Labels = labelMap
			networkConfig = &network.NetworkingConfig{
				EndpointsConfig: map[string]*network.EndpointSettings{
					"traefik-net": {},
				},
			}
		}

		dockerService := NewDockerService()
		if err := dockerService.ContainerCreateAndStart(containerConfig, hostConfig, networkConfig, containerName); err != nil {
			logs.GetLogger().Errorf("failed to create job container, job_uuid: %s, error: %v", deployJob.Uuid, err)
			NewEcpJobService().UpdateEcpJobEntityMessage(deployJob.Uuid, "failed to create container")
			return
		}
		logs.GetLogger().Warnf("job_uuid: %s, starting container, container name: %s", deployJob.Uuid, containerName)

		time.Sleep(3 * time.Second)
		if !dockerService.IsExistContainer(containerName) {
			logs.GetLogger().Warnf("job_uuid: %s, not found container", deployJob.Uuid)
			NewEcpJobService().UpdateEcpJobEntityMessage(deployJob.Uuid, "failed to start container")
			return
		}
		logs.GetLogger().Warnf("job_uuid: %s, started container, container name: %s", deployJob.Uuid, containerName)

		if err = NewEcpJobService().UpdateEcpJobEntityContainerName(deployJob.Uuid, containerName); err != nil {
			logs.GetLogger().Errorf("failed to save job to db, error: %v", err)
			return
		}
	}()

	var inferenceResp models.EcpImageResp
	inferenceResp.UUID = deployJob.Uuid
	inferenceResp.ServiceUrl = "http://" + apiUrl
	inferenceResp.HealthPath = deployJob.HealthPath
	inferenceResp.ServicePortMapping = portMaps
	inferenceResp.Price = totalCost
	inferenceResp.LogUrl = logUrl

	var portMap []string
	if len(deployJob.Ports) > 1 {
		for _, pm := range portMaps {
			portMap = append(portMap, fmt.Sprintf("%d:%d", pm.ContainerPort, pm.ExternalPort))
		}
	}
	err = NewEcpJobService().UpdateEcpJobEntityPortsAndServiceUrl(deployJob.Uuid, strings.Join(portMap, ","), inferenceResp.ServiceUrl)
	if err != nil {
		logs.GetLogger().Errorf("failed to update job service url to db, error: %v", err)
		return
	}

	c.JSON(http.StatusOK, util.CreateSuccessResponse(inferenceResp))
}

func handleMultiPort(ports []int) (map[nat.Port][]nat.PortBinding, []models.PortMap, error) {
	var portRange = conf.GetConfig().API.PortRange
	var usedPort = make(map[int]int)
	var count int
	for i := 0; i < len(ports); i++ {
		for j := 0; i < len(portRange); j++ {
			if _, ok := usedPort[portRange[j]]; !ok {
				if util.IsPortAvailable(portRange[j]) {
					usedPort[portRange[j]] = ports[i]
					count++
					break
				}
			}
		}
	}
	if len(ports) > count {
		return nil, nil, fmt.Errorf("port number unavailable, need %d, available: %d", len(ports), count)
	}

	var mapPorts []models.PortMap
	var portMap = make(map[nat.Port][]nat.PortBinding)
	for hostP, containerP := range usedPort {
		portMap[nat.Port(fmt.Sprintf("%d/tcp", containerP))] = []nat.PortBinding{
			{
				HostIP:   "0.0.0.0",
				HostPort: strconv.Itoa(hostP),
			},
		}
		mapPorts = append(mapPorts, models.PortMap{
			ContainerPort: containerP,
			ExternalPort:  hostP,
		})
	}

	return portMap, mapPorts, nil
}

func checkPriceForDocker(userPrice string, duration int, resource *models.HardwareResource) (bool, float64, error) {
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
	gpuPrice, err := parsePrice(priceConfig.TARGET_GPU_DEFAULT)
	if err != nil {
		return false, 0, fmt.Errorf("failed to converting GPU price: %v", err)
	}

	// Calculate total cost
	cpuCost := float64(resource.CPU) * cpuPrice
	memoryCost := formatGiB(resource.Memory) * memoryPrice
	storageCost := formatGiB(resource.Storage) * storagePrice
	gpuCost := float64(resource.GPU) * gpuPrice

	totalCost := cpuCost + memoryCost + storageCost + gpuCost

	if userPayPrice == 0 {
		logs.GetLogger().Warnf("user's price is 0, use cp price")
		return true, totalCost, nil
	}

	// Compare user's price with total cost
	return userPayPrice >= totalCost, totalCost, nil
}

func checkResourceForImage(resource *models.HardwareResource) (bool, string, int64, int64, []string, error) {
	list, err := NewEcpJobService().GetEcpJobList([]string{models.CreatedStatus, models.RunningStatus})
	if err != nil {
		return false, "", 0, 0, nil, err
	}

	var taskGpuMap = make(map[string][]string)
	for _, g := range list {
		if len(strings.TrimSpace(g.GpuIndex)) > 0 {
			taskGpuMap[g.GpuName] = append(taskGpuMap[g.GpuName], strings.Split(strings.TrimSpace(g.GpuIndex), ",")...)
		}
	}

	dockerService := NewDockerService()
	containerLogStr, err := dockerService.ContainerLogs("resource-exporter")
	if err != nil {
		return false, "", 0, 0, nil, err
	}

	var nodeResource models.NodeResource
	if err := json.Unmarshal([]byte(containerLogStr), &nodeResource); err != nil {
		return false, "", 0, 0, nil, err
	}

	needCpu := resource.CPU
	var needMemory, needStorage float64
	var indexs []string
	if resource.Memory > 0 {
		needMemory = formatGiB(resource.Memory)

	}
	if resource.Storage > 0 {
		needMemory = formatGiB(resource.Storage)
	}

	remainderCpu, _ := strconv.ParseInt(nodeResource.Cpu.Free, 10, 64)

	var remainderMemory, remainderStorage float64
	if len(strings.Split(strings.TrimSpace(nodeResource.Memory.Free), " ")) > 0 {
		remainderMemory, _ = strconv.ParseFloat(strings.Split(strings.TrimSpace(nodeResource.Memory.Free), " ")[0], 64)
	}
	if len(strings.Split(strings.TrimSpace(nodeResource.Storage.Free), " ")) > 0 {
		remainderStorage, err = strconv.ParseFloat(strings.Split(strings.TrimSpace(nodeResource.Storage.Free), " ")[0], 64)
	}

	type gpuData struct {
		num    int
		indexs []string
	}

	var gpuMap = make(map[string]gpuData)
	if nodeResource.Gpu.AttachedGpus > 0 {
		for _, detail := range nodeResource.Gpu.Details {
			if detail.Status == models.Available {
				if checkGpu(detail.ProductName, detail.Index, taskGpuMap) {
					continue
				}

				data, ok := gpuMap[detail.ProductName]
				if ok {
					data.num += 1
					data.indexs = append(data.indexs, detail.Index)
					gpuMap[detail.ProductName] = data
				} else {
					indexs = make([]string, 0)
					indexs = append(indexs, detail.Index)
					var dataNew = gpuData{
						num:    1,
						indexs: indexs,
					}
					gpuMap[detail.ProductName] = dataNew
				}
			}
		}
	}

	logs.GetLogger().Infof("checkResourceForImage: needCpu: %d, needMemory: %.2f, needStorage: %.2f, needGpu: %d, gpuName: %s", needCpu, needMemory, needStorage, resource.GPU, resource.GPUModel)
	logs.GetLogger().Infof("checkResourceForImage: remainingCpu: %d, remainingMemory: %.2f, remainingStorage: %.2f, remainingGpu: %+v", remainderCpu, remainderMemory, remainderStorage, gpuMap)
	if needCpu <= remainderCpu && needMemory <= remainderMemory && needStorage <= remainderStorage {
		if resource.GPUModel != "" {
			var flag bool
			for k, gd := range gpuMap {
				if strings.ToUpper(k) == resource.GPUModel && gd.num > 0 {
					indexs = gd.indexs
					flag = true
					break
				}
			}
			if flag {
				return true, nodeResource.CpuName, needCpu, int64(needMemory), indexs, nil
			} else {
				return false, nodeResource.CpuName, needCpu, int64(needMemory), indexs, nil
			}
		}
		return true, nodeResource.CpuName, needCpu, int64(needMemory), indexs, nil
	}
	return false, nodeResource.CpuName, needCpu, int64(needMemory), indexs, nil
}

func parsePrice(priceStr string) (float64, error) {
	return strconv.ParseFloat(priceStr, 64)
}

func formatTiB(bytes int64) float64 {
	return float64(bytes) / float64(1<<40)
}

func formatGiB(bytes int64) float64 {
	return float64(bytes) / float64(1<<30)
}

func formatMiB(bytes int64) float64 {
	return float64(bytes) / float64(1<<20)
}

func formatKiB(bytes int64) float64 {
	return float64(bytes) / float64(1<<10)
}

func BytesToHumanReadable(bytes int64) string {
	switch {
	case bytes >= 1<<40:
		return fmt.Sprintf("%.2f Ti", formatTiB(bytes))
	case bytes >= 1<<30:
		return fmt.Sprintf("%.2f Gi", formatGiB(bytes))
	case bytes >= 1<<20:
		return fmt.Sprintf("%.2f Mi", formatMiB(bytes))
	case bytes >= 1<<10:
		return fmt.Sprintf("%.2f Ki", formatKiB(bytes))
	}
	return ""
}

func checkGpu(gpuName, index string, taskUseGpu map[string][]string) bool {
	taskUseIndex, ok := taskUseGpu[gpuName]
	if ok {
		for _, taskU := range taskUseIndex {
			if taskU == index {
				return true
			}
		}
	}
	return false
}

func CheckWalletWhiteListForEcp(walletAddress string) bool {
	walletWhiteListUrl := conf.GetConfig().API.WalletWhiteList
	if walletWhiteListUrl == "" {
		return true
	}
	whiteList, err := getWalletList(walletWhiteListUrl)
	if err != nil {
		logs.GetLogger().Errorf("get whiteList By url failed, url: %s, error: %v", err)
		return true
	}

	for _, address := range whiteList {
		if walletAddress == address {
			return true
		}
	}
	return false
}

func CheckWalletBlackListForEcp(walletAddress string) bool {
	walletBlackListUrl := conf.GetConfig().API.WalletBlackList
	if walletBlackListUrl == "" {
		return false
	}
	blackList, err := getWalletList(walletBlackListUrl)
	if err != nil {
		logs.GetLogger().Errorf("get blacklist By url failed, url: %s, error: %v", err)
		return true
	}

	for _, address := range blackList {
		if walletAddress == address {
			return true
		}
	}
	return false
}

func parseDockerfileConfigForEcp(job models.EcpImageJobReq) (*models.DeployJobParam, error) {
	var deployParam *models.DeployJobParam
	if job.RunCommands != nil && len(job.RunCommands) != 0 {
		content, envs, ports := generateDockerfile(job)
		cpRepoPath, _ := os.LookupEnv("CP_PATH")
		buildFolder := filepath.Join(cpRepoPath, "build/ecp", job.Uuid)
		if err := os.MkdirAll(buildFolder, os.ModePerm); err != nil {
			return nil, fmt.Errorf("failed to create dir, error: %v", err)
		}
		dockerfileFile := filepath.Join(buildFolder, "Dockerfile")
		err := os.WriteFile(dockerfileFile, []byte(content), 0644)
		if err != nil {
			return nil, fmt.Errorf("failed to save Dockerfile: %v", err)
		}

		imageName := fmt.Sprintf("ecp-image/%s:%d", job.Uuid, time.Now().Unix())
		if conf.GetConfig().Registry.ServerAddress != "" {
			imageName = fmt.Sprintf("%s/%s:%d",
				strings.TrimSpace(conf.GetConfig().Registry.ServerAddress), job.Uuid, time.Now().Unix())
		}
		imageName = strings.ToLower(imageName)

		if _, err := os.Stat(dockerfileFile); err != nil {
			return nil, fmt.Errorf("failed not found Dockerfile, path: %s", dockerfileFile)
		}
		deployParam.BuildImagePath = buildFolder
		deployParam.BuildImageName = imageName
		deployParam.Envs = envs
		deployParam.Ports = ports
	}
	return deployParam, nil
}

func generateDockerfile(config models.EcpImageJobReq) (string, []string, []int) {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("FROM %s\n", config.Image))

	var envs []string
	var ports []int
	if config.WorkDir != "" {
		builder.WriteString(fmt.Sprintf("WORKDIR %s\n", config.WorkDir))
	}

	for key, value := range config.Envs {
		builder.WriteString(fmt.Sprintf("ENV %s=%s\n", key, value))
		envs = append(envs, fmt.Sprintf("%s=%s", key, value))
	}

	for _, cmd := range config.RunCommands {
		builder.WriteString(fmt.Sprintf("RUN %s\n", cmd))
	}

	for _, port := range config.Ports {
		builder.WriteString(fmt.Sprintf("EXPOSE %d\n", port))
		ports = append(ports, port...)
	}

	if len(config.Cmd) > 0 {
		cmdStr := strings.Join(config.Cmd, " ")
		builder.WriteString(fmt.Sprintf("CMD [%s]\n", cmdStr))
	}

	return builder.String(), envs, ports
}

func ValidateName(name string) error {
	regex := `^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`
	re := regexp.MustCompile(regex)
	if !re.MatchString(name) {
		return fmt.Errorf("invalid field value: %s, must match regex %s", name, regex)
	}
	return nil
}

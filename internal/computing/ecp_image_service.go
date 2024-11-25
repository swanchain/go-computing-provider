package computing

import (
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
	"github.com/filswan/go-swan-lib/logs"
	"github.com/gin-gonic/gin"
	"github.com/swanchain/go-computing-provider/conf"
	"github.com/swanchain/go-computing-provider/internal/models"
	"github.com/swanchain/go-computing-provider/util"
	"net/http"
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
		c.JSON(http.StatusOK, util.CreateSuccessResponse(util.NoAvailableResourcesError))
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

	//if !CheckWalletWhiteListForEcp("") {
	//	logs.GetLogger().Errorf("This cp does not accept tasks from wallet addresses outside the whitelist")
	//	c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.SpaceCheckWhiteListError))
	//	return
	//}
	//
	//if CheckWalletBlackListForEcp("") {
	//	logs.GetLogger().Errorf("This cp does not accept tasks from wallet addresses inside the blacklist")
	//	c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.SpaceCheckBlackListError))
	//	return
	//}

	if job.JobType == 0 {
		job.JobType = 1
	}

	if strings.TrimSpace(job.UUID) == "" {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.UbiTaskParamError, "missing required field: [uuid]"))
		return
	}

	if strings.TrimSpace(job.Name) == "" {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.UbiTaskParamError, "missing required field: [name]"))
		return
	}

	if strings.TrimSpace(job.Image) == "" {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.UbiTaskParamError, "missing required field: [image]"))
		return
	}

	if job.JobType != models.MiningJobType && job.JobType != models.InferenceJobType {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.UbiTaskParamError, "invalidate value: [job_type], support: 1 or 2"))
		return
	}

	if len(job.Sign) == 0 {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.BadParamError, "missing required field: [sign]"))
		return
	}

	//cpAccountAddress, err := contract.GetCpAccountAddress()
	//if err != nil {
	//	logs.GetLogger().Errorf("get cp account contract address failed, error: %v", err)
	//	c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.GetCpAccountError))
	//	return
	//}
	//
	//signature, err := verifySignature(conf.GetConfig().UBI.UbiEnginePk, fmt.Sprintf("%s%s", cpAccountAddress, job.UUID), job.Sign)
	//if err != nil {
	//	logs.GetLogger().Errorf("failed to verifySignature for ecp job, error: %+v", err)
	//	c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.SignatureError, "verify sign data occur error"))
	//	return
	//}
	//
	//logs.GetLogger().Infof("ubi task sign verifing, task_id: %s, verify: %v", job.UUID, signature)
	//if !signature {
	//	c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.SignatureError, "signature verify failed"))
	//	return
	//}

	var totalCost float64
	var checkPriceFlag bool
	if !conf.GetConfig().API.Pricing {
		checkPriceFlag, totalCost, err = checkPriceForDocker(job.Price, job.Duration, job.Resource)
		if err != nil {
			logs.GetLogger().Errorf("failed to check price, job_uuid: %s, error: %v", job.UUID, err)
			c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.CheckPriceError))
			return
		}

		if !checkPriceFlag {
			logs.GetLogger().Errorf("bid below the set price, job_uuid: %s, pid: %s, need: %0.4f", job.UUID, job.Price, totalCost)
			c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.BelowPriceError))
			return
		}
	}

	var env []string
	for k, v := range job.Envs {
		env = append(env, fmt.Sprintf("%s=%s", k, v))
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

	var needResource container.Resources
	var useIndexs []string
	if job.Resource.GPUModel != "" && job.Resource.GPU > 0 {
		for i := 0; i < job.Resource.GPU; i++ {
			if i >= len(indexs) {
				break
			}
			useIndexs = append(useIndexs, indexs[i])
			env = append(env, fmt.Sprintf("CUDA_VISIBLE_DEVICES=%s", strings.Join(useIndexs, ",")))
		}

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

	if err = NewEcpJobService().SaveEcpJobEntity(&models.EcpJobEntity{
		Uuid:       job.UUID,
		Name:       job.Name,
		Image:      job.Image,
		Env:        strings.Join(env, ","),
		JobType:    job.JobType,
		Cmd:        strings.Join(job.Cmd, ","),
		Cpu:        job.Resource.CPU,
		Memory:     job.Resource.Memory,
		Storage:    job.Resource.Storage,
		GpuName:    job.Resource.GPUModel,
		GpuIndex:   strings.Join(useIndexs, ","),
		Status:     models.CreatedStatus,
		CreateTime: time.Now().Unix(),
	}); err != nil {
		logs.GetLogger().Errorf("failed to save job to db, error: %v", err)
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.SaveTaskEntityError))
		return
	}

	if job.JobType == models.MiningJobType {
		imageJob.DeployMining(c, job, needResource, env, totalCost)
		return
	} else if job.JobType == models.InferenceJobType {
		imageJob.DeployInference(c, job, needResource, env, totalCost)
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
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.BadParamError, "missing required field: job_uuid"))
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

func (*ImageJobService) DeployMining(c *gin.Context, job models.EcpImageJobReq, needResource container.Resources, env []string, totalCost float64) {
	containerName := job.Name + "-" + generateString(5)
	go func() {
		err := NewDockerService().PullImage(job.Image)
		if err != nil {
			logs.GetLogger().Errorf("failed to pull %s image, job_uuid: %s, error: %v", job.Image, job.UUID, err)
			NewEcpJobService().UpdateEcpJobEntityMessage(job.UUID, fmt.Sprintf("failed to pull image: %s", job.Image))
			return
		}

		hostConfig := &container.HostConfig{
			Resources:  needResource,
			Privileged: true,
		}
		containerConfig := &container.Config{
			Image:        job.Image,
			Env:          env,
			AttachStdout: true,
			AttachStderr: true,
			Tty:          true,
		}
		dockerService := NewDockerService()
		if err := dockerService.ContainerCreateAndStart(containerConfig, hostConfig, nil, containerName); err != nil {
			logs.GetLogger().Errorf("failed to create job container, job_uuid: %s, error: %v", job.UUID, err)
			NewEcpJobService().UpdateEcpJobEntityMessage(job.UUID, "failed to create container")
			return
		}
		logs.GetLogger().Warnf("job_uuid: %s, starting container, container name: %s", job.UUID, containerName)

		time.Sleep(3 * time.Second)
		if !dockerService.IsExistContainer(containerName) {
			logs.GetLogger().Warnf("job_uuid: %s, not found container", job.UUID)
			NewEcpJobService().UpdateEcpJobEntityMessage(job.UUID, "failed to start container")
			return
		}
		logs.GetLogger().Warnf("job_uuid: %s, started container, container name: %s", job.UUID, containerName)

		if err = NewEcpJobService().UpdateEcpJobEntityContainerName(job.UUID, containerName); err != nil {
			logs.GetLogger().Errorf("failed to save job to db, error: %v", err)
			return
		}
	}()

	c.JSON(http.StatusOK, util.CreateSuccessResponse(map[string]interface{}{
		"price": totalCost,
	}))
}

func (*ImageJobService) DeployInference(c *gin.Context, job models.EcpImageJobReq, needResource container.Resources, env []string, totalCost float64) {
	var containerName = job.Name + "-" + generateString(5)
	var apiUrl string
	var portBinding map[nat.Port][]nat.PortBinding
	var portMaps []models.PortMap
	var labelMap map[string]string
	var err error

	if len(job.Ports) > 1 {
		portBinding, portMaps, err = handleMultiPort(job.Ports)
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
		if err := NewDockerService().PullImage(job.Image); err != nil {
			logs.GetLogger().Errorf("failed to pull %s image, job_uuid: %s, error: %v", job.Image, job.UUID, err)
			NewEcpJobService().UpdateEcpJobEntityMessage(job.UUID, fmt.Sprintf("failed to pull image: %s", job.Image))
			return
		}

		hostConfig := &container.HostConfig{
			Resources:  needResource,
			Privileged: true,
		}
		containerConfig := &container.Config{
			Image:        job.Image,
			Env:          env,
			Cmd:          job.Cmd,
			AttachStdout: true,
			AttachStderr: true,
			Tty:          true,
		}

		var networkConfig *network.NetworkingConfig
		if len(job.Ports) > 1 {
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
			logs.GetLogger().Errorf("failed to create job container, job_uuid: %s, error: %v", job.UUID, err)
			NewEcpJobService().UpdateEcpJobEntityMessage(job.UUID, "failed to create container")
			return
		}
		logs.GetLogger().Warnf("job_uuid: %s, starting container, container name: %s", job.UUID, containerName)

		time.Sleep(3 * time.Second)
		if !dockerService.IsExistContainer(containerName) {
			logs.GetLogger().Warnf("job_uuid: %s, not found container", job.UUID)
			NewEcpJobService().UpdateEcpJobEntityMessage(job.UUID, "failed to start container")
			return
		}
		logs.GetLogger().Warnf("job_uuid: %s, started container, container name: %s", job.UUID, containerName)

		if err = NewEcpJobService().UpdateEcpJobEntityContainerName(job.UUID, containerName); err != nil {
			logs.GetLogger().Errorf("failed to save job to db, error: %v", err)
			return
		}
	}()

	var inferenceResp models.EcpImageResp
	inferenceResp.UUID = job.UUID
	inferenceResp.ServiceUrl = "http://" + apiUrl
	inferenceResp.HealthPath = job.HealthPath
	inferenceResp.ServicePortMapping = portMaps
	inferenceResp.Price = totalCost
	c.JSON(http.StatusOK, util.CreateSuccessResponse(inferenceResp))
}

func handleMultiPort(ports []int) (map[nat.Port][]nat.PortBinding, []models.PortMap, error) {
	var portRange = conf.GetConfig().API.PortRange
	var usedPort = make(map[int]int)
	var count int
	for i := 0; i < len(ports); i++ {
		for j := 0; i < len(portRange); j++ {
			if _, ok := usedPort[portRange[j]]; !ok {
				if util.CheckPortAvailable(portRange[j]) {
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

func checkPriceForDocker(userPrice string, duration int, resource models.HardwareResource) (bool, float64, error) {
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

func checkResourceForImage(resource models.HardwareResource) (bool, string, int64, int64, []string, error) {
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

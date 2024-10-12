package computing

import (
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types/container"
	"github.com/filswan/go-swan-lib/logs"
	"github.com/gin-gonic/gin"
	"github.com/swanchain/go-computing-provider/internal/models"
	"github.com/swanchain/go-computing-provider/util"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type ImageJobService struct {
}

func NewImageJobService() *ImageJobService {
	return &ImageJobService{}
}

func (*ImageJobService) CheckJobCondition(c *gin.Context) {
	var job models.EcpJobCreateReq
	if err := c.ShouldBindJSON(&job); err != nil {
		logs.GetLogger().Errorf("failed to parse json, error: %v", err)
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.JsonError))
		return
	}
	logs.GetLogger().Infof("check job condition, received Data: %+v", job)

	checkPriceFlag, totalCost, err := checkPrice(job.Price, job.Duration, job.Resource)
	if err != nil {
		logs.GetLogger().Errorf("failed to check price, job_uuid: %s, error: %v", job.UUID, err)
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.JsonError))
		return
	}

	if !checkPriceFlag {
		logs.GetLogger().Errorf("bid below the set price, job_uuid: %s, pid: %s, need: %0.4f", job.UUID, job.Price, totalCost)
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.BelowPriceError))
		return
	}

	receive, _, _, _, err := checkResourceForImage(job.Resource)
	if receive {
		c.JSON(http.StatusOK, util.CreateSuccessResponse("success"))
	} else {
		c.JSON(http.StatusOK, util.CreateSuccessResponse(util.NoAvailableResourcesError))
	}
	return
}

func (*ImageJobService) DeployJob(c *gin.Context) {
	var job models.EcpJobCreateReq
	if err := c.ShouldBindJSON(&job); err != nil {
		logs.GetLogger().Errorf("failed to parse json, error: %v", err)
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.JsonError))
		return
	}
	logs.GetLogger().Infof("Job received Data: %+v", job)

	checkPriceFlag, totalCost, err := checkPrice(job.Price, job.Duration, job.Resource)
	if err != nil {
		logs.GetLogger().Errorf("failed to check price, job_uuid: %s, error: %v", job.UUID, err)
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.JsonError))
		return
	}

	if !checkPriceFlag {
		logs.GetLogger().Errorf("bid below the set price, job_uuid: %s, pid: %s, need: %0.4f", job.UUID, job.Price, totalCost)
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.BelowPriceError))
		return
	}

	if strings.TrimSpace(job.Name) == "" {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.UbiTaskParamError, "missing required field: name"))
		return
	}

	if strings.TrimSpace(job.Image) == "" {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.UbiTaskParamError, "missing required field: image"))
		return
	}

	if err := NewDockerService().PullImage(job.Image); err != nil {
		logs.GetLogger().Errorf("failed to pull %s image, error: %v", job.Image, err)
		return
	}

	var env []string
	for k, v := range job.Envs {
		env = append(env, fmt.Sprintf("%s=%s", k, v))
	}

	_, _, needCpu, _, err := checkResourceForImage(job.Resource)

	if err := NewDockerService().PullImage(job.Image); err != nil {
		logs.GetLogger().Errorf("failed to pull %s image, error: %v", job.Image, err)
		return
	}

	var needResource container.Resources
	needResource = container.Resources{
		CPUQuota: needCpu * 100000,
		Memory:   job.Resource.Memory,
		DeviceRequests: []container.DeviceRequest{
			{
				Driver:       "nvidia",
				Count:        -1,
				DeviceIDs:    []string{},
				Capabilities: [][]string{{"gpu"}},
				Options:      nil,
			},
		},
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

	containerName := job.Name + "-" + generateString(5)
	dockerService := NewDockerService()
	if err := dockerService.ContainerCreateAndStart(containerConfig, hostConfig, containerName); err != nil {
		logs.GetLogger().Errorf("failed to create job container, job_uuid: %s, error: %v", job.UUID, err)
		return
	}
	logs.GetLogger().Warnf("job_uuid: %s, starting container, container name: %s", job.UUID, containerName)

	time.Sleep(3 * time.Second)
	if !dockerService.IsExistContainer(containerName) {
		logs.GetLogger().Warnf("job_uuid: %s, not found container", job.UUID)
		return
	}
	logs.GetLogger().Warnf("job_uuid: %s, started container, container name: %s", job.UUID, containerName)

	if err = NewEcpJobService().SaveEcpJobEntity(&models.EcpJobEntity{
		Uuid:          job.UUID,
		Name:          job.Name,
		Image:         job.Image,
		Env:           strings.Join(env, ","),
		ContainerName: containerName,
		CreateTime:    time.Now().Unix(),
	}); err != nil {
		logs.GetLogger().Errorf("failed to save job to db, error: %v", err)
		return
	}
	c.JSON(http.StatusOK, util.CreateSuccessResponse("success"))
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
		if status, ok := containerStatus[entity.ContainerName]; ok {
			fmt.Printf("container name: %s, status: %s \n", entity.ContainerName, status)
			result = append(result, models.EcpJobStatusResp{Uuid: entity.Uuid, Status: status})
		}
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

func checkPrice(userPrice string, duration int, resource models.HardwareResource) (bool, float64, error) {
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
	cpuCost := float64(resource.CPU) * cpuPrice * float64(duration/3600)
	memoryCost := formatGiB(resource.Memory) * memoryPrice * float64(duration/3600)
	storageCost := formatGiB(resource.Storage) * storagePrice * float64(duration/3600)
	gpuCost := float64(resource.GPU) * gpuPrice * float64(duration/3600)

	totalCost := cpuCost + memoryCost + storageCost + gpuCost

	// Compare user's price with total cost
	return userPayPrice >= totalCost, totalCost, nil
}

func checkResourceForImage(resource models.HardwareResource) (bool, string, int64, int64, error) {
	dockerService := NewDockerService()
	containerLogStr, err := dockerService.ContainerLogs("resource-exporter")
	if err != nil {
		return false, "", 0, 0, err
	}

	var nodeResource models.NodeResource
	if err := json.Unmarshal([]byte(containerLogStr), &nodeResource); err != nil {
		return false, "", 0, 0, err
	}

	needCpu := resource.CPU
	var needMemory, needStorage float64
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

	var gpuMap = make(map[string]int)
	if nodeResource.Gpu.AttachedGpus > 0 {
		for _, detail := range nodeResource.Gpu.Details {
			if detail.Status == models.Available {
				gpuMap[detail.ProductName] += 1
			}
		}
	}

	logs.GetLogger().Infof("checkResourceForImage: needCpu: %d, needMemory: %.2f, needStorage: %.2f, needGpu: %d, gpuName: %s", needCpu, needMemory, needStorage, resource.GPU, resource.GPUModel)
	logs.GetLogger().Infof("checkResourceForImage: remainingCpu: %d, remainingMemory: %.2f, remainingStorage: %.2f, remainingGpu: %+v", remainderCpu, remainderMemory, remainderStorage, gpuMap)
	if needCpu <= remainderCpu && needMemory <= remainderMemory && needStorage <= remainderStorage {
		if resource.GPUModel != "" {
			var flag bool
			for k, num := range gpuMap {
				if strings.ToUpper(k) == resource.GPUModel && num > 0 {
					flag = true
					break
				}
			}
			if flag {
				return true, nodeResource.CpuName, needCpu, int64(needMemory), nil
			} else {
				return false, nodeResource.CpuName, needCpu, int64(needMemory), nil
			}
		}
		return true, nodeResource.CpuName, needCpu, int64(needMemory), nil
	}
	return false, nodeResource.CpuName, needCpu, int64(needMemory), nil
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

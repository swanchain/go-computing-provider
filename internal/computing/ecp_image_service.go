package computing

import (
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

	var needResource container.Resources
	needResource = container.Resources{
		DeviceRequests: []container.DeviceRequest{
			{
				Driver:       "nvidia",
				Count:        -1,
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
	ecpJobs, err := NewEcpJobService().GetEcpJobs()
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
	memoryCost := float64(resource.Memory) * memoryPrice * float64(duration/3600)
	storageCost := float64(resource.Storage) * storagePrice * float64(duration/3600)
	gpuCost := float64(resource.GPU) * gpuPrice * float64(duration/3600)

	totalCost := cpuCost + memoryCost + storageCost + gpuCost

	// Compare user's price with total cost
	return userPayPrice >= totalCost, totalCost, nil
}

func parsePrice(priceStr string) (float64, error) {
	return strconv.ParseFloat(priceStr, 64)
}

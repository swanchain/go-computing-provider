package computing

import (
	"fmt"
	"github.com/docker/docker/api/types/container"
	"github.com/filswan/go-swan-lib/logs"
	"github.com/gin-gonic/gin"
	"github.com/swanchain/go-computing-provider/internal/models"
	"github.com/swanchain/go-computing-provider/util"
	"net/http"
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
	for k, v := range job.EnvVar {
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
		logs.GetLogger().Errorf("failed to create job container, job_uuid: %s, error: %v", job.Uuid, err)
		return
	}
	logs.GetLogger().Warnf("job_uuid: %s, starting container, container name: %s", job.Uuid, containerName)

	time.Sleep(3 * time.Second)
	if !dockerService.IsExistContainer(containerName) {
		logs.GetLogger().Warnf("job_uuid: %s, not found container", job.Uuid)
		return
	}
	logs.GetLogger().Warnf("job_uuid: %s, started container, container name: %s", job.Uuid, containerName)

	err := NewEcpJobService().SaveEcpJobEntity(&models.EcpJobEntity{
		Uuid:          job.Uuid,
		Name:          job.Name,
		Image:         job.Image,
		Env:           strings.Join(env, ","),
		ContainerName: containerName,
		CreateTime:    time.Now().Unix(),
	})
	if err != nil {
		logs.GetLogger().Errorf("failed to save job to db, error: %v", err)
		return
	}
	c.JSON(http.StatusOK, util.CreateSuccessResponse("success"))
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
	c.JSON(http.StatusOK, util.CreateSuccessResponse("success"))
}

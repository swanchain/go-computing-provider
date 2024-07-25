package computing

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/filswan/go-swan-lib/logs"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/swanchain/go-computing-provider/build"
	"github.com/swanchain/go-computing-provider/conf"
	"github.com/swanchain/go-computing-provider/constants"
	"github.com/swanchain/go-computing-provider/internal/contract"
	"github.com/swanchain/go-computing-provider/internal/contract/account"
	"github.com/swanchain/go-computing-provider/internal/contract/ecp"
	"github.com/swanchain/go-computing-provider/internal/models"
	"github.com/swanchain/go-computing-provider/util"
	"github.com/swanchain/go-computing-provider/wallet"
	"io"
	batchv1 "k8s.io/api/batch/v1"
	coreV1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func DoUbiTaskForK8s(c *gin.Context) {

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

	if ubiTask.ResourceType != models.RESOURCE_TYPE_CPU && ubiTask.ResourceType != models.RESOURCE_TYPE_GPU {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.UbiTaskParamError, "the value of resource_type is 0 or 1"))
		return
	}
	if ubiTask.Type == 0 {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.UbiTaskParamError, "missing required field: type"))
		return
	}

	if strings.TrimSpace(ubiTask.InputParam) == "" {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.UbiTaskParamError, "missing required field: input_param"))
		return
	}

	if strings.TrimSpace(ubiTask.VerifyParam) == "" {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.UbiTaskParamError, "missing required field: verify_param"))
		return
	}

	if strings.TrimSpace(ubiTask.CheckCode) == "" {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.UbiTaskParamError, "missing required field: check_code"))
		return
	}

	if strings.TrimSpace(ubiTask.Signature) == "" {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.UbiTaskParamError, "missing required field: signature"))
		return
	}

	if ubiTask.DeadLine == 0 {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.UbiTaskParamError, "missing required field: deadline"))
		return
	}

	cpAccountAddress, err := contract.GetCpAccountAddress()
	if err != nil {
		logs.GetLogger().Errorf("get cp account contract address failed, error: %v", err)
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.GetCpAccountError))
		return
	}

	signature, err := verifySignature(conf.GetConfig().UBI.UbiEnginePk, fmt.Sprintf("%s%d", cpAccountAddress, ubiTask.ID), ubiTask.Signature)
	if err != nil {
		logs.GetLogger().Errorf("verifySignature for ubi task failed, error: %+v", err)
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.UbiTaskParamError, "sign data failed"))
		return
	}

	logs.GetLogger().Infof("ubi task sign verifing, task_id: %d, type: %s, verify: %v", ubiTask.ID, models.UbiTaskTypeStr(ubiTask.Type), signature)
	if !signature {
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.UbiTaskParamError, "signature verify failed"))
		return
	}

	var gpuFlag = "0"
	if ubiTask.ResourceType == 1 {
		gpuFlag = "1"
	}

	var taskEntity = new(models.TaskEntity)
	taskEntity.Id = int64(ubiTask.ID)
	taskEntity.Type = ubiTask.Type
	taskEntity.Name = ubiTask.Name
	taskEntity.ResourceType = ubiTask.ResourceType
	taskEntity.InputParam = ubiTask.InputParam
	taskEntity.VerifyParam = ubiTask.VerifyParam
	taskEntity.Status = models.TASK_RECEIVED_STATUS
	taskEntity.CreateTime = time.Now().Unix()
	taskEntity.Deadline = ubiTask.DeadLine
	taskEntity.CheckCode = ubiTask.CheckCode
	err = NewTaskService().SaveTaskEntity(taskEntity)
	if err != nil {
		logs.GetLogger().Errorf("save task entity failed, error: %v", err)
		return
	}

	var envFilePath string
	envFilePath = filepath.Join(os.Getenv("CP_PATH"), "fil-c2.env")
	envVars, err := godotenv.Read(envFilePath)
	if err != nil {
		logs.GetLogger().Errorf("reading fil-c2-env.env failed, error: %v", err)
		return
	}

	c2GpuConfig := envVars["RUST_GPU_TOOLS_CUSTOM_GPU"]
	c2GpuName := convertGpuName(strings.TrimSpace(c2GpuConfig))
	nodeName, architecture, needCpu, needMemory, needStorage, err := checkResourceAvailableForUbi(ubiTask.ResourceType, c2GpuName, ubiTask.Resource)
	if err != nil {
		taskEntity.Status = models.TASK_FAILED_STATUS
		NewTaskService().SaveTaskEntity(taskEntity)
		logs.GetLogger().Errorf("check resource failed, error: %v", err)
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.CheckResourcesError))
		return
	}

	if nodeName == "" {
		taskEntity.Status = models.TASK_FAILED_STATUS
		taskEntity.Error = "No resources available"
		NewTaskService().SaveTaskEntity(taskEntity)
		logs.GetLogger().Warnf("ubi task id: %d, type: %s, not found a resources available", ubiTask.ID, models.GetResourceTypeStr(ubiTask.ResourceType))
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.NoAvailableResourcesError))
		return
	}

	var ubiTaskImage string
	if architecture == constants.CPU_AMD {
		ubiTaskImage = build.UBITaskImageAmdCpu
		if gpuFlag == "1" {
			ubiTaskImage = build.UBITaskImageAmdGpu
		}
	} else if architecture == constants.CPU_INTEL {
		ubiTaskImage = build.UBITaskImageIntelCpu
		if gpuFlag == "1" {
			ubiTaskImage = build.UBITaskImageIntelGpu
		}
	}

	mem := strings.Split(strings.TrimSpace(ubiTask.Resource.Memory), " ")[1]
	memUnit := strings.ReplaceAll(mem, "B", "")
	disk := strings.Split(strings.TrimSpace(ubiTask.Resource.Storage), " ")[1]
	diskUnit := strings.ReplaceAll(disk, "B", "")
	memQuantity, err := resource.ParseQuantity(fmt.Sprintf("%d%s", needMemory, memUnit))
	if err != nil {
		taskEntity.Status = models.TASK_FAILED_STATUS
		NewTaskService().SaveTaskEntity(taskEntity)
		logs.GetLogger().Error("get memory failed, error: %+v", err)
		return
	}

	storageQuantity, err := resource.ParseQuantity(fmt.Sprintf("%d%s", needStorage, diskUnit))
	if err != nil {
		taskEntity.Status = models.TASK_FAILED_STATUS
		NewTaskService().SaveTaskEntity(taskEntity)
		logs.GetLogger().Error("get storage failed, error: %+v", err)
		return
	}

	maxMemQuantity, err := resource.ParseQuantity(fmt.Sprintf("%d%s", needMemory*2, memUnit))
	if err != nil {
		taskEntity.Status = models.TASK_FAILED_STATUS
		NewTaskService().SaveTaskEntity(taskEntity)
		logs.GetLogger().Error("get memory failed, error: %+v", err)
		return
	}

	maxStorageQuantity, err := resource.ParseQuantity(fmt.Sprintf("%d%s", needStorage*2, diskUnit))
	if err != nil {
		taskEntity.Status = models.TASK_FAILED_STATUS
		NewTaskService().SaveTaskEntity(taskEntity)
		logs.GetLogger().Error("get storage failed, error: %+v", err)
		return
	}

	resourceRequirements := coreV1.ResourceRequirements{
		Limits: coreV1.ResourceList{
			coreV1.ResourceCPU:              *resource.NewQuantity(needCpu*2, resource.DecimalSI),
			coreV1.ResourceMemory:           maxMemQuantity,
			coreV1.ResourceEphemeralStorage: maxStorageQuantity,
			"nvidia.com/gpu":                resource.MustParse(gpuFlag),
		},
		Requests: coreV1.ResourceList{
			coreV1.ResourceCPU:              *resource.NewQuantity(needCpu, resource.DecimalSI),
			coreV1.ResourceMemory:           memQuantity,
			coreV1.ResourceEphemeralStorage: storageQuantity,
			"nvidia.com/gpu":                resource.MustParse(gpuFlag),
		},
	}

	go func() {
		var namespace = "ubi-task-" + strconv.Itoa(ubiTask.ID)
		var err error
		defer func() {
			if err := recover(); err != nil {
				logs.GetLogger().Errorf("do zk task painc, error: %+v", err)
				return
			}

			ubiTaskRun, err := NewTaskService().GetTaskEntity(int64(ubiTask.ID))
			if err != nil {
				logs.GetLogger().Errorf("get ubi task detail from db failed, ubiTaskId: %d, error: %+v", ubiTask.ID, err)
				return
			}

			if ubiTaskRun.TxHash != "" {
				ubiTaskRun.Status = models.TASK_SUBMITTED_STATUS
			} else {
				ubiTaskRun.Status = models.TASK_FAILED_STATUS
				k8sService := NewK8sService()
				k8sService.k8sClient.CoreV1().Namespaces().Delete(context.TODO(), namespace, metaV1.DeleteOptions{})
			}
			err = NewTaskService().SaveTaskEntity(ubiTaskRun)
		}()

		if ubiTaskImage == "" {
			logs.GetLogger().Errorf("please check the log output of the resource-exporter pod to see if cpu_name is intel or amd")
			return
		}

		k8sService := NewK8sService()
		if _, err = k8sService.GetNameSpace(context.TODO(), namespace, metaV1.GetOptions{}); err != nil {
			if errors.IsNotFound(err) {
				k8sNamespace := &v1.Namespace{
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

		receiveUrl := fmt.Sprintf("%s:%d/api/v1/computing/cp/receive/ubi", k8sService.GetAPIServerEndpoint(), conf.GetConfig().API.Port)
		execCommand := []string{"ubi-bench", "c2"}
		JobName := strings.ToLower(models.UbiTaskTypeStr(ubiTask.Type)) + "-" + strconv.Itoa(ubiTask.ID)
		filC2Param := envVars["FIL_PROOFS_PARAMETER_CACHE"]
		if gpuFlag == "0" {
			delete(envVars, "RUST_GPU_TOOLS_CUSTOM_GPU")
			envVars["BELLMAN_NO_GPU"] = "1"
		}

		delete(envVars, "FIL_PROOFS_PARAMETER_CACHE")
		var useEnvVars []v1.EnvVar
		for k, v := range envVars {
			useEnvVars = append(useEnvVars, v1.EnvVar{
				Name:  k,
				Value: v,
			})
		}

		useEnvVars = append(useEnvVars, v1.EnvVar{
			Name:  "RECEIVE_PROOF_URL",
			Value: receiveUrl,
		},
			v1.EnvVar{
				Name:  "TASKID",
				Value: strconv.Itoa(ubiTask.ID),
			},
			v1.EnvVar{
				Name:  "NAME_SPACE",
				Value: namespace,
			},
			v1.EnvVar{
				Name:  "PARAM_URL",
				Value: ubiTask.InputParam,
			},
		)

		job := &batchv1.Job{
			ObjectMeta: metaV1.ObjectMeta{
				Name:      JobName,
				Namespace: namespace,
			},
			Spec: batchv1.JobSpec{
				Template: v1.PodTemplateSpec{
					Spec: v1.PodSpec{
						NodeName:     nodeName,
						NodeSelector: generateLabel(strings.ReplaceAll(c2GpuName, " ", "-")),
						Containers: []v1.Container{
							{
								Name:  JobName + generateString(5),
								Image: ubiTaskImage,
								Env:   useEnvVars,
								VolumeMounts: []v1.VolumeMount{
									{
										Name:      "proof-params",
										MountPath: "/var/tmp/filecoin-proof-parameters",
									},
								},
								Command:         execCommand,
								Resources:       resourceRequirements,
								ImagePullPolicy: coreV1.PullIfNotPresent,
							},
						},
						Volumes: []v1.Volume{
							{
								Name: "proof-params",
								VolumeSource: v1.VolumeSource{
									HostPath: &v1.HostPathVolumeSource{
										Path: filC2Param,
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
		*job.Spec.TTLSecondsAfterFinished = 120

		if _, err = k8sService.k8sClient.BatchV1().Jobs(namespace).Create(context.TODO(), job, metaV1.CreateOptions{}); err != nil {
			logs.GetLogger().Errorf("Failed creating ubi task job: %v", err)
			return
		}

		err = wait.PollImmediate(2*time.Second, 60*time.Second, func() (bool, error) {
			pods, err := k8sService.k8sClient.CoreV1().Pods(namespace).List(context.TODO(), metaV1.ListOptions{
				LabelSelector: fmt.Sprintf("job-name=%s", JobName),
			})
			if err != nil {
				return false, err
			}

			for _, p := range pods.Items {
				for _, condition := range p.Status.Conditions {
					if condition.Type != coreV1.PodReady && condition.Status != coreV1.ConditionTrue {
						return false, nil
					}
				}
			}
			return true, nil
		})
		if err != nil {
			logs.GetLogger().Errorf("Failed waiting pods create: %v", err)
			return
		}

		pods, err := k8sService.k8sClient.CoreV1().Pods(namespace).List(context.TODO(), metaV1.ListOptions{
			LabelSelector: fmt.Sprintf("job-name=%s", JobName),
		})
		if err != nil {
			logs.GetLogger().Errorf("Failed list ubi pods: %v", err)
			return
		}

		var podName string
		for _, pod := range pods.Items {
			podName = pod.Name
			break
		}
		if podName == "" {
			return
		}

		req := k8sService.k8sClient.CoreV1().Pods(namespace).GetLogs(podName, &v1.PodLogOptions{
			Container: "",
			Follow:    true,
		})

		podLogs, err := req.Stream(context.Background())
		if err != nil {
			logs.GetLogger().Errorf("Error opening log stream: %v", err)
			return
		}
		defer podLogs.Close()

		cpRepoPath, _ := os.LookupEnv("CP_PATH")
		ubiLogFileName := filepath.Join(cpRepoPath, "ubi-fcp.log")
		logFile, err := os.OpenFile(ubiLogFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			logs.GetLogger().Errorf("opening ubi-fcp log file failed, error: %v", err)
			return
		}
		defer logFile.Close()

		if _, err = io.Copy(logFile, podLogs); err != nil {
			logs.GetLogger().Errorf("write ubi-fcp log to file failed, error: %v", err)
			return
		}
	}()

	c.JSON(http.StatusOK, util.CreateSuccessResponse("success"))
}

func ReceiveUbiProof(c *gin.Context) {
	var c2Proof models.UbiC2Proof
	var err error
	if err := c.ShouldBindJSON(&c2Proof); err != nil {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.JsonError))
		return
	}
	logs.GetLogger().Infof("task_id: %s, C2 proof out received: %+v", c2Proof.TaskId, c2Proof)

	taskId, err := strconv.Atoi(c2Proof.TaskId)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.JsonError))
		return
	}

	ubiTask, err := NewTaskService().GetTaskEntity(int64(taskId))
	if err != nil {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.JsonError))
		return
	}
	go func() {
		defer func() {
			if err := recover(); err != nil {
				logs.GetLogger().Errorf("taskId: %d, submit zk-task proof catch painc error: %v", taskId, err)
			}
		}()
		err = submitUBIProof(c2Proof, ubiTask)
		if err != nil {
			return
		}
	}()

	c.JSON(http.StatusOK, util.CreateSuccessResponse("success"))
}

func DoUbiTaskForDocker(c *gin.Context) {
	var ubiTask models.UBITaskReq
	if err := c.ShouldBindJSON(&ubiTask); err != nil {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.JsonError))
		return
	}

	logs.GetLogger().Infof("ubi task received: id: %d, deadline: %d,resource_type: %d, type: %s, input_param: %s, signature: %s",
		ubiTask.ID, ubiTask.DeadLine, ubiTask.ResourceType, models.UbiTaskTypeStr(ubiTask.Type), ubiTask.InputParam, ubiTask.Signature)

	if ubiTask.ID == 0 {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.UbiTaskParamError, "missing required field: id"))
		return
	}
	if strings.TrimSpace(ubiTask.Name) == "" {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.UbiTaskParamError, "missing required field: name"))
		return
	}

	if ubiTask.ResourceType != models.RESOURCE_TYPE_CPU && ubiTask.ResourceType != models.RESOURCE_TYPE_GPU {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.UbiTaskParamError, "the value of resource_type is 0 or 1"))
		return
	}
	if ubiTask.Type == 0 {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.UbiTaskParamError, "missing required field: type"))
		return
	}

	if strings.TrimSpace(ubiTask.InputParam) == "" {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.UbiTaskParamError, "missing required field: input_param"))
		return
	}

	if strings.TrimSpace(ubiTask.VerifyParam) == "" {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.UbiTaskParamError, "missing required field: verify_param"))
		return
	}

	if strings.TrimSpace(ubiTask.CheckCode) == "" {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.UbiTaskParamError, "missing required field: check_code"))
		return
	}

	//if strings.TrimSpace(ubiTask.Signature) == "" {
	//	c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.UbiTaskParamError, "missing required field: signature"))
	//	return
	//}

	if ubiTask.DeadLine == 0 {
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.UbiTaskParamError, "missing required field: deadline"))
		return
	}

	//cpAccountAddress, err := contract.GetCpAccountAddress()
	//if err != nil {
	//	logs.GetLogger().Errorf("get cp account contract address failed, error: %v", err)
	//	c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.GetCpAccountError))
	//	return
	//}

	//signature, err := verifySignature(conf.GetConfig().UBI.UbiEnginePk, fmt.Sprintf("%s%d", cpAccountAddress, ubiTask.ID), ubiTask.Signature)
	//if err != nil {
	//	logs.GetLogger().Errorf("verifySignature for ubi task failed, error: %+v", err)
	//	c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.SignatureError, "verify sign data occur error"))
	//	return
	//}
	//
	//logs.GetLogger().Infof("ubi task sign verifing, task_id: %d, verify: %v", ubiTask.ID, signature)
	//if !signature {
	//	c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.SignatureError, "signature verify failed"))
	//	return
	//}

	var gpuFlag = "0"
	if ubiTask.ResourceType == 1 {
		gpuFlag = "1"
	}

	var taskEntity = new(models.TaskEntity)
	taskEntity.Id = int64(ubiTask.ID)
	taskEntity.Type = ubiTask.Type
	taskEntity.Name = ubiTask.Name
	taskEntity.ResourceType = ubiTask.ResourceType
	taskEntity.InputParam = ubiTask.InputParam
	taskEntity.VerifyParam = ubiTask.VerifyParam
	taskEntity.Status = models.TASK_RECEIVED_STATUS
	taskEntity.CreateTime = time.Now().Unix()
	taskEntity.Deadline = ubiTask.DeadLine
	taskEntity.CheckCode = ubiTask.CheckCode
	err := NewTaskService().SaveTaskEntity(taskEntity)
	if err != nil {
		logs.GetLogger().Errorf("save task entity failed, error: %v", err)
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.SaveTaskEntityError))
		return
	}

	var gpuName string
	gpuConfig, ok := os.LookupEnv("RUST_GPU_TOOLS_CUSTOM_GPU")
	if ok {
		gpuName = convertGpuName(strings.TrimSpace(gpuConfig))
	}

	_, architecture, _, needMemory, err := checkResourceForUbi(ubiTask.Resource, gpuName, ubiTask.ResourceType)
	if err != nil {
		taskEntity.Status = models.TASK_FAILED_STATUS
		NewTaskService().SaveTaskEntity(taskEntity)
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.CheckResourcesError))
		return
	}

	var ubiTaskImage string
	if architecture == constants.CPU_AMD {
		ubiTaskImage = build.UBITaskImageAmdCpu
		if gpuFlag == "1" {
			ubiTaskImage = build.UBITaskImageAmdGpu
		}
	} else if architecture == constants.CPU_INTEL {
		ubiTaskImage = build.UBITaskImageIntelCpu
		if gpuFlag == "1" {
			ubiTaskImage = build.UBITaskImageIntelGpu
		}
	}

	go func() {
		defer func() {
			if err := recover(); err != nil {
				logs.GetLogger().Errorf("do zk task painc, error: %+v", err)
				return
			}
		}()

		if ubiTaskImage == "" {
			logs.GetLogger().Errorf("please check the log output of the resource-exporter container to see if cpu_name is intel or amd")
			return
		}

		if err := NewDockerService().PullImage(ubiTaskImage); err != nil {
			logs.GetLogger().Errorf("pull %s image failed, error: %v", ubiTaskImage, err)
			return
		}

		receiveUrl := fmt.Sprintf("http://127.0.0.1:%d/api/v1/computing/cp/docker/receive/ubi", conf.GetConfig().API.Port)
		execCommand := []string{"ubi-bench", "c2"}
		JobName := strings.ToLower(models.UbiTaskTypeStr(ubiTask.Type)) + "-" + strconv.Itoa(ubiTask.ID)

		var env = []string{"RECEIVE_PROOF_URL=" + receiveUrl}
		env = append(env, "TASKID="+strconv.Itoa(ubiTask.ID))
		env = append(env, "PARAM_URL="+ubiTask.InputParam)

		var needResource container.Resources
		if gpuFlag == "0" {
			env = append(env, "BELLMAN_NO_GPU=1")
			needResource = container.Resources{
				Memory: needMemory * 1024 * 1024 * 1024,
			}
		} else {
			gpuEnv, ok := os.LookupEnv("RUST_GPU_TOOLS_CUSTOM_GPU")
			if ok {
				env = append(env, "RUST_GPU_TOOLS_CUSTOM_GPU="+gpuEnv)
			}
			needResource = container.Resources{
				Memory: needMemory * 1024 * 1024 * 1024,
				DeviceRequests: []container.DeviceRequest{
					{
						Driver:       "nvidia",
						Count:        -1,
						Capabilities: [][]string{{"gpu"}},
						Options:      nil,
					},
				},
			}
		}

		filC2Param, ok := os.LookupEnv("FIL_PROOFS_PARAMETER_CACHE")
		if !ok {
			filC2Param = "/var/tmp/filecoin-proof-parameters"
		}

		hostConfig := &container.HostConfig{
			Binds:       []string{fmt.Sprintf("%s:/var/tmp/filecoin-proof-parameters", filC2Param)},
			Resources:   needResource,
			NetworkMode: network.NetworkHost,
		}
		containerConfig := &container.Config{
			Image:        ubiTaskImage,
			Cmd:          execCommand,
			Env:          env,
			AttachStdout: true,
			AttachStderr: true,
			Tty:          true,
		}

		containerName := JobName + generateString(5)
		dockerService := NewDockerService()
		if err = dockerService.ContainerCreateAndStart(containerConfig, hostConfig, containerName); err != nil {
			logs.GetLogger().Errorf("create ubi task container failed, error: %v", err)
			return
		}

		time.Sleep(3 * time.Second)
		if !dockerService.IsExistContainer(containerName) {
			return
		}

		containerLogStream, err := dockerService.GetContainerLogStream(containerName)
		if err != nil {
			logs.GetLogger().Errorf("get docker container log stream failed, error: %v", err)
			return
		}
		defer containerLogStream.Close()

		cpRepoPath, _ := os.LookupEnv("CP_PATH")
		ubiLogFileName := filepath.Join(cpRepoPath, "ubi-ecp.log")
		logFile, err := os.OpenFile(ubiLogFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			logs.GetLogger().Errorf("opening ubi-ecp log file failed, error: %v", err)
			return
		}
		defer logFile.Close()

		if _, err = io.Copy(logFile, containerLogStream); err != nil {
			logs.GetLogger().Errorf("write ubi-ecp log to file failed, error: %v", err)
			return
		}
	}()
	c.JSON(http.StatusOK, util.CreateSuccessResponse("success"))
}

func checkResourceForUbi(resource *models.TaskResource, gpuName string, resourceType int) (bool, string, int64, int64, error) {
	dockerService := NewDockerService()
	containerLogStr, err := dockerService.ContainerLogs("resource-exporter")
	if err != nil {
		return false, "", 0, 0, err
	}

	var nodeResource models.NodeResource
	if err := json.Unmarshal([]byte(containerLogStr), &nodeResource); err != nil {
		return false, "", 0, 0, err
	}

	needCpu, _ := strconv.ParseInt(resource.CPU, 10, 64)
	var needMemory, needStorage float64
	if len(strings.Split(strings.TrimSpace(resource.Memory), " ")) > 0 {
		needMemory, err = strconv.ParseFloat(strings.Split(strings.TrimSpace(resource.Memory), " ")[0], 64)

	}
	if len(strings.Split(strings.TrimSpace(resource.Storage), " ")) > 0 {
		needStorage, err = strconv.ParseFloat(strings.Split(strings.TrimSpace(resource.Storage), " ")[0], 64)
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

	logs.GetLogger().Infof("checkResourceForUbi: needCpu: %d, needMemory: %.2f, needStorage: %.2f", needCpu, needMemory, needStorage)
	logs.GetLogger().Infof("checkResourceForUbi: remainingCpu: %d, remainingMemory: %.2f, remainingStorage: %.2f, remainingGpu: %+v", remainderCpu, remainderMemory, remainderStorage, gpuMap)
	if needCpu <= remainderCpu && needMemory <= remainderMemory && needStorage <= remainderStorage {
		if resourceType == 1 {
			if gpuName != "" {
				var flag bool
				for k, num := range gpuMap {
					if strings.ToUpper(k) == gpuName && num > 0 {
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
		}
		return true, nodeResource.CpuName, needCpu, int64(needMemory), nil
	}
	return false, nodeResource.CpuName, needCpu, int64(needMemory), nil
}

func GetCpResource(c *gin.Context) {
	location, err := getLocation()
	if err != nil {
		logs.GetLogger().Error(err)
		location = "-"
	}

	dockerService := NewDockerService()
	containerLogStr, err := dockerService.ContainerLogs("resource-exporter")
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.ServerError, err.Error()))
		return
	}

	var nodeResource models.NodeResource
	if err := json.Unmarshal([]byte(containerLogStr), &nodeResource); err != nil {
		logs.GetLogger().Warnf("hardware info parse to json failed, restarting resource-exporter")
		if err = RestartResourceExporter(); err != nil {
			logs.GetLogger().Errorf("restartResourceExporter failed, error: %v", err)
		}
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.JsonError))
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
		ClusterInfo:      []*models.NodeResource{&nodeResource},
		NodeName:         conf.GetConfig().API.NodeName,
		NodeId:           GetNodeId(cpRepo),
		CpAccountAddress: cpAccountAddress,
	})
}

func submitUBIProof(c2Proof models.UbiC2Proof, task *models.TaskEntity) error {
	chainUrl, err := conf.GetRpcByNetWorkName()
	if err != nil {
		logs.GetLogger().Errorf("get rpc url failed, taskId: %s, error: %v", c2Proof.TaskId, err)
		return err
	}
	client, err := ethclient.Dial(chainUrl)
	if err != nil {
		logs.GetLogger().Errorf("dial rpc connect failed, taskId: %s, error: %v", c2Proof.TaskId, err)
		return err
	}
	client.Close()

	localWallet, err := wallet.SetupWallet(wallet.WalletRepo)
	if err != nil {
		logs.GetLogger().Errorf("setup wallet failed, taskId: %s,error: %v", c2Proof.TaskId, err)
		return err
	}

	_, workerAddress, err := GetOwnerAddressAndWorkerAddress()
	if err != nil {
		logs.GetLogger().Errorf("get worker address failed, taskId: %s,error: %v", c2Proof.TaskId, err)
		return err
	}

	ki, err := localWallet.FindKey(workerAddress)
	if err != nil || ki == nil {
		logs.GetLogger().Errorf("taskId: %s,the address: %s, private key %v", c2Proof.TaskId, workerAddress, wallet.ErrKeyInfoNotFound)
		return err
	}
	var workerPrivateKey = ki.PrivateKey
	ki = nil

	taskStub, err := ecp.NewTaskStub(client, ecp.WithTaskContractAddress(task.Contract), ecp.WithTaskPrivateKey(workerPrivateKey))
	if err != nil {
		logs.GetLogger().Errorf("create ubi task client failed, taskId: %s, contract: %s, error: %v", c2Proof.TaskId, task.Contract, err)
		return err
	}

	var blockNumber uint64

loopTask:
	for {
		select {
		case <-time.After(30 * time.Second):
			logs.GetLogger().Errorf("get ubi task info, taskId: %s timeout", c2Proof.TaskId)
			break loopTask
		default:
			blockNumber, err = client.BlockNumber(context.Background())
			if err != nil {
				logs.GetLogger().Warnf("get ubi task info failed, taskId: %s, msg: %s, retrying", c2Proof.TaskId, err.Error())
				time.Sleep(3 * time.Second)
				continue
			} else {
				break loopTask
			}
		}
	}

	remainingTime := task.Deadline - int64(blockNumber)
	if remainingTime < 0 {
		logs.GetLogger().Warnf("taskId: %s proof submission deadline has passed,current: %d, deadline: %d, , deadlineTime: %d", c2Proof.TaskId, blockNumber, task.Deadline, remainingTime)
		task.Status = models.TASK_FAILED_STATUS
		task.Error = fmt.Sprintf("create contract deadline has passed")
		return NewTaskService().SaveTaskEntity(task)
	}

	if conf.GetConfig().UBI.AggregateCommits {
		if err = submitTaskToSequencer(c2Proof.Proof, task, remainingTime); err != nil {
			task.Status = models.TASK_FAILED_STATUS
		} else {
			task.Status = models.TASK_SUBMITTED_STATUS
		}
	} else {
		taskContractAddress, err := taskStub.CreateTaskContract(c2Proof.Proof, task, remainingTime)
		if taskContractAddress != "" {
			task.Status = models.TASK_SUBMITTED_STATUS
			task.Contract = taskContractAddress
			logs.GetLogger().Infof("taskId: %s, taskContractAddress: %s", c2Proof.TaskId, taskContractAddress)
		} else if err != nil {
			task.Status = models.TASK_FAILED_STATUS
			task.Error = fmt.Sprintf("%s", err.Error())
			logs.GetLogger().Errorf("taskId: %s, create task contract failed, error: %v", c2Proof.TaskId, err)
		}
	}
	return NewTaskService().SaveTaskEntity(task)
}

func GetTaskInfoOnChain(taskContract string) (models.EcpTaskInfo, error) {
	var taskInfo models.EcpTaskInfo

	chainRpc, err := conf.GetRpcByNetWorkName()
	if err != nil {
		return taskInfo, err
	}
	client, err := ethclient.Dial(chainRpc)
	if err != nil {
		return taskInfo, err
	}
	defer client.Close()

	taskStub, err := ecp.NewTaskStub(client, ecp.WithTaskContractAddress(taskContract))
	if err != nil {
		logs.GetLogger().Errorf("create ubi task client failed, error: %v", err)
		return taskInfo, err
	}

loopTask:
	for {
		select {
		case <-time.After(10 * time.Second):
			logs.GetLogger().Errorf("get ubi task info, contract address: %s, timeout", taskContract)
			break loopTask
		default:
			taskInfo, err = taskStub.GetTaskInfo()
			if err != nil {
				logs.GetLogger().Warnf("get ubi task info failed, contract address: %s, retrying", taskContract)
				time.Sleep(time.Second)
				continue
			} else {
				break loopTask
			}
		}
	}
	if taskInfo.InputParam != "" {
		return taskInfo, nil
	}
	return taskInfo, err
}

func reportClusterResourceForDocker() {
	dockerService := NewDockerService()
	containerLogStr, err := dockerService.ContainerLogs("resource-exporter")
	if err != nil {
		if err = RestartResourceExporter(); err != nil {
			logs.GetLogger().Errorf("restartResourceExporter failed, error: %v", err)
		}
		return
	}

	var nodeResource models.NodeResource
	if err := json.Unmarshal([]byte(containerLogStr), &nodeResource); err != nil {
		if err = RestartResourceExporter(); err != nil {
			logs.GetLogger().Errorf("restartResourceExporter failed, error: %v", err)
		}
		return
	}

	var freeGpuMap = make(map[string]int)
	if nodeResource.Gpu.AttachedGpus > 0 {
		for _, g := range nodeResource.Gpu.Details {
			if g.Status == models.Available {
				freeGpuMap[g.ProductName] += 1
			} else {
				freeGpuMap[g.ProductName] = 0
			}
		}
	}
	logs.GetLogger().Infof("collect hardware resource, freeCpu:%s, freeMemory: %s, freeStorage: %s, freeGpu: %v",
		nodeResource.Cpu.Free, nodeResource.Memory.Free, nodeResource.Storage.Free, freeGpuMap)
}

func CronTaskForEcp() {
	go func() {
		ticker := time.NewTicker(10 * time.Minute)
		for range ticker.C {
			NewDockerService().CleanResource()
		}
	}()

	go func() {
		ticker := time.NewTicker(3 * time.Minute)
		for range ticker.C {
			reportClusterResourceForDocker()
		}
	}()

	go func() {
		ticker := time.NewTicker(5 * time.Minute)
		for range ticker.C {
			var taskList []models.TaskEntity
			oneHourAgo := time.Now().Add(-1 * time.Hour).Unix()
			err := NewTaskService().Model(&models.TaskEntity{}).Where("status !=? and status !=? and create_time <?", models.TASK_SUBMITTED_STATUS, models.TASK_FAILED_STATUS, oneHourAgo).
				Or("tx_hash !='' and status =?", models.TASK_FAILED_STATUS).Find(&taskList).Error
			if err != nil {
				logs.GetLogger().Errorf("Failed get task list, error: %+v", err)
				return
			}

			for _, entity := range taskList {
				ubiTask := entity
				if ubiTask.TxHash != "" {
					ubiTask.Status = models.TASK_SUBMITTED_STATUS
				} else {
					ubiTask.Status = models.TASK_FAILED_STATUS
				}
				NewTaskService().SaveTaskEntity(&ubiTask)
			}
		}
	}()

	go func() {
		defer func() {
			if err := recover(); err != nil {
				logs.GetLogger().Errorf("GetUbiTaskReward, error: %+v", err)
			}
		}()
		ticker := time.NewTicker(3 * time.Minute)
		for range ticker.C {

			taskList, err := NewTaskService().GetTaskListNoReward()
			if err != nil {
				logs.GetLogger().Errorf("failed to get task list, error: %+v", err)
				return
			}

			taskGroups := handleTasksToGroup(taskList)
			for _, group := range taskGroups {
				taskList, err := NewSequencer().QueryTask(group.Ids...)
				if err != nil {
					logs.GetLogger().Errorf("failed to query task, task ids: %v, error: %v", group.Ids, err)
					continue
				}

				var taskMap = make(map[int64]SequenceTask)
				for _, t := range taskList.List {
					taskMap[int64(t.Id)] = t
				}

				for _, item := range group.Items {
					if t, ok := taskMap[item.Id]; ok {
						item.Reward = t.Reward
						item.SettlementCid = t.SettlementCid
						item.SequenceCid = t.SequenceCid
						item.Sequencer = 1

						var status int
						switch t.Status {
						case "Verified":
							status = models.TASK_VERIFIED_STATUS
						case "rewarded":
							status = models.TASK_REWARDED_STATUS
						case "NSC":
							status = models.TASK_NSC_STATUS
						}
						item.Status = status
						NewTaskService().UpdateTaskEntityByTaskId(item)
					}
				}
			}
		}
	}()
}

func SyncCpAccountInfo() {
	cpAccountAddress, err := contract.GetCpAccountAddress()
	if err != nil {
		logs.GetLogger().Fatalf("get cp account contract address failed, error: %v", err)
		return
	}

	cpAccount, err := account.GetAccountInfo()
	if err != nil {
		logs.GetLogger().Errorf("get cpAccount failed, error: %v", err)
		return
	}

	var cpInfo = new(models.CpInfoEntity)
	cpInfo.NodeId = cpAccount.NodeId
	cpInfo.OwnerAddress = cpAccount.OwnerAddress
	cpInfo.Beneficiary = cpAccount.Beneficiary
	cpInfo.WorkerAddress = cpAccount.WorkerAddress
	cpInfo.ContractAddress = cpAccountAddress
	cpInfo.CreateAt = time.Now().Format("2006-01-02 15:04:05")
	cpInfo.UpdateAt = time.Now().Format("2006-01-02 15:04:05")
	cpInfo.MultiAddresses = cpAccount.MultiAddresses
	cpInfo.Version = cpAccount.Version
	cpInfo.TaskTypes = cpAccount.TaskTypes
	if err = NewCpInfoService().SaveCpInfoEntity(cpInfo); err != nil {
		logs.GetLogger().Errorf("save cp info to db failed, error: %v", err)
		return
	}
}

func RestartResourceExporter() error {
	resourceExporterContainerName := "resource-exporter"
	dockerService := NewDockerService()
	dockerService.RemoveContainerByName(resourceExporterContainerName)
	err := dockerService.PullImage(build.UBIResourceExporterDockerImage)
	if err != nil {
		return fmt.Errorf("pull %s image failed, error: %v", build.UBIResourceExporterDockerImage, err)
	}

	err = dockerService.ContainerCreateAndStart(&container.Config{
		Image:        build.UBIResourceExporterDockerImage,
		AttachStdout: true,
		AttachStderr: true,
		Tty:          true,
	}, nil, resourceExporterContainerName)
	if err != nil {
		return fmt.Errorf("create resource-exporter container failed, error: %v", err)
	}
	return nil
}

func submitTaskToSequencer(proof string, task *models.TaskEntity, timeOut int64) error {
	var err error
	var taskReq struct {
		Id           int64  `json:"id"`
		Type         int    `json:"type"`
		ResourceType int    `json:"resource_type"`
		InputParam   string `json:"input_param"`
		VerifyParam  string `json:"verify_param"`
		Deadline     int64  `json:"deadline"`
		CheckCode    string `json:"check_code"`
		Proof        string `json:"proof"`
	}

	taskReq.Id = task.Id
	taskReq.Type = task.Type
	taskReq.ResourceType = task.ResourceType
	taskReq.InputParam = task.InputParam
	taskReq.VerifyParam = task.VerifyParam
	taskReq.Deadline = task.Deadline
	taskReq.CheckCode = task.CheckCode
	taskReq.Proof = proof

	data, err := json.Marshal(&taskReq)
	if err != nil {
		return err
	}

	timeOutCh := time.After(time.Second * time.Duration(timeOut))
outerLoop:
	for {
		select {
		case <-timeOutCh:
			err = fmt.Errorf("submit task to sequencer timed out")
			break outerLoop
		default:
			sendTaskProof, err := NewSequencer().SendTaskProof(data)
			if err != nil {
				logs.GetLogger().Warnf("taskId: %d submit task to sequencer failed, error: %v, retrying", task.Id, err)
				time.Sleep(2 * time.Second)
				continue
			}
			task.BlockHash = sendTaskProof.BlockHash
			task.Sign = sendTaskProof.Sign
			task.Sequencer = 1
			break outerLoop
		}
	}

	return err
}

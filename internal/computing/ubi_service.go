package computing

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
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
	"math"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func DoUbiTaskForK8s(c *gin.Context) {
	if !conf.GetConfig().UBI.EnableSequencer && !conf.GetConfig().UBI.AutoChainProof {
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.RejectZkTaskError))
		return
	}

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

	balance, err := checkBalance(cpAccountAddress)
	if err != nil || !balance {
		logs.GetLogger().Errorf("failed check cp account balance, error: %v", err)
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.CheckBalanceError))
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
	c2GpuName = strings.ToUpper(c2GpuName)
	nodeName, architecture, needCpu, needMemory, needStorage, gpuIndex, err := checkResourceAvailableForUbi(ubiTask.ResourceType, c2GpuName, ubiTask.Resource)
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

			if ubiTaskRun.Contract != "" || ubiTaskRun.BlockHash != "" {
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

		if len(strings.TrimSpace(filC2Param)) == 0 {
			logs.GetLogger().Warnf("task_id: %d, `FIL_PROOFS_PARAMETER_CACHE` variable is not configured", ubiTask.ID)
			return
		}

		if gpuFlag == "0" {
			delete(envVars, "RUST_GPU_TOOLS_CUSTOM_GPU")
			envVars["BELLMAN_NO_GPU"] = "1"
		} else {
			if len(gpuIndex) > 0 {
				envVars["CUDA_VISIBLE_DEVICES"] = gpuIndex[0]
			}
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
						RestartPolicy: v1.RestartPolicyNever,
					},
				},
				BackoffLimit:            new(int32),
				TTLSecondsAfterFinished: new(int32),
			},
		}

		*job.Spec.BackoffLimit = 1
		*job.Spec.TTLSecondsAfterFinished = 300

		if _, err = k8sService.k8sClient.BatchV1().Jobs(namespace).Create(context.TODO(), job, metaV1.CreateOptions{}); err != nil {
			logs.GetLogger().Errorf("Failed creating ubi task job: %v", err)
			return
		}

		err = wait.PollImmediate(2*time.Second, 60*time.Second, func() (bool, error) {
			pods, err := k8sService.k8sClient.CoreV1().Pods(namespace).List(context.TODO(), metaV1.ListOptions{
				LabelSelector: fmt.Sprintf("job-name=%s", JobName),
			})
			if err != nil {
				logs.GetLogger().Errorf("failed get pod, taskId: %d, error: %v，retrying", ubiTask.ID, err)
				return false, err
			}

			if len(pods.Items) == 0 {
				return false, nil
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
			logs.GetLogger().Errorf("failed get pod name, taskId: %d", ubiTask.ID)
			return
		}
		logs.GetLogger().Infof("successfully get pod name, taskId: %d, podName: %s", ubiTask.ID, podName)

		NewTaskService().UpdateTaskStatusById(ubiTask.ID, models.TASK_RUNNING_STATUS)
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
	logs.GetLogger().Infof("task_id: %s, The ZK proof computation is complete: %+v", c2Proof.TaskId, c2Proof)

	taskId, err := strconv.Atoi(c2Proof.TaskId)
	if err != nil {
		logs.GetLogger().Errorf("failed to parse task id, task_id: %s, error: %v", c2Proof.TaskId, err)
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.JsonError))
		return
	}

	ubiTask, err := NewTaskService().GetTaskEntity(int64(taskId))
	if err != nil {
		logs.GetLogger().Errorf("failed to get task info, task_id: %s, error: %v", c2Proof.TaskId, err)
		c.JSON(http.StatusBadRequest, util.CreateErrorResponse(util.JsonError))
		return
	}
	go func() {
		defer func() {
			if err := recover(); err != nil {
				logs.GetLogger().Errorf("taskId: %d, submit zk-task proof catch painc error: %v", taskId, err)
			}
		}()
		submitUBIProof(c2Proof, ubiTask)
	}()

	c.JSON(http.StatusOK, util.CreateSuccessResponse("success"))
}

func DoUbiTaskForDocker(c *gin.Context) {
	if !conf.GetConfig().UBI.EnableSequencer && !conf.GetConfig().UBI.AutoChainProof {
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.RejectZkTaskError))
		return
	}

	var ubiTask models.UBITaskReq
	if err := c.ShouldBindJSON(&ubiTask); err != nil {
		logs.GetLogger().Error("failed to parse json, error: %v", err)
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

	balance, err := checkBalance(cpAccountAddress)
	if err != nil || !balance {
		logs.GetLogger().Errorf("failed check cp account balance, error: %v", err)
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.CheckBalanceError))
		return
	}

	signature, err := verifySignature(conf.GetConfig().UBI.UbiEnginePk, fmt.Sprintf("%s%d", cpAccountAddress, ubiTask.ID), ubiTask.Signature)
	if err != nil {
		logs.GetLogger().Errorf("verifySignature for ubi task failed, error: %+v", err)
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.SignatureError, "verify sign data occur error"))
		return
	}

	logs.GetLogger().Infof("ubi task sign verifing, task_id: %d, verify: %v", ubiTask.ID, signature)
	if !signature {
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.SignatureError, "signature verify failed"))
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
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.SaveTaskEntityError))
		return
	}

	var gpuName string
	gpuConfig, ok := os.LookupEnv("RUST_GPU_TOOLS_CUSTOM_GPU")
	if ok {
		gpuName = convertGpuName(strings.TrimSpace(gpuConfig))
	}

	_, architecture, _, needMemory, indexs, err := checkResourceForUbi(ubiTask.Resource, gpuName, ubiTask.ResourceType)
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
			logs.GetLogger().Errorf("failed to pull %s image, error: %v", ubiTaskImage, err)
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

			var useIndexs []string
			if len(indexs) > 0 {
				useIndexs = append(useIndexs, indexs[0])
				env = append(env, fmt.Sprintf("CUDA_VISIBLE_DEVICES=%s", strings.Join(useIndexs, ",")))
			} else {
				logs.GetLogger().Warnf("not resources available, task_id: %d", ubiTask.ID)
				return
			}
			needResource = container.Resources{
				Memory: needMemory * 1024 * 1024 * 1024,
				DeviceRequests: []container.DeviceRequest{
					{
						Driver:       "nvidia",
						DeviceIDs:    useIndexs,
						Capabilities: [][]string{{"gpu", "compute", "utility"}},
					},
				},
			}
		}

		filC2Param, ok := os.LookupEnv("FIL_PROOFS_PARAMETER_CACHE")
		if !ok {
			filC2Param = "/var/tmp/filecoin-proof-parameters"
		}
		if len(strings.TrimSpace(filC2Param)) == 0 {
			logs.GetLogger().Warnf("task_id: %d, `FIL_PROOFS_PARAMETER_CACHE` variable is not configured", ubiTask.ID)
			return
		}

		hostConfig := &container.HostConfig{
			Binds:       []string{fmt.Sprintf("%s:/var/tmp/filecoin-proof-parameters", filC2Param)},
			Resources:   needResource,
			NetworkMode: network.NetworkHost,
			Privileged:  true,
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
		logs.GetLogger().Warnf("task_id: %d, starting container, container name: %s", ubiTask.ID, containerName)

		dockerService := NewDockerService()
		if err = dockerService.ContainerCreateAndStart(containerConfig, hostConfig, nil, containerName); err != nil {
			logs.GetLogger().Errorf("failed to create ubi task container, task_id: %d, error: %v", ubiTask.ID, err)
			return
		}

		time.Sleep(3 * time.Second)
		if !dockerService.IsExistContainer(containerName) {
			logs.GetLogger().Warnf("task_id: %d, not found container", ubiTask.ID)
			return
		}
		logs.GetLogger().Warnf("task_id: %d, started container, container name: %s", ubiTask.ID, containerName)
		NewTaskService().UpdateTaskStatusById(ubiTask.ID, models.TASK_RUNNING_STATUS)

		containerLogStream, err := dockerService.GetContainerLogStream(context.TODO(), containerName)
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

func checkResourceForUbi(resource *models.TaskResource, gpuName string, resourceType int) (bool, string, int64, int64, []string, error) {
	dockerService := NewDockerService()
	containerLogStr, err := dockerService.ContainerLogs("resource-exporter")
	if err != nil {
		return false, "", 0, 0, nil, err
	}

	var nodeResource models.NodeResource
	if err := json.Unmarshal([]byte(containerLogStr), &nodeResource); err != nil {
		return false, "", 0, 0, nil, err
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

	type gpuData struct {
		num    int
		indexs []string
	}

	var indexs []string
	var gpuMap = make(map[string]gpuData)
	if nodeResource.Gpu.AttachedGpus > 0 {
		for _, detail := range nodeResource.Gpu.Details {
			if detail.Status == models.Available {
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

	logs.GetLogger().Infof("checkResourceForUbi: needCpu: %d, needMemory: %.2f, needStorage: %.2f", needCpu, needMemory, needStorage)
	logs.GetLogger().Infof("checkResourceForUbi: remainingCpu: %d, remainingMemory: %.2f, remainingStorage: %.2f, remainingGpu: %+v", remainderCpu, remainderMemory, remainderStorage, gpuMap)
	if needCpu <= remainderCpu && needMemory <= remainderMemory && needStorage <= remainderStorage {
		if resourceType == 1 {
			if gpuName != "" {
				var flag bool
				for k, gm := range gpuMap {
					if strings.ToUpper(k) == gpuName && gm.num > 0 {
						indexs = gm.indexs
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
		}
		return true, nodeResource.CpuName, needCpu, int64(needMemory), indexs, nil
	}
	return false, nodeResource.CpuName, needCpu, int64(needMemory), indexs, nil
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

	list, err := NewEcpJobService().GetEcpJobList([]string{models.CreatedStatus, models.RunningStatus})
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(util.JsonError))
		return
	}

	var taskGpuMap = make(map[string][]string)
	for _, g := range list {
		if len(strings.TrimSpace(g.GpuIndex)) > 0 {
			taskGpuMap[g.GpuName] = append(taskGpuMap[g.GpuName], strings.Split(strings.TrimSpace(g.GpuIndex), ",")...)
		}
	}

	if nodeResource.Gpu.AttachedGpus > 0 {
		for i, detail := range nodeResource.Gpu.Details {
			if detail.Status == models.Available {
				if checkGpu(detail.ProductName, detail.Index, taskGpuMap) {
					nodeResource.Gpu.Details[i].Status = models.Occupied
				}
			}
		}
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
		ClientType:       "ECP",
	})
}

func submitUBIProof(c2Proof models.UbiC2Proof, task *models.TaskEntity) {
	chainUrl, err := conf.GetRpcByNetWorkName()
	if err != nil {
		logs.GetLogger().Errorf("failed to get rpc url, taskId: %s, error: %v", c2Proof.TaskId, err)
		return
	}
	client, err := contract.GetEthClient(chainUrl)
	if err != nil {
		logs.GetLogger().Errorf("failed to dial rpc, taskId: %s, error: %v", c2Proof.TaskId, err)
		return
	}
	defer client.Close()

	var timeUnit int64 = 2
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		logs.GetLogger().Errorf("dial rpc connect failed, taskId: %s, error: %v", c2Proof.TaskId, err)
		return
	}
	if chainId.Int64() == 254 {
		timeUnit = 5
	}

	localWallet, err := wallet.SetupWallet(wallet.WalletRepo)
	if err != nil {
		logs.GetLogger().Errorf("failed to setup wallet, taskId: %s,error: %v", c2Proof.TaskId, err)
		return
	}

	_, workerAddress, err := GetOwnerAddressAndWorkerAddress()
	if err != nil {
		logs.GetLogger().Errorf("failed get worker address, taskId: %s,error: %v", c2Proof.TaskId, err)
		return
	}

	ki, err := localWallet.FindKey(workerAddress)
	if err != nil || ki == nil {
		logs.GetLogger().Errorf("taskId: %s,the address: %s, private key %v", c2Proof.TaskId, workerAddress, wallet.ErrKeyInfoNotFound)
		return
	}
	var workerPrivateKey = ki.PrivateKey
	ki = nil

	taskStub, err := ecp.NewTaskStub(client, ecp.WithTaskContractAddress(task.Contract), ecp.WithTaskPrivateKey(workerPrivateKey))
	if err != nil {
		logs.GetLogger().Errorf("create ubi task client failed, taskId: %s, contract: %s, error: %v", c2Proof.TaskId, task.Contract, err)
		return
	}

	var sequencerBalance float64
	if conf.GetConfig().UBI.EnableSequencer {
		cpAccountAddress, err := contract.GetCpAccountAddress()
		if err != nil {
			logs.GetLogger().Errorf("failed to get cp account contract address, error: %v", err)
			return
		}

		err = RetryFn(func() error {
			sequencerStub, err := ecp.NewSequencerStub(client, ecp.WithSequencerCpAccountAddress(cpAccountAddress))
			if err != nil {
				logs.GetLogger().Errorf("failed to get cp sequencer contract, error: %v", err)
				return err
			}
			sequencerBalanceStr, err := sequencerStub.GetCPBalance()
			if err != nil {
				logs.GetLogger().Errorf("failed to get cp sequencer contract, error: %v", err)
				return err
			}

			sequencerBalance, err = strconv.ParseFloat(sequencerBalanceStr, 64)
			if err != nil {
				logs.GetLogger().Errorf("failed to convert numbers for cp sequencer balance, sequencerBalance: %s, error: %v", sequencerBalanceStr, err)
				return err
			}
			return nil
		}, 3, 2*time.Second)
		if err != nil {
			return
		}
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

	task.Proof = c2Proof.Proof
	remainingTime := (task.Deadline - int64(blockNumber)) * timeUnit
	if remainingTime < 0 {
		logs.GetLogger().Warnf("taskId: %s proof submission deadline has passed, current: %d, deadline: %d, deadlineTime: %d", c2Proof.TaskId, blockNumber, task.Deadline, remainingTime)
		task.Status = models.TASK_FAILED_STATUS
		task.Error = fmt.Sprintf("create contract deadline has passed")
		if err = NewTaskService().SaveTaskEntity(task); err != nil {
			logs.GetLogger().Errorf("failed to save task info, taskId: %s, error: %v", c2Proof.TaskId, err)
		}
		return
	}

	if conf.GetConfig().UBI.EnableSequencer && conf.GetConfig().UBI.AutoChainProof {
		if sequencerBalance <= 0 {
			logs.GetLogger().Infof("taskId: %s starting to create task contract", c2Proof.TaskId)
			taskContractAddress, err := taskStub.CreateTaskContract(c2Proof.Proof, task, remainingTime)
			if taskContractAddress != "" {
				task.Status = models.TASK_SUBMITTED_STATUS
				task.Contract = taskContractAddress
				task.Sequencer = 0
				logs.GetLogger().Infof("successfully submitted to the chain, taskId: %s task contract address: %s", c2Proof.TaskId, taskContractAddress)
			} else {
				task.Status = models.TASK_FAILED_STATUS
				task.Error = fmt.Sprintf("%s", err.Error())
				logs.GetLogger().Errorf("taskId: %s, failed to create task contract, error: %v", c2Proof.TaskId, err)
			}
		} else {
			logs.GetLogger().Infof("taskId: %s starting to use sequencer to submit proof", c2Proof.TaskId)
			if err = submitTaskToSequencer(c2Proof.Proof, task, remainingTime, true); err != nil {
				logs.GetLogger().Errorf("failed to submitted to the sequencer, taskId: %d, error: %v", task.Id, err)
				task.Status = models.TASK_FAILED_STATUS
			} else {
				task.Status = models.TASK_SUBMITTED_STATUS
			}
		}
	} else if conf.GetConfig().UBI.EnableSequencer && !conf.GetConfig().UBI.AutoChainProof {
		if sequencerBalance > 0 {
			logs.GetLogger().Infof("taskId: %s starting to use sequencer to submit proof", c2Proof.TaskId)
			if err = submitTaskToSequencer(c2Proof.Proof, task, remainingTime, false); err != nil {
				logs.GetLogger().Errorf("failed to submitted to the sequencer, taskId: %d, error: %v", task.Id, err)
				task.Status = models.TASK_FAILED_STATUS
			} else {
				task.Status = models.TASK_SUBMITTED_STATUS
			}
		} else {
			task.Status = models.TASK_FAILED_STATUS
			logs.GetLogger().Warnf("taskId: %d, sequencer insufficient balance, sequencerBalance: %.6f", task.Id, sequencerBalance)
		}
	} else {
		logs.GetLogger().Infof("taskId: %s starting to create task contract", c2Proof.TaskId)
		taskContractAddress, err := taskStub.CreateTaskContract(c2Proof.Proof, task, remainingTime)
		if taskContractAddress != "" {
			task.Status = models.TASK_SUBMITTED_STATUS
			task.Contract = taskContractAddress
			task.Sequencer = 0
			logs.GetLogger().Infof("taskId: %s, taskContractAddress: %s", c2Proof.TaskId, taskContractAddress)
		} else if err != nil {
			task.Status = models.TASK_FAILED_STATUS
			task.Error = fmt.Sprintf("%s", err.Error())
			logs.GetLogger().Errorf("taskId: %s, failed to create task contract, error: %v", c2Proof.TaskId, err)
		}
	}

	if err = NewTaskService().SaveTaskEntity(task); err != nil {
		logs.GetLogger().Errorf("failed to save task info, taskId: %s, error: %v", c2Proof.TaskId, err)
	}
	return
}

func GetTaskInfoOnChain(taskContract string) (models.EcpTaskInfo, error) {
	var taskInfo models.EcpTaskInfo

	chainRpc, err := conf.GetRpcByNetWorkName()
	if err != nil {
		return taskInfo, err
	}
	client, err := contract.GetEthClient(chainRpc)
	if err != nil {
		return taskInfo, err
	}
	defer client.Close()

	taskStub, err := ecp.NewTaskStub(client, ecp.WithTaskContractAddress(taskContract))
	if err != nil {
		logs.GetLogger().Errorf("faile to get ubi task client, error: %v", err)
		return taskInfo, err
	}

	taskInfo, err = taskStub.GetTaskInfo()
	if taskInfo.InputParam != "" {
		return taskInfo, nil
	}
	return taskInfo, err
}

func GetAggregatedTaskInfo(taskContract string) (string, error) {
	chainRpc, err := conf.GetRpcByNetWorkName()
	if err != nil {
		return "", err
	}
	client, err := contract.GetEthClient(chainRpc)
	if err != nil {
		return "", err
	}
	defer client.Close()

	aggregatedTask, err := ecp.NewAggregatedTask(common.HexToAddress(taskContract), client)
	if err != nil {
		return "", fmt.Errorf("create ubi task client failed, error: %v", err)
	}
	blobCid, err := aggregatedTask.TaskBlobCID(&bind.CallOpts{})
	if err != nil {
		return "", fmt.Errorf("failed to get cid from contact, error: %v", err)
	}
	return blobCid, nil
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
		logs.GetLogger().Errorf("failed to convert json, container log: %s, error: %v", containerLogStr, err)
		if err = RestartResourceExporter(); err != nil {
			logs.GetLogger().Errorf("restartResourceExporter failed, error: %v", err)
		}
		return
	}

	list, err := NewEcpJobService().GetEcpJobList([]string{models.CreatedStatus, models.RunningStatus})
	if err != nil {
		logs.GetLogger().Errorf("failed to get ecp job task, error: %v", err)
		return
	}

	var taskGpuMap = make(map[string][]string)
	for _, g := range list {
		if len(strings.TrimSpace(g.GpuIndex)) > 0 {
			taskGpuMap[g.GpuName] = append(taskGpuMap[g.GpuName], strings.Split(strings.TrimSpace(g.GpuIndex), ",")...)
		}
	}

	var freeGpuMap = make(map[string]int)
	var useGpuMap = make(map[string]int)
	if nodeResource.Gpu.AttachedGpus > 0 {
		for _, g := range nodeResource.Gpu.Details {
			if g.Status == models.Available {
				if checkGpu(g.ProductName, g.Index, taskGpuMap) {
					freeGpuMap[g.ProductName] = 0
				} else {
					freeGpuMap[g.ProductName] += 1
				}
			} else {
				useGpuMap[g.ProductName] += 1
			}
		}
	}
	logs.GetLogger().Infof("collect hardware resource, freeCpu:%s, freeMemory: %s, freeStorage: %s, freeGpu: %v, useGpu: %v",
		nodeResource.Cpu.Free, nodeResource.Memory.Free, nodeResource.Storage.Free, freeGpuMap, useGpuMap)
}

func updateEcpTaskStatus() {
	list, err := NewEcpJobService().GetEcpJobList([]string{models.CreatedStatus, models.RunningStatus})
	if err != nil {
		logs.GetLogger().Errorf("failed to get ecp job task, error: %v", err)
		return
	}

	containerStatus, err := NewDockerService().GetContainerStatus()
	if err != nil {
		return
	}

	for _, entity := range list {
		if time.Now().Unix()-entity.CreateTime < 3600 {
			continue
		}
		if status, ok := containerStatus[entity.ContainerName]; ok {
			NewEcpJobService().UpdateEcpJobEntity(entity.Uuid, status)
		} else {
			NewEcpJobService().UpdateEcpJobEntity(entity.Uuid, models.TerminatedStatus)
		}
	}
}

func CronTaskForEcp() {
	if conf.GetConfig().API.AutoDeleteImage {
		go func() {
			ticker := time.NewTicker(2 * time.Hour)
			for range ticker.C {
				NewDockerService().CleanResourceForDocker()
			}
		}()
	}

	go func() {
		ticker := time.NewTicker(3 * time.Minute)
		for range ticker.C {
			reportClusterResourceForDocker()
		}
	}()

	go func() {
		ticker := time.NewTicker(30 * time.Second)
		for range ticker.C {
			updateEcpTaskStatus()
		}
	}()

	go func() {
		ticker := time.NewTicker(5 * time.Minute)
		for range ticker.C {
			var taskList []models.TaskEntity
			oneHourAgo := time.Now().Add(-1 * time.Hour).Unix()
			err := NewTaskService().Model(&models.TaskEntity{}).Where("status in (?,?) and create_time <?", models.TASK_RECEIVED_STATUS, models.TASK_RUNNING_STATUS, oneHourAgo).
				Or("tx_hash !='' and status =?", models.TASK_FAILED_STATUS).Find(&taskList).Error
			if err != nil {
				logs.GetLogger().Errorf("Failed get task list, error: %+v", err)
				return
			}

			for _, entity := range taskList {
				ubiTask := entity
				if ubiTask.Contract != "" || ubiTask.BlockHash != "" {
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
		ticker := time.NewTicker(10 * time.Minute)
		for range ticker.C {
			if err := syncTaskStatusForSequencerService(); err != nil {
				logs.GetLogger().Errorf("failed to sync task from sequencer, error: %v", err)
			}
		}
	}()

	go func() {
		defer func() {
			if err := recover(); err != nil {
				logs.GetLogger().Errorf("Scanner task payment events, error: %+v", err)
			}
		}()

		ticker := time.NewTicker(30 * time.Minute)
		for range ticker.C {
			NewTaskPaymentService().ScannerChainGetTaskPayment()
		}
	}()

	go func() {
		ticker := time.NewTicker(30 * time.Minute)
		for range ticker.C {
			GetCpBalance()
		}
	}()

}

func syncTaskStatusForSequencerService() error {
	taskList, err := NewTaskService().GetTaskListNoReward()
	if err != nil {
		return fmt.Errorf("failed to get task list, error: %+v", err)
	}

	taskGroups := handleTasksToGroup(taskList)
	var taskIdAndStatus = make(map[int64]string)
	for _, group := range taskGroups {
		taskList, err := NewSequencer().QueryTask(group.Type, group.Ids...)
		if err != nil {
			logs.GetLogger().Errorf("failed to query task, task ids: %v, error: %v", group.Ids, err)
			continue
		}

		var taskMap = make(map[int64]SequenceTask)
		for _, t := range taskList.Data.List {
			taskMap[int64(t.Id)] = t
		}

		for _, item := range group.Items {
			var taskAddress = item.SequenceTaskAddr
			if t, ok := taskMap[item.Id]; ok {
				item.Reward = "0.00"
				item.SequenceCid = t.SequenceCid
				item.SettlementCid = t.SettlementCid
				item.SequenceTaskAddr = t.SequenceTaskAddr
				item.SettlementTaskAddr = t.SettlementTaskAddr

				var status int
				if t.SequenceCid == "-1" {
					status = models.TASK_NSC_STATUS
				} else if t.SettlementTaskAddr != "" {
					status = models.TASK_REWARDED_STATUS
					if len(t.Reward) >= 4 {
						item.Reward = t.Reward[:4]
					} else {
						item.Reward = t.Reward
					}
				} else {
					switch t.Status {
					case "received":
						status = models.TASK_RECEIVED_STATUS
					case "running":
						status = models.TASK_RUNNING_STATUS
					case "submitted":
						status = models.TASK_SUBMITTED_STATUS
					case "verified":
						status = models.TASK_VERIFIED_STATUS
					case "rewarded":
						status = models.TASK_REWARDED_STATUS
					case "invalid":
						status = models.TASK_INVALID_STATUS
					case "repeated":
						status = models.TASK_REPEATED_STATUS
					case "timeout":
						status = models.TASK_TIMEOUT_STATUS
					case "verifyFailed":
						status = models.TASK_VERIFYFAILED_STATUS
					case "failed":
						status = models.TASK_FAILED_STATUS
					default:
						status = models.TASK_UNKNOWN_STATUS
					}
				}

				if item.Status == status && taskAddress == item.SequenceTaskAddr {
					continue
				}
				taskIdAndStatus[item.Id] = models.TaskStatusStr(status)
				item.Status = status
				NewTaskService().UpdateTaskEntityByTaskId(item)
			}
		}
	}
	if len(taskIdAndStatus) > 0 {
		logs.GetLogger().Infof("successfully updated the task status: %v", taskIdAndStatus)
	}
	return nil
}

func SyncCpAccountInfo() (*models.Account, error) {
	cpAccountAddress, err := contract.GetCpAccountAddress()
	if err != nil {
		logs.GetLogger().Errorf("get cp account contract address failed, error: %v", err)
		return nil, err
	}

	cpAccount, err := account.GetAccountInfo()
	if err != nil {
		logs.GetLogger().Errorf("get cpAccount failed, error: %v", err)
		return nil, err
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
		return nil, err
	}
	return &cpAccount, nil
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
	}, &container.HostConfig{
		RestartPolicy: container.RestartPolicy{
			Name:              container.RestartPolicyOnFailure,
			MaximumRetryCount: 3,
		},
		Privileged: true,
	}, nil, resourceExporterContainerName)
	if err != nil {
		return fmt.Errorf("create resource-exporter container failed, error: %v", err)
	}
	return nil
}

func submitTaskToSequencer(proof string, task *models.TaskEntity, timeOut int64, autoChainProof bool) error {
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

	start := time.Now()
	if autoChainProof {
		var flag bool
		for i := 0; i < 5; i++ {
			sendTaskProof, err := NewSequencer().SendTaskProof(data)
			if err != nil {
				logs.GetLogger().Warnf("taskId: %d submit task to sequencer failed, error: %v, retrying", task.Id, err)
				time.Sleep(2 * time.Second)
				continue
			}
			task.BlockHash = sendTaskProof.Data.BlockHash
			task.Sign = sendTaskProof.Data.Sign
			task.Sequencer = 1
			logs.GetLogger().Infof("successfully submitted to the sequencer, taskId: %d the sequencer receipt is block_hash: %s, sign: %s", task.Id, task.BlockHash, task.Sign)
			flag = true
			break
		}

		if flag {
			return nil
		}
		remainingTime := timeOut - int64(time.Now().Sub(start).Seconds())
		if !flag && remainingTime > 0 {
			chainUrl, err := conf.GetRpcByNetWorkName()
			if err != nil {
				return fmt.Errorf("failed to get rpc url, taskId: %d, error: %v", task.Id, err)
			}
			client, err := contract.GetEthClient(chainUrl)
			if err != nil {
				return fmt.Errorf("failed to dial rpc, taskId: %d, error: %v", task.Id, err)
			}
			client.Close()

			localWallet, err := wallet.SetupWallet(wallet.WalletRepo)
			if err != nil {
				return fmt.Errorf("failed to setup wallet, taskId: %d,error: %v", task.Id, err)
			}

			_, workerAddress, err := GetOwnerAddressAndWorkerAddress()
			if err != nil {
				return fmt.Errorf("failed to get worker address, taskId: %d,error: %v", task.Id, err)
			}

			ki, err := localWallet.FindKey(workerAddress)
			if err != nil || ki == nil {
				return fmt.Errorf("taskId: %d,the address: %s, private key %v", task.Id, workerAddress, wallet.ErrKeyInfoNotFound)
			}
			var workerPrivateKey = ki.PrivateKey
			ki = nil
			taskStub, err := ecp.NewTaskStub(client, ecp.WithTaskContractAddress(task.Contract), ecp.WithTaskPrivateKey(workerPrivateKey))
			if err != nil {
				return fmt.Errorf("failed to create ubi task client, taskId: %s, contract: %s, error: %v", task.Id, task.Contract, err)
			}

			logs.GetLogger().Infof("taskId: %d, failed to submit proof using sequencer, starting to create task contract", task.Id)
			taskContractAddress, err := taskStub.CreateTaskContract(task.Proof, task, remainingTime)
			if taskContractAddress != "" {
				task.Status = models.TASK_SUBMITTED_STATUS
				task.Contract = taskContractAddress
				task.Sequencer = 0
				logs.GetLogger().Infof("successfully submitted to the chain, taskId: %d task contract address: %s", task.Id, taskContractAddress)
				return nil
			} else {
				task.Status = models.TASK_FAILED_STATUS
				task.Error = fmt.Sprintf("%s", err.Error())
				return fmt.Errorf("taskId: %d, failed to create task contract, error: %v", task.Id, err)
			}
		}

		if remainingTime <= 0 {
			return fmt.Errorf("taskId: %d, proof submission deadline has passed, remainingTime: %d", task.Id, remainingTime)
		}
	} else {
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
				task.BlockHash = sendTaskProof.Data.BlockHash
				task.Sign = sendTaskProof.Data.Sign
				task.Sequencer = 1
				logs.GetLogger().Infof("successfully submitted to the sequencer, taskId: %d, the sequencer receipt is block_hash: %s, sign: %s", task.Id, task.BlockHash, task.Sign)
				break outerLoop
			}
		}

	}
	return err
}

func checkBalance(cpAccountAddress string) (bool, error) {
	cpBalance, err := NewCpBalanceService().GetCpBalance(cpAccountAddress)
	if err != nil || cpBalance == nil || cpBalance.WorkerBalance == 0 {
		return false, fmt.Errorf("failed to get cp balance, error: %v", err)
	}
	logs.GetLogger().Infof("cpAccount: %s, sequencer balance: %f ETH, worker address balance: %f ETH", cpAccountAddress, cpBalance.SequencerBalance, cpBalance.WorkerBalance)

	if conf.GetConfig().UBI.EnableSequencer && !conf.GetConfig().UBI.AutoChainProof {
		if cpBalance.SequencerBalance > 0.000001 {
			return true, nil
		} else {
			return false, fmt.Errorf("No sufficient balance in the sequencer account, current balance: %f ETH", cpBalance.SequencerBalance)
		}
	}

	if conf.GetConfig().UBI.EnableSequencer && conf.GetConfig().UBI.AutoChainProof {
		if cpBalance.SequencerBalance > 0.000001 || cpBalance.WorkerBalance > 0.00001 {
			return true, nil
		} else {
			return false, fmt.Errorf("Not enough funds in the sequencer account and worker address, sequencer balance: %f ETH, worker address: %f ETH", cpBalance.SequencerBalance,
				cpBalance.WorkerBalance)
		}
	}

	if !conf.GetConfig().UBI.EnableSequencer && conf.GetConfig().UBI.AutoChainProof {
		if cpBalance.WorkerBalance > 0.00001 {
			return true, nil
		} else {
			return false, fmt.Errorf("No sufficient balance in the worker address, current balance: %f ETH", cpBalance.WorkerBalance)
		}
	}

	return false, fmt.Errorf("unknown error")
}

func GetCpBalance() {
	cpAccountAddress, err := contract.GetCpAccountAddress()
	if err != nil {
		logs.GetLogger().Errorf("get cp account contract address failed, error: %v", err)
		return
	}

	chainUrl, err := conf.GetRpcByNetWorkName()
	if err != nil {
		logs.GetLogger().Errorf("failed to get rpc url, cpAccount: %s, error: %v", cpAccountAddress, err)
		return

	}
	client, err := contract.GetEthClient(chainUrl)
	if err != nil {
		logs.GetLogger().Errorf("failed to dial rpc, cpAccount: %d, error: %v", cpAccountAddress, err)
		return
	}
	client.Close()

	_, workerAddress, err := GetOwnerAddressAndWorkerAddress()
	if err != nil {
		logs.GetLogger().Errorf("failed to get worker address, cpAccount: %s,error: %v", cpAccountAddress, err)
		return
	}

	workerBalance, err := wallet.BalanceNumber(client, workerAddress)
	if err != nil {
		logs.GetLogger().Errorf("failed to get worker banlance, cpAccount: %s,error: %v", cpAccountAddress, err)
		return
	}

	sequencerStub, err := ecp.NewSequencerStub(client, ecp.WithSequencerCpAccountAddress(cpAccountAddress))
	if err != nil {
		logs.GetLogger().Errorf("failed to get cp sequencer contract, cpAccount: %s, error: %v", cpAccountAddress, err)
		return
	}
	sequencerBalanceStr, err := sequencerStub.GetCPBalance()
	if err != nil {
		logs.GetLogger().Errorf("failed to get cp sequencer contract, cpAccount: %s, error: %v", cpAccountAddress, err)
		return
	}

	sequencerBalance, err := strconv.ParseFloat(sequencerBalanceStr, 64)
	if err != nil {
		logs.GetLogger().Errorf("failed to convert numbers for cp sequencer balance, cpAccount: %s, sequencerBalance: %s, error: %v", cpAccountAddress, sequencerBalanceStr, err)
		return
	}

	logs.GetLogger().Infof("cpAccount: %s, sequencer balance: %s ETH, worker address balance: %s ETH", cpAccountAddress, fmt.Sprintf("%.6f", sequencerBalance),
		fmt.Sprintf("%.6f", workerBalance))

	var cpBalance = models.CpBalanceEntity{
		CpAccount:        cpAccountAddress,
		WorkerBalance:    roundToSixDecimal(workerBalance),
		SequencerBalance: roundToSixDecimal(sequencerBalance),
	}
	cpBalanceEntity, err := NewCpBalanceService().GetCpBalance(cpAccountAddress)
	if err != nil || cpBalanceEntity == nil || cpBalanceEntity.WorkerBalance == 0 {
		err = NewCpBalanceService().SaveCpBalance(cpBalance)
	} else {
		err = NewCpBalanceService().UpdateCpBalance(cpBalance)
	}

	if err != nil {
		logs.GetLogger().Errorf("failed to update cp balance, error: %v", err)
		return
	}
}

func roundToSixDecimal(num float64) float64 {
	return math.Round(num*1e6) / 1e6
}
func RetryFn(fn func() error, maxRetries int, delay time.Duration) error {
	var err error
	for i := 0; i < maxRetries; i++ {
		err = fn()
		if err == nil {
			return nil
		}
		fmt.Printf("Attempt %d failed: %v\n", i+1, err)
		time.Sleep(delay)
	}
	return fmt.Errorf("after %d attempts, last error: %w", maxRetries, err)
}

func RestartTraefikService() error {
	dockerService := NewDockerService()
	if err := dockerService.CreateNetwork("traefik-net"); err != nil {
		return err
	}

	traefikServiceContainerName := "traefik-service"

	dockerService.RemoveContainerByName(traefikServiceContainerName)
	err := dockerService.PullImage(build.TraefikServerDockerImage)
	if err != nil {
		return fmt.Errorf("pull %s image failed, error: %v", build.UBIResourceExporterDockerImage, err)
	}

	err = dockerService.ContainerCreateAndStart(&container.Config{
		Image: build.TraefikServerDockerImage,
		Cmd: []string{
			"--api.insecure=true",
			"--log.level=INFO",
			"--providers.docker.exposedbydefault=false",
			"--providers.docker=true",
			"--entrypoints.web.address=:80",
		},
		AttachStdout: true,
		AttachStderr: true,
		Tty:          true,
	}, &container.HostConfig{
		PortBindings: map[nat.Port][]nat.PortBinding{
			"80/tcp": {{
				HostIP:   "0.0.0.0",
				HostPort: strconv.Itoa(traefikListenPortMapHost),
			}},
			"8080/tcp": {{
				HostIP:   "0.0.0.0",
				HostPort: "9080",
			}},
		},
		Binds: []string{"/var/run/docker.sock:/var/run/docker.sock"},
		RestartPolicy: container.RestartPolicy{
			Name:              container.RestartPolicyOnFailure,
			MaximumRetryCount: 3,
		},
		Privileged: true,
	}, &network.NetworkingConfig{
		EndpointsConfig: map[string]*network.EndpointSettings{
			"traefik-net": {},
		},
	}, traefikServiceContainerName)
	if err != nil {
		return fmt.Errorf("create traefik-service container failed, error: %v", err)
	}
	return nil
}

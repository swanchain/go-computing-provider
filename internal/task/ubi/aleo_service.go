package ubi

import (
	"context"
	"fmt"
	"github.com/filswan/go-swan-lib/logs"
	"github.com/joho/godotenv"
	"github.com/swanchain/go-computing-provider/build"
	"github.com/swanchain/go-computing-provider/conf"
	"github.com/swanchain/go-computing-provider/constants"
	"github.com/swanchain/go-computing-provider/internal/computing"
	"github.com/swanchain/go-computing-provider/internal/models"
	batchv1 "k8s.io/api/batch/v1"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const AleoConfigFile = "aleo_proof.env"

type AleoTask struct {
}

func NewAleoTask() *AleoTask {
	return &AleoTask{}
}

func (aleo *AleoTask) TaskType() string {
	return TASK_ZK_TYPE_ALEO
}

func (aleo *AleoTask) DoTask(ubiTask models.UBITaskReq, resourceRequirements coreV1.ResourceRequirements, imageName, namespace, nodeName string, useGpu bool) {
	var envFilePath string
	envFilePath = filepath.Join(os.Getenv("CP_PATH"), AleoConfigFile)
	envVars, err := godotenv.Read(envFilePath)
	if err != nil {
		logs.GetLogger().Errorf("reading %s failed, error: %v", AleoConfigFile, err)
		return
	}

	var taskType = "CPU"
	if useGpu {
		delete(envVars, "RUST_GPU_TOOLS_CUSTOM_GPU")
		envVars["BELLMAN_NO_GPU"] = "1"
		taskType = "GPU"
	}

	go func() {
		var err error
		defer func() {
			key := constants.REDIS_UBI_ALEO_PERFIX + strconv.Itoa(ubiTask.ID)
			ubiTaskRun, _ := computing.RetrieveUbiTaskMetadata(key)
			if ubiTaskRun.TaskId == "" {
				ubiTaskRun = new(models.CacheUbiTaskDetail)
				ubiTaskRun.TaskId = strconv.Itoa(ubiTask.ID)
				ubiTaskRun.TaskType = taskType
				ubiTaskRun.ZkType = ubiTask.ZkType
				ubiTaskRun.CreateTime = time.Now().Format("2006-01-02 15:04:05")
			}

			if err == nil {
				ubiTaskRun.Status = constants.UBI_TASK_RUNNING_STATUS
			} else {
				ubiTaskRun.Status = constants.UBI_TASK_FAILED_STATUS
				k8sService := computing.NewK8sService()
				k8sService.Client.CoreV1().Namespaces().Delete(context.TODO(), namespace, metaV1.DeleteOptions{})
			}
			computing.SaveUbiTaskMetadata(ubiTaskRun)
		}()

		domain := conf.GetConfig().API.Domain
		if strings.HasPrefix(domain, ".") {
			domain = domain[1:]
		}
		domain = "ubi-aleo." + domain
		localIp, err := getLocalIp()
		if err != nil {
			logs.GetLogger().Errorf("check resource failed, error: %v", err)
			return
		}

		receiveUrl := fmt.Sprintf("https://%s:%d/api/v1/computing/cp/receive/ubi", domain, conf.GetConfig().API.Port)
		execCommand := []string{"/mnt/init/cmd", "start", "--prover"}
		JobName := strings.ToLower(ubiTask.ZkType) + "-" + strconv.Itoa(ubiTask.ID)

		var useEnvVars []coreV1.EnvVar
		for k, v := range envVars {
			useEnvVars = append(useEnvVars, coreV1.EnvVar{
				Name:  k,
				Value: v,
			})
		}

		useEnvVars = append(useEnvVars,
			coreV1.EnvVar{
				Name:  "RECEIVE_PROOF_URL",
				Value: receiveUrl,
			},
			coreV1.EnvVar{
				Name:  "TASKID",
				Value: strconv.Itoa(ubiTask.ID),
			},
			coreV1.EnvVar{
				Name:  "TASK_TYPE",
				Value: strconv.Itoa(ubiTask.Type),
			},
			coreV1.EnvVar{
				Name:  "ZK_TYPE",
				Value: ubiTask.ZkType,
			},
			coreV1.EnvVar{
				Name:  "NAME_SPACE",
				Value: namespace,
			},
			coreV1.EnvVar{
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
				Template: coreV1.PodTemplateSpec{
					Spec: coreV1.PodSpec{
						NodeName: nodeName,
						Containers: []coreV1.Container{
							{
								Name:            JobName + generateString(5),
								Image:           imageName,
								Env:             useEnvVars,
								Command:         execCommand,
								Resources:       resourceRequirements,
								ImagePullPolicy: coreV1.PullIfNotPresent,
							},
						},
						RestartPolicy: "Never",
						HostAliases: []coreV1.HostAlias{
							{
								IP:        localIp,
								Hostnames: []string{domain},
							},
						},
					},
				},
				BackoffLimit:            new(int32),
				TTLSecondsAfterFinished: new(int32),
			},
		}

		*job.Spec.BackoffLimit = 1
		*job.Spec.TTLSecondsAfterFinished = 120

		k8sService := computing.NewK8sService()
		if _, err = k8sService.Client.BatchV1().Jobs(namespace).Create(context.TODO(), job, metaV1.CreateOptions{}); err != nil {
			logs.GetLogger().Errorf("Failed creating FIL-C2 ubi task job: %v", err)
			return
		}
	}()
}

func (aleo *AleoTask) ReceiveUbiProof(c2Proof models.ReceiveProof) {
	var submitUBIProofTx string
	var err error
	defer func() {
		key := constants.REDIS_UBI_ALEO_PERFIX + c2Proof.TaskId
		ubiTask, _ := computing.RetrieveUbiTaskMetadata(key)
		if err == nil {
			ubiTask.Status = constants.UBI_TASK_SUCCESS_STATUS
		} else {
			ubiTask.Status = constants.UBI_TASK_FAILED_STATUS
		}
		ubiTask.Tx = submitUBIProofTx
		computing.SaveUbiTaskMetadata(ubiTask)
		if strings.TrimSpace(c2Proof.NameSpace) != "" {
			k8sService := computing.NewK8sService()
			k8sService.Client.CoreV1().Namespaces().Delete(context.TODO(), c2Proof.NameSpace, metaV1.DeleteOptions{})
		}
	}()
	retries := 3
	for i := 0; i < retries; i++ {
		submitUBIProofTx, err = submitUBIProof(c2Proof)
		if err == nil {
			break
		}
		time.Sleep(time.Second * 2)
	}
}

func (aleo *AleoTask) GetConfGpuName() string {
	var envFilePath string
	envFilePath = filepath.Join(os.Getenv("CP_PATH"), AleoConfigFile)
	envVars, err := godotenv.Read(envFilePath)
	if err != nil {
		logs.GetLogger().Errorf("reading %s failed, error: %v", AleoConfigFile, err)
		return ""
	}

	c2GpuConfig := envVars["RUST_GPU_TOOLS_CUSTOM_GPU"]
	return convertGpuName(strings.TrimSpace(c2GpuConfig))
}

// GetImageName architecture: AMD or INTEL
func (aleo *AleoTask) GetImageName(architecture string, useGpu bool) string {
	var ubiTaskImage string
	if architecture == constants.CPU_AMD {
		ubiTaskImage = build.UBITaskAleoProofImageAmdGpu
	} else if architecture == constants.CPU_INTEL {
		ubiTaskImage = build.UBITaskAleoProofImageAmdGpu
	}
	ubiTaskImage = build.UBITaskAleoProofImageAmdGpu
	return ubiTaskImage
}

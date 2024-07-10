package computing

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/filswan/go-mcs-sdk/mcs/api/common/logs"
	"github.com/robfig/cron/v3"
	"github.com/swanchain/go-computing-provider/conf"
	"github.com/swanchain/go-computing-provider/constants"
	"github.com/swanchain/go-computing-provider/internal/contract/fcp"
	"github.com/swanchain/go-computing-provider/internal/models"
	"io"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

var deployingChan = make(chan models.Job)
var TaskMap sync.Map

type CronTask struct {
	nodeId       string
	ownerAddress string
	location     string
}

func NewCronTask(nodeId string) *CronTask {
	ownerAddress, _, err := GetOwnerAddressAndWorkerAddress()
	if err != nil {
		logs.GetLogger().Errorf("get owner address failed, error: %v", err)
		return nil
	}

	location, err := getLocation()
	if err != nil {
		logs.GetLogger().Error(err)
	}

	return &CronTask{nodeId: nodeId, ownerAddress: ownerAddress, location: location}
}

func (task *CronTask) RunTask() {
	addNodeLabel()
	checkJobStatus()
	task.checkCollateralBalance()
	task.cleanAbnormalDeployment()
	task.setFailedUbiTaskStatus()
	task.watchNameSpaceForDeleted()
	task.updateUbiTaskReward()
	task.reportClusterResourceToHub()
	task.watchExpiredTask()
}

func checkJobStatus() {
	go func() {
		for {
			select {
			case job := <-deployingChan:
				TaskMap.Store(job.Uuid, &job)
			case <-time.After(3 * time.Second):
				TaskMap.Range(func(key, value any) bool {
					jobUuid := key.(string)
					job := value.(*models.Job)
					if reportJobStatus(jobUuid, job.Status) && job.Status == models.DEPLOY_TO_K8S {
						TaskMap.Delete(jobUuid)
					}
					return true
				})
			}
		}
	}()
}

func (task *CronTask) reportClusterResourceToHub() {
	c := cron.New(cron.WithSeconds())
	c.AddFunc("0/10 * * * * ?", func() {
		defer func() {
			if err := recover(); err != nil {
				logs.GetLogger().Errorf("Failed report cp resource's summary, error: %+v", err)
			}
		}()

		k8sService := NewK8sService()
		statisticalSources, err := k8sService.StatisticalSources(context.TODO())
		if err != nil {
			logs.GetLogger().Errorf("Failed k8s statistical sources, error: %+v", err)
			return
		}
		checkClusterProviderStatus(statisticalSources)
	})
	c.Start()
}

func (task *CronTask) watchNameSpaceForDeleted() {
	c := cron.New(cron.WithSeconds())
	c.AddFunc("0 0/50 * * * ?", func() {
		defer func() {
			if err := recover(); err != nil {
				logs.GetLogger().Errorf("watchNameSpaceForDeleted catch panic error: %+v", err)
			}
		}()
		service := NewK8sService()
		namespaces, err := service.ListNamespace(context.TODO())
		if err != nil {
			logs.GetLogger().Errorf("Failed get all namespace, error: %+v", err)
			return
		}

		for _, namespace := range namespaces {
			getPods, err := service.GetPods(namespace, "")
			if err != nil {
				logs.GetLogger().Errorf("Failed get pods form namespace,namepace: %s, error: %+v", namespace, err)
				continue
			}
			if !getPods && (strings.HasPrefix(namespace, constants.K8S_NAMESPACE_NAME_PREFIX) || strings.HasPrefix(namespace, "ubi-task")) {
				if err = service.DeleteNameSpace(context.TODO(), namespace); err != nil {
					logs.GetLogger().Errorf("Failed delete namespace, namepace: %s, error: %+v", namespace, err)
				}
			}
		}
		NewDockerService().CleanResource()
	})
	c.Start()
}

func (task *CronTask) watchExpiredTask() {
	c := cron.New(cron.WithSeconds())
	c.AddFunc("0 0/5 * * * ?", func() {
		defer func() {
			if err := recover(); err != nil {
				logs.GetLogger().Errorf("watchExpiredTask catch panic error: %+v", err)
			}
		}()

		jobList, err := NewJobService().GetJobList()
		if err != nil {
			logs.GetLogger().Errorf("Failed watchExpiredTask get job data, error: %+v", err)
			return
		}

		deployments, err := NewK8sService().k8sClient.AppsV1().Deployments(metav1.NamespaceAll).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			fmt.Println("Error listing deployments:", err)
			return
		}

		var deployOnK8s = make(map[string]string)
		for _, deploy := range deployments.Items {
			if strings.HasPrefix(deploy.Namespace, constants.K8S_NAMESPACE_NAME_PREFIX) {
				deployOnK8s[deploy.Name] = deploy.Namespace
			}
		}

		if len(jobList) == 0 {
			service := NewK8sService()
			namespaces, err := service.ListNamespace(context.TODO())
			if err != nil {
				logs.GetLogger().Errorf("Failed get all namespace, error: %+v", err)
				return
			}

			for _, namespace := range namespaces {
				if strings.HasPrefix(namespace, constants.K8S_NAMESPACE_NAME_PREFIX) {
					service.DeleteNameSpace(context.TODO(), namespace)
				}
			}
			return
		}

		var deleteSpaceIds []string
		for _, job := range jobList {
			if _, ok := deployOnK8s[job.K8sDeployName]; ok {
				delete(deployOnK8s, job.K8sDeployName)
			}

			if _, err = NewK8sService().k8sClient.AppsV1().Deployments(job.NameSpace).Get(context.TODO(), job.K8sDeployName, metav1.GetOptions{}); err != nil && errors.IsNotFound(err) {
				deleteSpaceIds = append(deleteSpaceIds, job.SpaceUuid)
				continue
			}

			if len(strings.TrimSpace(job.TaskUuid)) != 0 {
				taskStatus, err := checkTaskStatusByHub(job.TaskUuid, task.nodeId)
				if err != nil {
					logs.GetLogger().Errorf("Failed check task status by Orchestrator service, error: %+v", err)
					return
				}
				if strings.Contains(taskStatus, "no task found") {
					logs.GetLogger().Infof("task_uuid: %s, task not found on the orchestrator service, starting to delete it.", job.TaskUuid)
					deleteJob(job.NameSpace, job.SpaceUuid, "cron task, no task found")
					deleteSpaceIds = append(deleteSpaceIds, job.SpaceUuid)
					continue
				}
				if strings.Contains(taskStatus, "Terminated") || strings.Contains(taskStatus, "Terminated") ||
					strings.Contains(taskStatus, "Cancelled") || strings.Contains(taskStatus, "Failed") {
					logs.GetLogger().Infof("task_uuid: %s, current status is %s, starting to delete it.", job.TaskUuid, taskStatus)
					if err = deleteJob(job.NameSpace, job.SpaceUuid, "cron task, abnormal state"); err == nil {
						deleteSpaceIds = append(deleteSpaceIds, job.SpaceUuid)
						continue
					}
				}
			}

			if time.Now().Unix() > job.ExpireTime {
				logs.GetLogger().Infof("<timer-task> space_uuid: %s has expired, the job starting terminated", job.SpaceUuid)
				if err = deleteJob(job.NameSpace, job.SpaceUuid, "cron task, run time expired"); err == nil {
					deleteSpaceIds = append(deleteSpaceIds, job.SpaceUuid)
					continue
				}
			}
		}

		for deploymentName, nameSpace := range deployOnK8s {
			if nameSpace != "" && deploymentName != "" {
				spaceUuid := strings.TrimPrefix(deploymentName, constants.K8S_DEPLOY_NAME_PREFIX)
				deleteJob(nameSpace, spaceUuid, "cron task, task left on k8s")
			}
		}
		for _, spaceUuid := range deleteSpaceIds {
			NewJobService().DeleteJobEntityBySpaceUuId(spaceUuid)
		}

	})
	c.Start()
}

func (task *CronTask) checkCollateralBalance() {
	c := cron.New(cron.WithSeconds())
	c.AddFunc("0 0/10 * * * * ?", func() {
		defer func() {
			if err := recover(); err != nil {
				logs.GetLogger().Errorf("task job: [checkCollateralBalance], error: %+v", err)
			}
		}()

		result, err := checkFcpCollateralBalance()
		if err != nil {
			logs.GetLogger().Errorf("check collateral balance failed, error: %+v", err)
			return
		}

		floatResult, err := strconv.ParseFloat(result, 64)
		if err != nil {
			logs.GetLogger().Errorf("parse collateral balance failed, error: %+v", err)
			return
		}

		if floatResult <= conf.GetConfig().HUB.BalanceThreshold {
			logs.GetLogger().Warnf("No sufficient collateral Balance, the current collateral balance is: %0.3f. Please run: computing-provider collateral [fromWalletAddress] [amount]", floatResult)
		}
	})
	c.Start()
}

func (task *CronTask) cleanAbnormalDeployment() {
	c := cron.New(cron.WithSeconds())
	c.AddFunc("* 0/30 * * * ?", func() {
		defer func() {
			if err := recover(); err != nil {
				logs.GetLogger().Errorf("task job: [cleanAbnormalDeployment], error: %+v", err)
			}
		}()

		k8sService := NewK8sService()
		namespaces, err := k8sService.ListNamespace(context.TODO())
		if err != nil {
			logs.GetLogger().Errorf("Failed get all namespace, error: %+v", err)
			return
		}

		for _, namespace := range namespaces {
			if strings.HasPrefix(namespace, constants.K8S_NAMESPACE_NAME_PREFIX) {
				deployments, err := k8sService.k8sClient.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
				if err != nil {
					logs.GetLogger().Errorf("Error getting deployments in namespace %s: %v\n", namespace, err)
					continue
				}

				for _, deployment := range deployments.Items {
					creationTimestamp := deployment.ObjectMeta.CreationTimestamp.Time
					currentTime := time.Now()
					age := currentTime.Sub(creationTimestamp)
					if deployment.Status.AvailableReplicas == 0 && age.Hours() >= 2 {
						logs.GetLogger().Infof("Cleaning up deployment %s in namespace %s", deployment.Name, namespace)
						err := k8sService.k8sClient.AppsV1().Deployments(namespace).Delete(context.TODO(), deployment.Name, metav1.DeleteOptions{})
						if err != nil {
							if errors.IsNotFound(err) {
								logs.GetLogger().Errorf("Deployment %s not found. Ignoring", deployment.Name)
							} else {
								logs.GetLogger().Errorf("Error deleting deployment %s: %v", deployment.Name, err)
							}
						} else {
							logs.GetLogger().Errorf("abnormal Deployment %s deleted successfully.", deployment.Name)
						}
					}
				}
			}
		}
	})
	c.Start()
}

func (task *CronTask) setFailedUbiTaskStatus() {
	c := cron.New(cron.WithSeconds())
	c.AddFunc("0 0/8 * * * ?", func() {
		defer func() {
			if err := recover(); err != nil {
				logs.GetLogger().Errorf("task job: [setFailedUbiTaskStatus], error: %+v", err)
			}
		}()

		var taskList []models.TaskEntity
		oneHourAgo := time.Now().Add(-1 * time.Hour).Unix()
		err := NewTaskService().Model(&models.TaskEntity{}).Where("status in (?,?)", models.TASK_RECEIVED_STATUS, models.TASK_RUNNING_STATUS).
			Or("status ==? and tx_hash !=''", models.TASK_FAILED_STATUS).
			Or("status=? and tx_hash==''", models.TASK_SUCCESS_STATUS).Find(&taskList).Error
		if err != nil {
			logs.GetLogger().Errorf("Failed get task list, error: %+v", err)
			return
		}

		for _, entity := range taskList {
			ubiTask := entity

			if ubiTask.CreateTime < oneHourAgo {
				JobName := strings.ToLower(models.UbiTaskTypeStr(ubiTask.Type)) + "-" + strconv.Itoa(int(ubiTask.Id))
				k8sNameSpace := "ubi-task-" + strconv.Itoa(int(ubiTask.Id))
				NewK8sService().k8sClient.BatchV1().Jobs(k8sNameSpace).Delete(context.TODO(), JobName, metav1.DeleteOptions{})
				ubiTask.Status = models.TASK_FAILED_STATUS
			}

			if ubiTask.TxHash != "" {
				ubiTask.Status = models.TASK_SUCCESS_STATUS
			} else {
				ubiTask.Status = models.TASK_FAILED_STATUS
			}

			NewTaskService().SaveTaskEntity(&ubiTask)
		}
	})
	c.Start()
}

func (task *CronTask) updateUbiTaskReward() {
	c := cron.New(cron.WithSeconds())
	c.AddFunc("0 0/30 * * * ?", func() {
		defer func() {
			if err := recover(); err != nil {
				logs.GetLogger().Errorf("task job: [updateUbiTaskReward], error: %+v", err)
			}
		}()

		taskList, err := NewTaskService().GetTaskListNoReward()
		if err != nil {
			logs.GetLogger().Errorf("get task list failed, error: %+v", err)
			return
		}

		for _, entity := range taskList {
			ubiTask := entity
			err = getReward(ubiTask)
			if err != nil {
				logs.GetLogger().Errorf("taskId: %d, %v", ubiTask.Id, err)
				continue
			}
		}
	})
	c.Start()
}

func addNodeLabel() {
	k8sService := NewK8sService()
	nodes, err := k8sService.k8sClient.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	nodeGpuInfoMap, err := k8sService.GetResourceExporterPodLog(context.TODO())
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	logs.GetLogger().Infof("collect all node: %d", len(nodes.Items))
	for _, node := range nodes.Items {
		cpNode := node
		if collectInfo, ok := nodeGpuInfoMap[cpNode.Name]; ok {
			for _, detail := range collectInfo.Gpu.Details {
				if err = k8sService.AddNodeLabel(cpNode.Name, detail.ProductName); err != nil {
					logs.GetLogger().Errorf("add node label, nodeName %s, gpuName: %s, error: %+v", cpNode.Name, detail.ProductName, err)
					continue
				}
			}
			err := k8sService.AddNodeLabel(cpNode.Name, collectInfo.CpuName)
			if err != nil {
				logs.GetLogger().Errorf("nodeName: %s, error: %v", cpNode.Name, err)
			}
		}
	}
}

func checkTaskStatusByHub(taskUuid, nodeId string) (string, error) {
	url := fmt.Sprintf("%s/check_task_status_with_node_id/%s/%s", conf.GetConfig().HUB.ServerUrl, taskUuid, nodeId)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Add("offset", "0")
	req.Header.Add("limit", "10")
	req.Header.Add("Authorization", "Bearer "+conf.GetConfig().HUB.AccessToken)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making HTTP request:", err)
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var taskStatus struct {
		Data struct {
			JobStatus  string `json:"job_status"`
			TaskStatus string `json:"task_status"`
		} `json:"data"`
		Message string `json:"message"`
		Status  string `json:"status"`
	}
	err = json.Unmarshal(respBody, &taskStatus)
	if err != nil {
		logs.GetLogger().Errorf("check_task_status_with_node_id resp: %s", string(respBody))
		return "", err
	}
	if taskStatus.Status == "failed" {
		return taskStatus.Message, nil
	}
	return taskStatus.Status, nil
}

func reportJobStatus(jobUuid string, deployStatus int) bool {
	var job = new(models.JobEntity)
	job.JobUuid = jobUuid
	job.DeployStatus = deployStatus
	if err := NewJobService().UpdateJobEntityByJobUuid(job); err != nil {
		logs.GetLogger().Errorf("update job info by jobUuid failed, error: %v", err)
	}
	return true
}

func checkFcpCollateralBalance() (string, error) {

	chainRpc, err := conf.GetRpcByNetWorkName()
	if err != nil {
		return "", err
	}
	client, err := ethclient.Dial(chainRpc)
	if err != nil {
		return "", err
	}
	defer client.Close()

	fcpCollateralStub, err := fcp.NewCollateralStub(client)
	if err != nil {
		return "", err
	}

	fcpCollateralInfo, err := fcpCollateralStub.CollateralInfo()
	if err != nil {
		return "", err
	}
	return fcpCollateralInfo.AvailableBalance, nil
}

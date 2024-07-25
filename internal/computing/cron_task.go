package computing

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/filswan/go-mcs-sdk/mcs/api/common/logs"
	"github.com/robfig/cron/v3"
	"github.com/swanchain/go-computing-provider/conf"
	"github.com/swanchain/go-computing-provider/constants"
	"github.com/swanchain/go-computing-provider/internal/contract/fcp"
	"github.com/swanchain/go-computing-provider/internal/models"
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
	nodeId string
}

func NewCronTask(nodeId string) *CronTask {
	return &CronTask{nodeId: nodeId}
}

func (task *CronTask) RunTask() {
	addNodeLabel()
	checkJobStatus()
	task.checkCollateralBalance()
	task.cleanAbnormalDeployment()
	task.setFailedUbiTaskStatus()
	task.watchNameSpaceForDeleted()
	task.reportClusterResource()
	task.watchExpiredTask()
	task.GetUbiTaskReward()
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

func (task *CronTask) reportClusterResource() {
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
			logs.GetLogger().Errorf("failed to collect k8s statistical sources, error: %+v", err)
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
	c.AddFunc("0 0/10 * * * ?", func() {
		defer func() {
			if err := recover(); err != nil {
				logs.GetLogger().Errorf("watchExpiredTask catch panic error: %+v", err)
			}
		}()

		jobList, err := NewJobService().GetJobList(models.UN_DELETEED_FLAG)
		if err != nil {
			logs.GetLogger().Errorf("failed to get job data, error: %+v", err)
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

		if len(deployOnK8s) == 0 && len(jobList) > 0 {
			for _, job := range jobList {
				if err = NewJobService().DeleteJobEntityBySpaceUuId(job.SpaceUuid, models.JOB_COMPLETED_STATUS); err != nil {
					logs.GetLogger().Infof("failed to delete job from db, space_uuid: %s, error: %v", job.SpaceUuid, err)
				}
			}
		} else if len(deployOnK8s) > 0 && len(jobList) == 0 {
			for _, namespace := range deployOnK8s {
				if err = NewK8sService().DeleteNameSpace(context.TODO(), namespace); err != nil {
					logs.GetLogger().Errorf("failed to delete job from k8s, namespace: %s, error: %v", namespace, err)
				}
			}
		} else if len(deployOnK8s) > 0 && len(jobList) > 0 {
			var deleteSpaceIds []string
			for _, job := range jobList {
				if _, ok := deployOnK8s[job.K8sDeployName]; ok {
					delete(deployOnK8s, job.K8sDeployName)
				}

				if time.Now().Sub(time.Unix(job.CreateTime, 0)).Hours() > 2 {
					deleteSpaceIds = append(deleteSpaceIds, job.SpaceUuid)
					continue
				}

				checkFcpJobInfoInChain(job)

				if job.Status == models.JOB_TERMINATED_STATUS || job.Status == models.JOB_COMPLETED_STATUS {
					logs.GetLogger().Infof("task_uuid: %s, current status is %s, starting to delete it.", job.TaskUuid, models.GetJobStatus(job.Status))
					if err = deleteJob(job.NameSpace, job.SpaceUuid, "cron task, abnormal state"); err == nil {
						deleteSpaceIds = append(deleteSpaceIds, job.SpaceUuid)
						continue
					}
				} else if job.Status == models.JOB_RUNNING_STATUS {
					if !checkHealth(job.RealUrl) {
						if err = deleteJob(job.NameSpace, job.SpaceUuid, "cron task, job service is not available"); err == nil {
							deleteSpaceIds = append(deleteSpaceIds, job.SpaceUuid)
							continue
						}
					}
				}

				if time.Now().Unix() > job.ExpireTime {
					if err = deleteJob(job.NameSpace, job.SpaceUuid, "cron-task the task execution time has expired"); err == nil {
						deleteSpaceIds = append(deleteSpaceIds, job.SpaceUuid)
						continue
					}
				}
			}

			for deploymentName, nameSpace := range deployOnK8s {
				if nameSpace != "" && deploymentName != "" {
					spaceUuid := strings.TrimPrefix(deploymentName, constants.K8S_DEPLOY_NAME_PREFIX)
					deleteJob(nameSpace, spaceUuid, "cron-task the obsolete task left in the k8s")
				}
			}
			for _, spaceUuid := range deleteSpaceIds {
				NewJobService().DeleteJobEntityBySpaceUuId(spaceUuid, models.JOB_COMPLETED_STATUS)
			}
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
			Or("status=? and tx_hash==''", models.TASK_SUBMITTED_STATUS).Find(&taskList).Error
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
				ubiTask.Status = models.TASK_SUBMITTED_STATUS
			} else {
				ubiTask.Status = models.TASK_FAILED_STATUS
			}

			NewTaskService().SaveTaskEntity(&ubiTask)
		}
	})
	c.Start()
}

func (task *CronTask) CheckJobReward() {
	c := cron.New(cron.WithSeconds())
	c.AddFunc("* 0/30 * * * ?", func() {
		defer func() {
			if err := recover(); err != nil {
				logs.GetLogger().Errorf("task job: [cleanAbnormalDeployment], error: %+v", err)
			}
		}()

		//jobList, err := NewJobService().GetJobListByNoReward()
		//if err != nil {
		//	logs.GetLogger().Errorf("failed to get job data, error: %+v", err)
		//	return
		//}
		//for _, job := range jobList {
		//
		//}

	})
	c.Start()
}

func (task *CronTask) GetUbiTaskReward() {
	c := cron.New(cron.WithSeconds())
	c.AddFunc("* 0/10 * * ?", func() {
		defer func() {
			if err := recover(); err != nil {
				logs.GetLogger().Errorf("task job: [GetUbiTaskReward], error: %+v", err)
			}
		}()

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

func reportJobStatus(jobUuid string, deployStatus int) bool {
	var job = new(models.JobEntity)
	job.JobUuid = jobUuid
	job.DeployStatus = deployStatus
	if deployStatus < models.DEPLOY_TO_K8S {
		job.PodStatus = models.POD_UNKNOWN_STATUS
		job.Status = models.JOB_DEPLOY_STATUS
	} else {
		job.PodStatus = models.POD_RUNNING_STATUS
		job.Status = models.JOB_RUNNING_STATUS
	}

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

func checkFcpJobInfoInChain(job *models.JobEntity) {
	var taskInfo models.TaskInfoOnChain
	var err error
	chainRpc, err := conf.GetRpcByNetWorkName()
	if err != nil {
		return
	}
	client, err := ethclient.Dial(chainRpc)
	if err != nil {
		return
	}
	defer client.Close()
	taskManagerStub, err := fcp.NewTaskManagerStub(client)
	if err != nil {
		return
	}
	taskInfo, err = taskManagerStub.GetTaskInfo(job.TaskUuid)
	if err != nil {
		return
	}

	if taskInfo.TaskStatus == models.COMPLETED {
		job.Status = models.JOB_COMPLETED_STATUS
	} else if taskInfo.TaskStatus == models.TERMINATED {
		job.Status = models.JOB_TERMINATED_STATUS
	}
}

func checkHealth(url string) bool {
	response, err := http.Get(url)
	if err != nil {
		return false
	}
	defer response.Body.Close()
	return response.StatusCode == 200
}

type TaskGroup struct {
	Items []*models.TaskEntity
	Ids   []int64
}

func handleTasksToGroup(list []*models.TaskEntity) []TaskGroup {
	var groups []TaskGroup
	const batchSize = 20
	for i := 0; i < len(list); i += batchSize {
		end := i + batchSize
		if end > len(list) {
			end = len(list)
		}
		batch := list[i:end]

		var ids []int64
		for _, item := range batch {
			ids = append(ids, item.Id)
		}

		group := TaskGroup{
			Items: batch,
			Ids:   ids,
		}
		groups = append(groups, group)
	}
	return groups
}

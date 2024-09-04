package computing

import (
	"context"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/filswan/go-mcs-sdk/mcs/api/common/logs"
	"github.com/robfig/cron/v3"
	"github.com/swanchain/go-computing-provider/conf"
	"github.com/swanchain/go-computing-provider/constants"
	"github.com/swanchain/go-computing-provider/internal/contract/fcp"
	"github.com/swanchain/go-computing-provider/internal/models"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var NetworkPolicyFlag bool

var deployingChan = make(chan models.Job)
var TaskMap sync.Map

type CronTask struct {
	nodeId string
}

func NewCronTask(nodeId string) *CronTask {
	return &CronTask{nodeId: nodeId}
}

func (task *CronTask) RunTask() {
	checkClusterNetworkPolicy()
	addNodeLabel()
	checkJobStatus()
	task.checkCollateralBalance()
	task.cleanAbnormalDeployment()
	task.setFailedUbiTaskStatus()
	task.watchNameSpaceForDeleted()
	task.reportClusterResource()
	task.watchExpiredTask()
	task.getUbiTaskReward()
	task.checkJobReward()
	task.cleanImageResource()
}

func checkClusterNetworkPolicy() {
	var err error
	defer func() {
		if err != nil {
			NewK8sService().DeleteGlobalNetworkSet(models.NetworkNetset)
			NewK8sService().DeleteGlobalNetworkPolicy(models.NetworkGlobalSubnet)
			NewK8sService().DeleteGlobalNetworkPolicy(models.NetworkGlobalOutAccess)
			NewK8sService().DeleteGlobalNetworkPolicy(models.NetworkGlobalInAccess)
			NewK8sService().DeleteGlobalNetworkPolicy(models.NetworkGlobalNamespace)
			NewK8sService().DeleteGlobalNetworkPolicy(models.NetworkGlobalDns)
		}
	}()

	netset, err := NewK8sService().GetGlobalNetworkSet(models.NetworkNetset)
	if err != nil {
		if errors.IsNotFound(err) {
			generateNewNetworkPolicy()
		} else {
			logs.GetLogger().Error(err)
		}
		return
	}

	if netset != nil {
		_, err = NewK8sService().GetGlobalNetworkPolicy(models.NetworkGlobalSubnet)
		if err != nil {
			if errors.IsNotFound(err) {
				err = NewK8sService().GenerateGlobalNetworkPoliciesForSubnet()
				if err != nil {
					logs.GetLogger().Error(err)
					return
				}
			}
		}

		_, err = NewK8sService().GetGlobalNetworkPolicy(models.NetworkGlobalOutAccess)
		if err != nil {
			if errors.IsNotFound(err) {
				err = NewK8sService().GenerateGlobalNetworkPoliciesForOutAccess()
				if err != nil {
					logs.GetLogger().Error(err)
					return
				}
			}
		}

		_, err = NewK8sService().GetGlobalNetworkPolicy(models.NetworkGlobalInAccess)
		if err != nil {
			if errors.IsNotFound(err) {
				err = NewK8sService().GenerateGlobalNetworkForInAccess()
				if err != nil {
					logs.GetLogger().Error(err)
					return
				}
			}
		}

		_, err = NewK8sService().GetGlobalNetworkPolicy(models.NetworkGlobalNamespace)
		if err != nil {
			if errors.IsNotFound(err) {
				err = NewK8sService().GenerateGlobalNetworkPoliciesForNamespace()
				if err != nil {
					logs.GetLogger().Error(err)
					return
				}
			}
		}

		_, err = NewK8sService().GetGlobalNetworkPolicy(models.NetworkGlobalDns)
		if err != nil {
			if errors.IsNotFound(err) {
				err = NewK8sService().GenerateGlobalNetworkPoliciesForDNS()
				if err != nil {
					logs.GetLogger().Error(err)
					return
				}
			}
		}
	}

	NetworkPolicyFlag = true
}

func generateNewNetworkPolicy() {
	var err error
	defer func() {
		if err != nil {
			NewK8sService().DeleteGlobalNetworkSet(models.NetworkNetset)
			NewK8sService().DeleteGlobalNetworkPolicy(models.NetworkGlobalSubnet)
			NewK8sService().DeleteGlobalNetworkPolicy(models.NetworkGlobalOutAccess)
			NewK8sService().DeleteGlobalNetworkPolicy(models.NetworkGlobalInAccess)
			NewK8sService().DeleteGlobalNetworkPolicy(models.NetworkGlobalNamespace)
			NewK8sService().DeleteGlobalNetworkPolicy(models.NetworkGlobalDns)
		}
	}()

	if err = NewK8sService().GenerateGlobalNetworkSet(); err != nil {
		logs.GetLogger().Error(err)
		return
	}

	err = NewK8sService().GenerateGlobalNetworkPoliciesForSubnet()
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	err = NewK8sService().GenerateGlobalNetworkPoliciesForOutAccess()
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	err = NewK8sService().GenerateGlobalNetworkForInAccess()
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	err = NewK8sService().GenerateGlobalNetworkPoliciesForNamespace()
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}
	err = NewK8sService().GenerateGlobalNetworkPoliciesForDNS()
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}
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
	})
	c.Start()
}

func (task *CronTask) cleanImageResource() {
	c := cron.New(cron.WithSeconds())
	c.AddFunc("* 0/30 * * * ?", func() {
		defer func() {
			if err := recover(); err != nil {
				logs.GetLogger().Errorf("cleanImageResource catch panic error: %+v", err)
			}
		}()
		NewDockerService().CleanResource()
	})
	c.Start()
}

func (task *CronTask) watchExpiredTask() {
	c := cron.New(cron.WithSeconds())
	c.AddFunc("* 0/10 * * * ?", func() {
		defer func() {
			if err := recover(); err != nil {
				logs.GetLogger().Errorf("watchExpiredTask catch panic error: %+v", err)
			}
		}()

		jobList, err := NewJobService().GetJobList(models.All_FLAG)
		if err != nil {
			logs.GetLogger().Errorf("failed to get job data, error: %+v", err)
			return
		}

		var deleteSpaceIds []string
		var deleteJobIds []string

		for _, job := range jobList {
			if job.DeleteAt == models.DELETED_FLAG && job.DeployStatus == models.DEPLOY_TO_K8S {
				continue
			}

			currentTime := time.Now()
			createdTime := time.Unix(job.CreateTime, 0)
			createDuration := currentTime.Sub(createdTime)
			if createDuration.Hours() <= 2 {
				continue
			}

			if job.NameSpace != "" && job.K8sDeployName != "" {
				_, err = NewK8sService().k8sClient.AppsV1().Deployments(job.NameSpace).Get(context.TODO(), job.K8sDeployName, metav1.GetOptions{})
				if err != nil {
					if errors.IsNotFound(err) {
						// delete job
						logs.GetLogger().Warnf("not found deployment on the cluster, space_uuid: %s, deployment: %s", job.SpaceUuid, job.K8sDeployName)
						deleteSpaceIds = append(deleteSpaceIds, job.SpaceUuid+"_"+job.JobUuid)
						deleteJobIds = append(deleteJobIds, job.JobUuid)
						continue
					}
					logs.GetLogger().Errorf("failed to get deployment: %s, error: %v", job.K8sDeployName, err)
					continue
				}
			}

			checkFcpJobInfoInChain(job)

			if job.Status == models.JOB_TERMINATED_STATUS || job.Status == models.JOB_COMPLETED_STATUS {
				logs.GetLogger().Infof("task_uuid: %s, current status is %s, starting to delete it.", job.TaskUuid, models.GetJobStatus(job.Status))
				if err = DeleteJob(job.NameSpace, job.JobUuid, "cron-task abnormal state"); err == nil {
					deleteSpaceIds = append(deleteSpaceIds, job.SpaceUuid+"_"+job.JobUuid)
					continue
				}
			}

			if time.Now().Unix() > job.ExpireTime {
				// Compatible with old versions
				DeleteJob(job.NameSpace, job.SpaceUuid, "compatible with old versions, cron-task the task execution time has expired")
				if err = DeleteJob(job.NameSpace, job.JobUuid, "cron-task the task execution time has expired"); err == nil {
					deleteSpaceIds = append(deleteSpaceIds, job.SpaceUuid)
					deleteJobIds = append(deleteJobIds, job.JobUuid)
					continue
				}
			}

		}

		for _, spaceUuid := range deleteSpaceIds {
			split := strings.Split(spaceUuid, "_")
			if len(split) == 2 {
				NewJobService().DeleteJobEntityBySpaceUuId(split[0], split[1], models.JOB_COMPLETED_STATUS)
			}
		}

		for _, jobUuid := range deleteJobIds {
			logs.GetLogger().Infof("corn-task starting delete job, job_uuid: %s", jobUuid)
			NewJobService().DeleteJobEntityByJobUuId(jobUuid, models.JOB_COMPLETED_STATUS)
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
		err := NewTaskService().Model(&models.TaskEntity{}).Where("status in (?,?) and create_time <?", models.TASK_RECEIVED_STATUS, models.TASK_RUNNING_STATUS, oneHourAgo).Find(&taskList).Error
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

			if ubiTask.Contract != "" || ubiTask.BlockHash != "" {
				ubiTask.Status = models.TASK_SUBMITTED_STATUS
			} else {
				ubiTask.Status = models.TASK_FAILED_STATUS
			}

			NewTaskService().SaveTaskEntity(&ubiTask)
		}
	})
	c.Start()
}

func (task *CronTask) checkJobReward() {
	c := cron.New(cron.WithSeconds())
	c.AddFunc("* * 0/20 * * ?", func() {
		defer func() {
			if err := recover(); err != nil {
				logs.GetLogger().Errorf("task job: [checkJobReward], error: %+v", err)
			}
		}()

		jobList, err := NewJobService().GetJobListByNoReward()
		if err != nil {
			logs.GetLogger().Errorf("failed to get job data, error: %+v", err)
			return
		}
		for _, job := range jobList {
			jobCopy := job
			go NewTaskManagerContract().Scan(jobCopy)
		}
	})
	c.Start()
}

func (task *CronTask) getUbiTaskReward() {
	c := cron.New(cron.WithSeconds())
	c.AddFunc("* 0/10 * * * ?", func() {
		defer func() {
			if err := recover(); err != nil {
				logs.GetLogger().Errorf("task job: [GetUbiTaskReward], error: %+v", err)
			}
		}()
		if err := syncTaskStatusForSequencerService(); err != nil {
			logs.GetLogger().Errorf("failed to sync task from sequencer, error: %v", err)
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

	expiredTime := taskInfo.StartTimestamp + taskInfo.Duration
	if expiredTime > 0 {
		job.ExpireTime = expiredTime
	}
	if err := NewJobService().UpdateJobEntityByJobUuid(job); err != nil {
		logs.GetLogger().Errorf("failed to update job info, space_uuid: %s, error: %v", job.SpaceUuid, err)
	}
}

type TaskGroup struct {
	Items []*models.TaskEntity
	Ids   []int64
	Type  int // 1: contract  2: sequncer
}

func handleTasksToGroup(list []*models.TaskEntity) []TaskGroup {
	var groups []TaskGroup
	var group TaskGroup
	var group1 TaskGroup

	const batchSize = 20
	for i := 0; i < len(list); i++ {
		if list[i].Sequencer == 1 {
			if len(group1.Items) > batchSize {
				groups = append(groups, group1)
				group1 = TaskGroup{}
			}
			group1.Items = append(group1.Items, list[i])
			group1.Ids = append(group1.Ids, list[i].Id)
			group1.Type = 2
		} else if list[i].Sequencer == 0 {
			if len(group.Items) > batchSize {
				groups = append(groups, group)
				group = TaskGroup{}
			}
			group.Items = append(group.Items, list[i])
			group.Ids = append(group.Ids, list[i].Id)
			group.Type = 1
		}
	}
	if len(group.Items) > 0 {
		groups = append(groups, group)
	}
	if len(group1.Items) > 0 {
		groups = append(groups, group1)
	}
	return groups
}

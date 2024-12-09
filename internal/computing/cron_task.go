package computing

import (
	"context"
	"fmt"
	"github.com/swanchain/go-computing-provider/internal/contract"
	"strconv"
	"strings"
	"sync"
	"time"

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
	checkJobStatus()
	task.addLabelToNode()
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

func CheckClusterNetworkPolicy() {
	var err error
	NetworkPolicyFlag = false
	netset, err := NewK8sService().GetGlobalNetworkSet(models.NetworkNetset)
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}
	if netset.GetName() == "" {
		return
	}

	if netset != nil {
		subNetGnp, err := NewK8sService().GetGlobalNetworkPolicy(models.NetworkGlobalSubnet)
		if err != nil {
			logs.GetLogger().Error(err)
			return
		}
		if subNetGnp.Name == "" {
			return
		}

		outAccessGnp, err := NewK8sService().GetGlobalNetworkPolicy(models.NetworkGlobalOutAccess)
		if err != nil {
			logs.GetLogger().Error(err)
			return
		}
		if outAccessGnp.Name == "" {
			return
		}

		inAccessGnp, err := NewK8sService().GetGlobalNetworkPolicy(models.NetworkGlobalInAccess)
		if err != nil {
			logs.GetLogger().Error(err)
			return
		}
		if inAccessGnp.Name == "" {
			return
		}

		namespaceGnp, err := NewK8sService().GetGlobalNetworkPolicy(models.NetworkGlobalNamespace)
		if err != nil {
			logs.GetLogger().Error(err)
			return
		}
		if namespaceGnp.Name == "" {
			return
		}

		dnsGnp, err := NewK8sService().GetGlobalNetworkPolicy(models.NetworkGlobalDns)
		if err != nil {
			logs.GetLogger().Error(err)
			return
		}
		if dnsGnp.Name == "" {
			return
		}

		podInNsGnp, err := NewK8sService().GetGlobalNetworkPolicy(models.NetworkGlobalPodInNamespace)
		if err != nil {
			logs.GetLogger().Error(err)
			return
		}
		if podInNsGnp.Name == "" {
			return
		}
	}

	NetworkPolicyFlag = true
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

func (task *CronTask) addLabelToNode() {
	c := cron.New(cron.WithSeconds())
	c.AddFunc("0 */10 * * * ?", func() {
		defer func() {
			if err := recover(); err != nil {
				logs.GetLogger().Errorf("failed to add label for cluster node, error: %+v", err)
			}
		}()
		addNodeLabel()
	})
	c.Start()
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
		NewDockerService().CleanResourceForK8s()
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

		var deleteSpaceIdAndJobUuid = make(map[string]string)
		for _, jobCopy := range jobList {
			job := jobCopy
			jobUuidDeployName := constants.K8S_DEPLOY_NAME_PREFIX + strings.ToLower(job.JobUuid)
			if _, ok := deployOnK8s[jobUuidDeployName]; ok {
				var nameSpace = job.NameSpace
				if job.Status == models.JOB_TERMINATED_STATUS || job.Status == models.JOB_COMPLETED_STATUS {
					if strings.TrimSpace(nameSpace) == "" {
						nameSpace = constants.K8S_NAMESPACE_NAME_PREFIX + strings.ToLower(job.WalletAddress)
					}
					if err = DeleteJob(nameSpace, job.JobUuid, "cron-task abnormal state"); err != nil {
						logs.GetLogger().Errorf("failed to use jobUuid: %s delete job, error: %v", job.JobUuid, err)
					}
					continue
				}
			} else {
				// compatible with space_uuid
				spaceUuidDeployName := constants.K8S_DEPLOY_NAME_PREFIX + strings.ToLower(job.SpaceUuid)
				if _, ok = deployOnK8s[spaceUuidDeployName]; ok {
					if NewJobService().GetJobEntityBySpaceUuid(job.SpaceUuid) > 0 && time.Now().Unix() < job.ExpireTime {
						if job.Status != models.JOB_RUNNING_STATUS {
							foundDeployment, err := NewK8sService().k8sClient.AppsV1().Deployments(job.NameSpace).Get(context.TODO(), job.K8sDeployName, metav1.GetOptions{})
							if err != nil {
								continue
							}
							if foundDeployment.Status.AvailableReplicas > 0 {
								job.PodStatus = models.POD_RUNNING_STATUS
								job.Status = models.JOB_RUNNING_STATUS
								NewJobService().UpdateJobEntityByJobUuid(job)
							}
						}
						continue
					}

					var nameSpace = job.NameSpace
					if job.Status == models.JOB_TERMINATED_STATUS || job.Status == models.JOB_COMPLETED_STATUS {
						if strings.TrimSpace(nameSpace) == "" {
							nameSpace = constants.K8S_NAMESPACE_NAME_PREFIX + strings.ToLower(job.WalletAddress)
						}

						if err = DeleteJob(nameSpace, job.SpaceUuid, "cron-task abnormal state, compatible with space_uuid"); err != nil {
							logs.GetLogger().Errorf("failed to use spaceUuid: %s delete job, error: %v", job.SpaceUuid, err)
						}
						continue
					}
				}
			}

			if job.DeleteAt == models.DELETED_FLAG {
				continue
			}

			currentTime := time.Now()
			createdTime := time.Unix(job.CreateTime, 0)
			createDuration := currentTime.Sub(createdTime)

			if job.NameSpace != "" && job.K8sDeployName != "" {
				foundDeployment, err := NewK8sService().k8sClient.AppsV1().Deployments(job.NameSpace).Get(context.TODO(), job.K8sDeployName, metav1.GetOptions{})
				if err != nil {
					if createDuration.Hours() <= 2 && job.Status != models.JOB_RUNNING_STATUS {
						continue
					}
					if errors.IsNotFound(err) {
						// delete job
						logs.GetLogger().Warnf("not found deployment on the cluster, job_uuid: %s, deployment: %s", job.JobUuid, job.K8sDeployName)
						deleteSpaceIdAndJobUuid[job.JobUuid] = job.SpaceUuid + "_" + job.JobUuid
						continue
					}
					logs.GetLogger().Errorf("failed to get deployment: %s, error: %v", job.K8sDeployName, err)
					continue
				}

				if foundDeployment.Status.AvailableReplicas == 0 && createDuration.Hours() > 2 { // need to delete
					DeleteJob(job.NameSpace, job.JobUuid, "cron-task correction status")
					deleteSpaceIdAndJobUuid[job.JobUuid] = job.SpaceUuid + "_" + job.JobUuid
					continue
				} else {
					if job.Status != models.JOB_RUNNING_STATUS {
						job.PodStatus = models.POD_RUNNING_STATUS
						job.Status = models.JOB_RUNNING_STATUS
						NewJobService().UpdateJobEntityByJobUuid(job)
					}
				}
			}

			checkFcpJobInfoInChain(job)

			if job.Status == models.JOB_TERMINATED_STATUS || job.Status == models.JOB_COMPLETED_STATUS || time.Now().Unix() > job.ExpireTime {
				expireTime := time.Unix(job.ExpireTime, 0).Format("2006-01-02 15:04:05")
				logs.GetLogger().Infof("job_uuid: %s, current status is %s, expire time: %s, starting to delete it.", job.JobUuid, models.GetJobStatus(job.Status), expireTime)
				if err = DeleteJob(job.NameSpace, job.JobUuid, "cron-task abnormal state"); err != nil {
					logs.GetLogger().Errorf("failed to use jobUuid: %s delete job, error: %v", job.JobUuid, err)
					continue
				}

				if err = DeleteJob(job.NameSpace, job.SpaceUuid, "compatible with old versions, cron-task abnormal state"); err != nil {
					logs.GetLogger().Errorf("failed to use spaceUuid: %s delete job, error: %v", job.SpaceUuid, err)
					continue
				}
				deleteSpaceIdAndJobUuid[job.JobUuid] = job.SpaceUuid + "_" + job.JobUuid
			}
		}

		for _, spaceUuidAndJobUuid := range deleteSpaceIdAndJobUuid {
			split := strings.Split(spaceUuidAndJobUuid, "_")
			if len(split) == 2 {
				NewJobService().DeleteJobEntityBySpaceUuId(split[0], split[1], models.JOB_COMPLETED_STATUS)
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
	//c.AddFunc("* * 0/20 * * ?", func() {
	c.AddFunc("0 */10 * * * ?", func() {
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
			if NewTaskManagerContract(job).Scan() {
				break
			}
		}
	})
	c.Start()
}

func (task *CronTask) getUbiTaskReward() {
	c := cron.New(cron.WithSeconds())
	c.AddFunc("0 */10 * * * ?", func() {
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

	for _, node := range nodes.Items {
		cpNode := node
		if collectInfo, ok := nodeGpuInfoMap[cpNode.Name]; ok {
			for _, detail := range collectInfo.Gpu.Details {
				if err = k8sService.AddNodeLabel(cpNode.Name, detail.ProductName); err != nil {
					logs.GetLogger().Errorf("add node label, nodeName %s, gpuName: %s, error: %+v", cpNode.Name, detail.ProductName, err)
					continue
				}
			}
			err := k8sService.AddNodeLabelForArchitecture(cpNode.Name, collectInfo.CpuName)
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
	client, err := contract.GetEthClient(chainRpc)
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
	client, err := contract.GetEthClient(chainRpc)
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
	if taskInfo.TaskUuid != "" {
		if taskInfo.TaskStatus == models.COMPLETED {
			job.Status = models.JOB_COMPLETED_STATUS
		} else if taskInfo.TaskStatus == models.TERMINATED {
			job.Status = models.JOB_TERMINATED_STATUS
		}

		expiredTime := taskInfo.StartTimestamp + taskInfo.Duration
		if expiredTime > 0 {
			job.ExpireTime = expiredTime
			if err := NewJobService().UpdateJobEntityByJobUuid(job); err != nil {
				logs.GetLogger().Errorf("update job info by jobUuid failed, error: %v", err)
			}
		}
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

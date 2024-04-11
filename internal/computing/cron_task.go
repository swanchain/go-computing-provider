package computing

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/filswan/go-mcs-sdk/mcs/api/common/logs"
	"github.com/gomodule/redigo/redis"
	"github.com/robfig/cron/v3"
	"github.com/swanchain/go-computing-provider/conf"
	"github.com/swanchain/go-computing-provider/constants"
	"io"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"math"
	"net/http"
	"strings"
	"time"
)

type CronTask struct {
}

func NewCronTask() *CronTask {
	return &CronTask{}
}

func (task *CronTask) RunTask() {
	task.checkCollateralBalance()
	task.cleanAbnormalDeployment()
	task.setFailedUbiTaskStatus()
}

func (task *CronTask) checkCollateralBalance() {
	c := cron.New(cron.WithSeconds())
	c.AddFunc("0/15 * * * * ?", func() {
		defer func() {
			if err := recover(); err != nil {
				logs.GetLogger().Errorf("task job: [checkCollateralBalance], error: %+v", err)
			}
		}()

		url := fmt.Sprintf("%s/cp/collateral/%s", conf.GetConfig().HUB.ServerUrl, conf.GetConfig().HUB.WalletAddress)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			logs.GetLogger().Errorf("create req failed: %+v", err)
			return
		}
		req.Header.Set("Authorization", "Bearer "+conf.GetConfig().HUB.AccessToken)

		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			logs.GetLogger().Errorf("send req failed: %+v", err)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			logs.GetLogger().Errorf("read response failed: %+v", err)
			return
		}
		var collateral struct {
			Data struct {
				Balance float64 `json:"balance"`
			} `json:"data"`
			Message string `json:"message"`
			Status  string `json:"status"`
		}
		err = json.Unmarshal(body, &collateral)
		if err != nil {
			logs.GetLogger().Errorf("json conversion failed: %+v", err)
			return
		}

		result := collateral.Data.Balance / 1e18
		result = math.Round(result*1000) / 1000
		if result <= conf.GetConfig().HUB.BalanceThreshold {
			logs.GetLogger().Warnf("No sufficient collateral Balance, the current collateral balance is: %0.3f. Please run: computing-provider collateral [fromWalletAddress] [amount]", result)
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
					if (deployment.Status.AvailableReplicas == 0 && age.Hours() >= 2) || age.Hours() > 24*15 {
						logs.GetLogger().Infof("Cleaning up deployment %s in namespace %s", deployment.Name, namespace)
						err := k8sService.k8sClient.AppsV1().Deployments(namespace).Delete(context.TODO(), deployment.Name, metav1.DeleteOptions{})
						if err != nil {
							if errors.IsNotFound(err) {
								logs.GetLogger().Errorf("Deployment %s not found. Ignoring", deployment.Name)
							} else {
								logs.GetLogger().Errorf("Error deleting deployment %s: %v", deployment.Name, err)
							}
						} else {
							logs.GetLogger().Errorf("Deployment %s deleted successfully.", deployment.Name)
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
				logs.GetLogger().Errorf("task job: [cleanAbnormalDeployment], error: %+v", err)
			}
		}()

		conn := redisPool.Get()
		prefix := constants.REDIS_UBI_C2_PERFIX + "*"
		keys, err := redis.Strings(conn.Do("KEYS", prefix))
		if err != nil {
			logs.GetLogger().Errorf("Failed get redis %s prefix, error: %+v", prefix, err)
			return
		}
		for _, key := range keys {
			ubiTask, err := RetrieveUbiTaskMetadata(key)
			if err != nil {
				logs.GetLogger().Errorf("Failed get ubi task from redis, key: %s, error: %+v", key, err)
				return
			}

			JobName := strings.ToLower(ubiTask.ZkType) + "-" + ubiTask.TaskId
			k8sNameSpace := "ubi-task-" + ubiTask.TaskId

			service := NewK8sService()
			if _, err = service.k8sClient.BatchV1().Jobs(k8sNameSpace).Get(context.TODO(), JobName, metav1.GetOptions{}); err != nil && errors.IsNotFound(err) {
				if ubiTask.Status != constants.UBI_TASK_SUCCESS_STATUS && ubiTask.Status != constants.UBI_TASK_FAILED_STATUS {
					ubiTask.Status = constants.UBI_TASK_FAILED_STATUS
				}
				SaveUbiTaskMetadata(ubiTask)
			}
		}
	})
	c.Start()
}

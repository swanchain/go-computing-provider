package computing

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/filswan/go-mcs-sdk/mcs/api/common/logs"
	"github.com/gomodule/redigo/redis"
	"github.com/lagrangedao/go-computing-provider/conf"
	"github.com/lagrangedao/go-computing-provider/constants"
	"github.com/lagrangedao/go-computing-provider/models"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func RunSyncTask() {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				logs.GetLogger().Errorf("Failed report cp resource's summary, error: %+v", err)
			}
		}()
		ticker := time.NewTicker(120 * time.Second)
		defer ticker.Stop()

		nodeId, _, _ := generateNodeID()
		location, err := getLocation()
		if err != nil {
			logs.GetLogger().Error(err)
		}

		reportClusterResource(location, nodeId)
		for range ticker.C {
			reportClusterResource(location, nodeId)
		}
	}()

	watchExpiredTask()
	watchNameSpaceForDeleted()
}

func reportClusterResource(location, nodeId string) {
	k8sService := NewK8sService()
	statisticalSources, err := k8sService.StatisticalSources(context.TODO())
	if err != nil {
		logs.GetLogger().Errorf("Failed k8s statistical sources, error: %+v", err)
		return
	}
	clusterSource := models.ClusterResource{
		NodeId:      nodeId,
		Region:      location,
		ClusterInfo: statisticalSources,
	}

	payload, err := json.Marshal(clusterSource)
	if err != nil {
		logs.GetLogger().Errorf("Failed convert to json, error: %+v", err)
		return
	}

	resp, err := http.Post(conf.GetConfig().LAD.ServerUrl+"/cp/summary", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		logs.GetLogger().Errorf("Failed send a request, error: %+v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logs.GetLogger().Errorf("The request returns a non-200 status code: %d", resp.StatusCode)
		return
	}
	logs.GetLogger().Info("report cluster node resources successfully")
}

func watchExpiredTask() {
	ticker := time.NewTicker(30 * time.Second)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				logs.GetLogger().Errorf("catch panic error: %+v", err)
			}
		}()

		var deleteKey []string
		for range ticker.C {
			conn := redisPool.Get()
			cursor := "0"
			prefix := constants.REDIS_FULL_PREFIX + "*"
			values, err := redis.Values(conn.Do("SCAN", cursor, "MATCH", prefix))
			if err != nil {
				logs.GetLogger().Errorf("Failed scan redis %s prefix, error: %+v", prefix, err)
				return
			}

			cursor, _ = redis.String(values[0], nil)
			keys, _ := redis.Strings(values[1], nil)
			for _, key := range keys {
				args := []interface{}{key}
				args = append(args, "k8s_namespace", "space_name", "expire_time")
				valuesStr, err := redis.Strings(conn.Do("HMGET", args...))
				if err != nil {
					logs.GetLogger().Errorf("Failed get redis key data, key: %s, error: %+v", key, err)
					return
				}

				if len(valuesStr) >= 3 {
					namespace := valuesStr[0]
					spaceName := valuesStr[1]
					expireTimeStr := valuesStr[2]
					expireTime, err := strconv.ParseInt(strings.TrimSpace(expireTimeStr), 10, 64)
					if err != nil {
						logs.GetLogger().Errorf("Failed convert time str: [%s], error: %+v", expireTimeStr, err)
						return
					}
					if time.Now().Unix() > expireTime {
						logs.GetLogger().Infof("The namespace: %s, spacename: %s, job has reached its runtime and will stop running.", namespace, spaceName)
						deleteJob(namespace, spaceName)
						deleteKey = append(deleteKey, key)
					}
				}
			}
			conn.Do("DEL", redis.Args{}.AddFlat(deleteKey)...)
			if len(deleteKey) > 0 {
				logs.GetLogger().Infof("Delete redis keys finished, keys: %+v", deleteKey)
				deleteKey = nil
			}

			if cursor == "0" {
				break
			}
		}
	}()
}

func watchNameSpaceForDeleted() {
	ticker := time.NewTicker(24 * time.Hour)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				logs.GetLogger().Errorf("catch panic error: %+v", err)
			}
		}()

		for range ticker.C {
			service := NewK8sService()
			namespaces, err := service.ListNamespace(context.TODO())
			if err != nil {
				logs.GetLogger().Errorf("Failed get all namespace, error: %+v", err)
				continue
			}

			for _, namespace := range namespaces {
				getPods, err := service.GetPods(namespace, "")
				if err != nil {
					logs.GetLogger().Errorf("Failed get pods form namespace,namepace: %s, error: %+v", namespace, err)
					continue
				}
				if !getPods && strings.HasPrefix(namespace, constants.K8S_NAMESPACE_NAME_PREFIX) {
					if err = service.DeleteNameSpace(context.TODO(), namespace); err != nil {
						logs.GetLogger().Errorf("Failed delete namespace, namepace: %s, error: %+v", namespace, err)
					}
				}
			}
		}
	}()
}
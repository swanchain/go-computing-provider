package computing

import (
	"encoding/json"
	"fmt"
	"github.com/filswan/go-mcs-sdk/mcs/api/common/logs"
	"github.com/swanchain/go-computing-provider/internal/models"
	corev1 "k8s.io/api/core/v1"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	ResourceCpu     string = "cpu"
	ResourceMem     string = "mem"
	ResourceStorage string = "storage"
)

type CpResourceSummary struct {
	ClusterInfo []*models.NodeResource
}

func GetNodeResource(allPods []corev1.Pod, node *corev1.Node) (map[string]GpuData, map[string]int64, *models.NodeResource) {
	var (
		usedCpu     int64
		usedMem     int64
		usedStorage int64
	)
	nodeGpu := make(map[string]GpuData)
	remainderResource := make(map[string]int64)

	var nodeResource = new(models.NodeResource)
	nodeResource.MachineId = node.Status.NodeInfo.MachineID

	for _, pod := range getPodsFromNode(allPods, node) {
		usedCpu += cpuInPod(&pod)
		usedMem += memInPod(&pod)
		usedStorage += storageInPod(&pod)

		// used gpu info
		gpuName, usedCount, usedIndexs := gpuInPod(&pod)
		if gpuName == "kubernetes.io/os" || gpuName == "" {
			continue
		}
		if v, ok := nodeGpu[gpuName]; ok {
			v.UsedIndex = append(v.UsedIndex, usedIndexs...)
			v.Used += usedCount
			nodeGpu[gpuName] = v
		} else {
			nodeGpu[gpuName] = GpuData{
				Used:      usedCount,
				UsedIndex: usedIndexs,
			}
		}
	}

	nodeResource.Cpu.Total = strconv.FormatInt(node.Status.Capacity.Cpu().Value(), 10)
	nodeResource.Cpu.Used = strconv.FormatInt(usedCpu, 10)
	nodeResource.Cpu.Free = strconv.FormatInt(node.Status.Capacity.Cpu().Value()-usedCpu, 10)
	nodeResource.Cpu.RemainderNum = node.Status.Capacity.Cpu().Value() - usedCpu
	remainderResource[ResourceCpu] = node.Status.Capacity.Cpu().Value() - usedCpu

	nodeResource.Vcpu.Total = nodeResource.Cpu.Total
	nodeResource.Vcpu.Used = nodeResource.Cpu.Used
	nodeResource.Vcpu.Free = nodeResource.Cpu.Free

	nodeResource.Memory.Total = fmt.Sprintf("%.2f GiB", float64(node.Status.Allocatable.Memory().Value()/1024/1024/1024))
	nodeResource.Memory.Used = fmt.Sprintf("%.2f GiB", float64(usedMem/1024/1024/1024))
	freeMemory := node.Status.Capacity.Memory().Value() - usedMem
	nodeResource.Memory.Free = fmt.Sprintf("%.2f GiB", float64(freeMemory/1024/1024/1024))
	nodeResource.Memory.RemainderNum = freeMemory
	remainderResource[ResourceMem] = freeMemory

	nodeResource.Storage.Total = fmt.Sprintf("%.2f GiB", float64(node.Status.Allocatable.StorageEphemeral().Value()/1024/1024/1024))
	nodeResource.Storage.Used = fmt.Sprintf("%.2f GiB", float64(usedStorage/1024/1024/1024))
	freeStorage := node.Status.Allocatable.StorageEphemeral().Value() - usedStorage
	nodeResource.Storage.Free = fmt.Sprintf("%.2f GiB", float64(freeStorage/1024/1024/1024))
	nodeResource.Storage.RemainderNum = freeStorage
	remainderResource[ResourceStorage] = freeStorage

	return nodeGpu, remainderResource, nodeResource
}

func getPodsFromNode(allPods []corev1.Pod, node *corev1.Node) (pods []corev1.Pod) {
	for _, pod := range allPods {
		if pod.Spec.NodeName == node.Name {
			pods = append(pods, pod)
		}
	}
	return pods
}

func storageInPod(pod *corev1.Pod) (storageUsed int64) {
	containers := pod.Spec.Containers
	for _, container := range containers {
		val, ok := container.Resources.Requests[corev1.ResourceEphemeralStorage]
		if !ok {
			continue
		}
		storageUsed += val.Value()
	}
	return storageUsed
}

func cpuInPod(pod *corev1.Pod) (cpuCount int64) {
	containers := pod.Spec.Containers
	for _, container := range containers {
		val, ok := container.Resources.Requests[corev1.ResourceCPU]
		if !ok {
			continue
		}
		cpuCount += val.Value()
	}
	return cpuCount
}

func memInPod(pod *corev1.Pod) (memCount int64) {
	containers := pod.Spec.Containers
	for _, container := range containers {
		val, ok := container.Resources.Requests[corev1.ResourceMemory]
		if !ok {
			continue
		}
		memCount += val.Value()
	}
	return memCount
}

func gpuInPod(pod *corev1.Pod) (gpuName string, gpuCount int, indexs []string) {
	containers := pod.Spec.Containers
	for _, container := range containers {
		var gIndex string
		for _, envVaule := range container.Env {
			if envVaule.Name == "NVIDIA_VISIBLE_DEVICES" {
				gIndex = envVaule.Value
				indexs = append(indexs, strings.Split(gIndex, ",")...)
				break
			}
		}
		if gIndex != "" {
			gpuCount += len(strings.Split(gIndex, ","))
		}
	}

	if pod.Spec.NodeSelector != nil {
		for k := range pod.Spec.NodeSelector {
			if k != "" {
				gpuName = k
			}
		}
	}
	return gpuName, gpuCount, indexs
}

func checkClusterProviderStatus(nodeResources []*models.NodeResource) {
	var policy models.ResourcePolicy
	cpPath, _ := os.LookupEnv("CP_PATH")
	resourcePolicy := filepath.Join(cpPath, "resource_policy.json")
	bytes, err := os.ReadFile(resourcePolicy)
	if err != nil {
		policy = defaultResourcePolicy()
	} else {
		if err = json.Unmarshal(bytes, &policy); err != nil {
			logs.GetLogger().Errorf("parse json failed, error: %v", err)
			return
		}
	}

	for _, node := range nodeResources {
		if node.Cpu.RemainderNum < policy.Cpu.Quota {
			logs.GetLogger().Warningf("Insufficient cpu resources, current cpu resource: %s less than %d", node.Cpu.Free, policy.Cpu.Quota)
			return
		}
		if node.Memory.RemainderNum < policy.Memory.Quota {
			logs.GetLogger().Warningf("Insufficient memory resources, current memory resource: %s less than %d %s", node.Memory.Free, policy.Memory.Quota, policy.Memory.Unit)
			return
		}
		if node.Storage.RemainderNum < policy.Storage.Quota {
			logs.GetLogger().Warningf("Insufficient storage resources, current storage resource: %s less than %d %s", node.Storage.Free, policy.Storage.Quota, policy.Storage.Unit)
			return
		}
	}
}

func defaultResourcePolicy() models.ResourcePolicy {
	return models.ResourcePolicy{
		Cpu: models.CpuQuota{
			Quota: 1,
		},
		Memory: models.Quota{
			Quota: 5,
			Unit:  "GiB",
		},
		Storage: models.Quota{
			Quota: 10,
			Unit:  "GiB",
		},
	}
}

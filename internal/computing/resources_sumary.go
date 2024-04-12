package computing

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/filswan/go-mcs-sdk/mcs/api/common/logs"
	"github.com/swanchain/go-computing-provider/internal/models"
	corev1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
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

func allActivePods(clientSet *kubernetes.Clientset) ([]corev1.Pod, error) {
	allPods, err := clientSet.CoreV1().Pods("").List(context.TODO(), metaV1.ListOptions{
		FieldSelector: "status.phase=Running",
	})
	if err != nil {
		return nil, err
	}
	return allPods.Items, nil
}

func GetNodeResource(allPods []corev1.Pod, node *corev1.Node) (map[string]int64, map[string]int64, *models.NodeResource) {
	var (
		usedCpu     int64
		usedMem     int64
		usedStorage int64
	)
	nodeGpu := make(map[string]int64)
	remainderResource := make(map[string]int64)

	var nodeResource = new(models.NodeResource)
	nodeResource.MachineId = node.Status.NodeInfo.MachineID
	nodeResource.Model = node.Status.NodeInfo.Architecture

	for _, pod := range getPodsFromNode(allPods, node) {
		usedCpu += cpuInPod(&pod)
		usedMem += memInPod(&pod)
		usedStorage += storageInPod(&pod)

		gpuName, count := gpuInPod(&pod)
		if v, ok := nodeGpu[gpuName]; ok {
			nodeGpu[gpuName] = v + count
		} else {
			nodeGpu[gpuName] = count
		}
	}

	nodeResource.Cpu.Total = strconv.FormatInt(node.Status.Capacity.Cpu().Value(), 10)
	nodeResource.Cpu.Used = strconv.FormatInt(usedCpu, 10)
	nodeResource.Cpu.Free = strconv.FormatInt(node.Status.Capacity.Cpu().Value()-usedCpu, 10)
	remainderResource[ResourceCpu] = node.Status.Capacity.Cpu().Value() - usedCpu

	nodeResource.Vcpu.Total = nodeResource.Cpu.Total
	nodeResource.Vcpu.Used = nodeResource.Cpu.Used
	nodeResource.Vcpu.Free = nodeResource.Cpu.Free

	nodeResource.Memory.Total = fmt.Sprintf("%.2f GiB", float64(node.Status.Allocatable.Memory().Value()/1024/1024/1024))
	nodeResource.Memory.Used = fmt.Sprintf("%.2f GiB", float64(usedMem/1024/1024/1024))
	freeMemory := node.Status.Capacity.Memory().Value() - usedMem
	nodeResource.Memory.Free = fmt.Sprintf("%.2f GiB", float64(freeMemory/1024/1024/1024))
	remainderResource[ResourceMem] = freeMemory

	nodeResource.Storage.Total = fmt.Sprintf("%.2f GiB", float64(node.Status.Allocatable.StorageEphemeral().Value()/1024/1024/1024))
	nodeResource.Storage.Used = fmt.Sprintf("%.2f GiB", float64(usedStorage/1024/1024/1024))
	freeStorage := node.Status.Allocatable.StorageEphemeral().Value() - usedStorage
	nodeResource.Storage.Free = fmt.Sprintf("%.2f GiB", float64(freeStorage/1024/1024/1024))
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

func gpuInPod(pod *corev1.Pod) (gpuName string, gpuCount int64) {
	containers := pod.Spec.Containers
	for _, container := range containers {
		val, ok := container.Resources.Requests["nvidia.com/gpu"]
		if !ok {
			continue
		}
		gpuCount += val.Value()
	}

	if pod.Spec.NodeSelector != nil {
		for k := range pod.Spec.NodeSelector {
			if k != "" {
				gpuName = k
			}
		}
	}
	return gpuName, gpuCount
}

func checkClusterProviderStatus() (string, error) {

	var policy models.ResourcePolicy
	currentDir, _ := os.Getwd()
	resourcePolicy := filepath.Join(currentDir, "resource_policy.json")
	bytes, err := os.ReadFile(resourcePolicy)
	if err != nil {
		policy = defaultResourcePolicy()
	} else {
		if err = json.Unmarshal(bytes, &policy); err != nil {
			return "", err
		}
	}

	service := NewK8sService()
	activePods, err := allActivePods(service.Client)
	if err != nil {
		return "", err
	}

	nodes, err := service.Client.CoreV1().Nodes().List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return "", err
	}

	collectGpu := make(map[string]int64)
	nodeGpuInfo, err := service.GetResourceExporterPodLog(context.TODO())
	if err != nil {
		logs.GetLogger().Error(err)
		return "", err
	}
	for _, gpu := range nodeGpuInfo {
		for _, gpuDetail := range gpu.Gpu.Details {
			collectGpu[gpuDetail.ProductName] = collectGpu[gpuDetail.ProductName] + 1
		}
	}

	nodeGpu := make(map[string]int64)
	nodeResource := make(map[string]int64)
	for _, node := range nodes.Items {
		gpuMap, remainderResource, _ := GetNodeResource(activePods, &node)
		for k, v := range gpuMap {
			nodeGpu[k] = nodeGpu[k] + v
		}
		for k, v := range remainderResource {
			nodeResource[k] = nodeGpu[k] + v
		}
	}

	remainGpu := make(map[string]int64)
	policyMap := make(map[string]int64)
	for name, num := range collectGpu {
		gpuName := strings.ReplaceAll(name, " ", "-")
		remainGpu[gpuName] = num - nodeGpu[gpuName]

		for _, gpu := range policy.Gpu {
			upperName := strings.ReplaceAll(gpu.Name, "Nvidia", "NVIDIA")
			if gpuName == upperName {
				policyMap[gpuName] = gpu.Quota
				break
			}
		}
	}

	var gpuFlag bool
	for name, quota := range policyMap {
		if remainGpu[name] > quota {
			gpuFlag = true
			break
		}
	}

	if gpuFlag {
		if nodeResource[ResourceCpu] < policy.Cpu.Quota || nodeResource[ResourceMem] < policy.Memory.Quota || nodeResource[ResourceStorage] < policy.Memory.Quota {
			logs.GetLogger().Infof("have gpu, status: %s", models.InactiveStatus)
			return models.InactiveStatus, nil
		}
		return models.ActiveStatus, nil
	} else {
		if nodeResource[ResourceCpu] < policy.Cpu.Quota || nodeResource[ResourceMem] < policy.Memory.Quota || nodeResource[ResourceStorage] < policy.Memory.Quota {
			logs.GetLogger().Infof("no gpu, status: %s", models.InactiveStatus)
			return models.InactiveStatus, nil
		} else {
			return models.ActiveStatus, nil
		}
	}
}

func defaultResourcePolicy() models.ResourcePolicy {
	return models.ResourcePolicy{
		Cpu: models.CpuQuota{
			Quota: 0,
		},
		Memory: models.Quota{
			Quota: 0,
			Unit:  "GiB",
		},
		Storage: models.Quota{
			Quota: 0,
			Unit:  "GiB",
		},
		Gpu: []models.GpuQuota{
			{
				Name:  "Nvidia-2060",
				Quota: 0,
			},
			{
				Name:  "Nvidia-3070",
				Quota: 0,
			},
			{
				Name:  "Nvidia-3080",
				Quota: 0,
			},
			{
				Name:  "Nvidia-3090",
				Quota: 0,
			},
			{
				Name:  "Nvidia-4090",
				Quota: 0,
			},
			{
				Name:  "Nvidia-A100",
				Quota: 0,
			},
			{
				Name:  "Nvidia-H100",
				Quota: 0,
			},
			{
				Name:  "Nvidia-A10G",
				Quota: 0,
			}, {
				Name:  "Nvidia-T4",
				Quota: 0,
			},
			{
				Name:  "Nvidia-2080-Ti",
				Quota: 0,
			},
			{
				Name:  "Nvidia-3060-Ti",
				Quota: 0,
			},
			{
				Name:  "Nvidia-3070-Ti",
				Quota: 0,
			},
			{
				Name:  "Nvidia-3080-Ti",
				Quota: 0,
			},
			{
				Name:  "Nvidia-4090-Ti",
				Quota: 0,
			},
		},
	}
}

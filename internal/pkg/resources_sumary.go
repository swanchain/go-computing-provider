package pkg

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/swanchain/go-computing-provider/internal/v1/models"
	corev1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"os"
	"path/filepath"
	"strconv"
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

func CheckClusterProviderStatus() {
	var policy models.ResourcePolicy
	cpPath, _ := os.LookupEnv("CP_PATH")
	resourcePolicy := filepath.Join(cpPath, "resource_policy.json")
	bytes, err := os.ReadFile(resourcePolicy)
	if err != nil {
		policy = defaultResourcePolicy()
	} else {
		if err = json.Unmarshal(bytes, &policy); err != nil {
			ulog.Errorf("parse json failed, error: %v", err)
			return
		}
	}
	service := NewK8sService()
	activePods, err := allActivePods(service.ClientSet)
	if err != nil {
		ulog.Errorf("get all active pod failed, error: %v", err)
		return
	}

	nodes, err := service.ClientSet.CoreV1().Nodes().List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		ulog.Errorf("get all node failed, error: %v", err)
		return
	}

	for _, node := range nodes.Items {
		_, remainderResource, nodeResource := GetNodeResource(activePods, &node)
		if remainderResource[ResourceCpu] < policy.Cpu.Quota {
			ulog.Warnf("Insufficient cpu resources, current cpu resource: %s less than %d", nodeResource.Cpu.Free, policy.Cpu.Quota)
			return
		}
		if remainderResource[ResourceMem] < policy.Memory.Quota {
			ulog.Warnf("Insufficient memory resources, current memory resource: %s less than %d %s", nodeResource.Memory.Free, policy.Memory.Quota, policy.Memory.Unit)
			return
		}
		if remainderResource[ResourceStorage] < policy.Storage.Quota {
			ulog.Warnf("Insufficient storage resources, current storage resource: %s less than %d %s", nodeResource.Storage.Free, policy.Storage.Quota, policy.Storage.Unit)
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

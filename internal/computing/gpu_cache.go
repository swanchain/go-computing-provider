package computing

import (
	"errors"
	"fmt"
	"github.com/filswan/go-mcs-sdk/mcs/api/common/logs"
	"sync"
)

type GPUInfo struct {
	Name      string
	Total     int
	Used      int
	Available int
	FreeIndex []string
}

type GPUManager struct {
	mu    sync.Mutex
	cache sync.Map
}

func NewGpuManager() *GPUManager {
	return &GPUManager{}
}

func (gm *GPUManager) UpdateGPU(key string, gpuInfo GPUInfo) {
	gm.mu.Lock()
	defer gm.mu.Unlock()
	gm.cache.Store(key, gpuInfo)
}

func (gm *GPUManager) GetGPU(key string) (GPUInfo, bool) {
	gpuInfo, ok := gm.cache.Load(key)
	if ok {
		return gpuInfo.(GPUInfo), true
	}
	return GPUInfo{}, false
}

func (gm *GPUManager) AllocateGPU(key string, count int) ([]string, error) {
	gm.mu.Lock()
	defer gm.mu.Unlock()
	gpuInfo, found := gm.GetGPU(key)
	if !found {
		return nil, fmt.Errorf("%s not found", key)
	}
	logs.GetLogger().Infof("gpu stats：total: %d, used: %d, available: %d, indexs: %v\n", gpuInfo.Total, gpuInfo.Used, gpuInfo.Available, gpuInfo.FreeIndex)
	if gpuInfo.Available < count {
		return nil, fmt.Errorf("insufficient GPU resources")
	}

	var allocateIndex []string
	for i := 0; i < count; i++ {
		if i >= len(gpuInfo.FreeIndex) {
			break
		}
		allocateIndex = append(allocateIndex, gpuInfo.FreeIndex[i])
	}
	remainingGpu := difference(gpuInfo.FreeIndex, allocateIndex)
	gpuInfo.FreeIndex = remainingGpu
	gpuInfo.Used += count
	gpuInfo.Available -= count
	gm.UpdateGPU(key, gpuInfo)
	return allocateIndex, nil
}

func (gm *GPUManager) ReleaseGPU(key string, count int, indexs []string) error {
	gm.mu.Lock()
	defer gm.mu.Unlock()
	gpuInfo, found := gm.GetGPU(key)
	if !found {
		return fmt.Errorf("%s not found", key)
	}
	if gpuInfo.Used < count {
		return errors.New("release count exceeds used GPUs")
	}
	gpuInfo.Used -= count
	gpuInfo.Available += count
	gpuInfo.FreeIndex = append(gpuInfo.FreeIndex, indexs...)
	gm.UpdateGPU(key, gpuInfo)
	return nil
}

func (gm *GPUManager) CheckAvailableGPU() (string, bool) {
	var availableGPU string
	gm.cache.Range(func(key, value interface{}) bool {
		gpuInfo := value.(GPUInfo)
		if gpuInfo.Available > 0 {
			availableGPU = key.(string)
			return false
		}
		return true
	})
	if availableGPU != "" {
		return availableGPU, true
	}
	return "", false
}

func difference(arr1, arr2 []string) []string {
	set := make(map[string]struct{})
	for _, v := range arr2 {
		set[v] = struct{}{}
	}
	var result []string
	for _, v := range arr1 {
		if _, found := set[v]; !found {
			result = append(result, v)
		}
	}
	return result
}

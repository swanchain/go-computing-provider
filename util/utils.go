package util

import (
	"fmt"
	"github.com/Masterminds/semver"
	"net"
	"sync"
	"sync/atomic"
)

func CheckPortAvailability(usedPort map[int32]struct{}) bool {
	startPort := 30000
	endPort := 32767
	var waitCheckPorts []int
	for port := startPort; port <= endPort; port++ {
		if _, ok := usedPort[int32(port)]; ok {
			continue
		}
		waitCheckPorts = append(waitCheckPorts, port)
	}

	var num = len(waitCheckPorts)
	var portCounter atomic.Int64
	var wg sync.WaitGroup
	sem := make(chan struct{}, 50)
	for _, port := range waitCheckPorts {
		wg.Add(1)
		go func(p int) {
			defer wg.Done()
			sem <- struct{}{}
			if IsPortAvailable(p) {
				portCounter.Add(1)
			}
			<-sem
		}(port)
	}
	wg.Wait()

	if len(usedPort) > 0 {
		num = num - 1
	}
	if num-int(portCounter.Load()) < 20 {
		return true
	}
	return false
}

func IsPortAvailable(port int) bool {
	address := fmt.Sprintf(":%d", port)
	ln, err := net.Listen("tcp", address)
	defer func() {
		if ln != nil {
			defer ln.Close()
		}
	}()

	if err != nil {
		return false
	}
	return true
}

func CheckVersion(targetVersionStr, versionStr string) error {
	version, err := semver.NewVersion(versionStr)
	if err != nil {
		return fmt.Errorf("invalid version format: %v", err)
	}
	targetVersion := semver.MustParse(targetVersionStr)
	if version.LessThan(targetVersion) {
		return fmt.Errorf("version is too low: %s (required >= %s)", version, targetVersionStr)
	}
	return nil
}

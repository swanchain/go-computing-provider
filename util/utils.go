package util

import (
	"context"
	"errors"
	"fmt"
	"github.com/filswan/go-swan-lib/logs"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

func CheckPortAvailability(usedPort map[int32]struct{}) bool {
	startPort := 30000
	endPort := 32767

	var wg sync.WaitGroup
	var num int64
	var portCounter atomic.Int64
	for port := startPort; port <= endPort; port++ {
		if _, ok := usedPort[int32(port)]; ok {
			continue
		}
		wg.Add(1)
		go func() {
			if startServer(&wg, port) {
				portCounter.Add(1)
			}
		}()
		num++
	}
	wg.Wait()

	if len(usedPort) > 0 {
		num = num - 1
	}
	if num == portCounter.Load() {
		return true
	}
	return false
}

func CheckPortAvailable(port int) bool {
	var flag bool
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		if startServer(&wg, port) {
			flag = true
		}
	}()
	wg.Wait()
	return flag
}

func startServer(wg *sync.WaitGroup, port int) bool {
	defer wg.Done()
	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", port),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		}),
	}
	defer srv.Shutdown(context.TODO())

	var isListen bool
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			isListen = true
		}
	}()
	time.Sleep(2 * time.Second)
	if isListen {
		return false
	}

	resp, err := http.Get(fmt.Sprintf("http://localhost:%d", port))
	if err != nil {
		logs.GetLogger().Errorf("Port %d is not accessible: %v", port, err)
		return false
	}
	defer resp.Body.Close()
	return true

}

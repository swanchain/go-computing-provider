package util

import (
	"context"
	"errors"
	"fmt"
	"github.com/filswan/go-mcs-sdk/mcs/api/common/logs"
	"github.com/valyala/gozstd"
	"io"
	"net/http"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

func SaveMcsFileByUrlToFile(fileName, mcsUrl string) error {
	resp, err := http.Get(mcsUrl)
	if err != nil {
		return fmt.Errorf("error making request to Space API: %+v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("space API response not OK. Status Code: %d", resp.StatusCode)
	}

	outBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read body failed, error:", err)
		return err
	}

	out, err := gozstd.Decompress(nil, outBytes)
	if err != nil {
		fmt.Println("decompress response bytes failed, error:", err)
		return err
	}
	return os.WriteFile(fileName, out, 0655)
}

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
			logs.GetLogger().Errorf("Port %d is closed: %v", port, err)
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

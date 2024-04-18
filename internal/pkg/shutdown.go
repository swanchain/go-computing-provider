package pkg

import (
	"context"
	"errors"
	"github.com/filswan/go-mcs-sdk/mcs/api/common/logs"
	"github.com/swanchain/go-computing-provider/conf"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type StopFunc func(context.Context) error

type ShutdownHandler struct {
	Component string
	StopFunc  StopFunc
}

func MonitorShutdown(triggerCh <-chan struct{}, handlers ...ShutdownHandler) <-chan struct{} {
	sigCh := make(chan os.Signal, 2)
	out := make(chan struct{})

	go func() {
		select {
		case sig := <-sigCh:
			logs.GetLogger().Warn("received shutdown", "signal", sig)
		case <-triggerCh:
			logs.GetLogger().Warn("received shutdown")
		}

		logs.GetLogger().Warn("Shutting down...")

		// Call all the handlers, logging on failure and success.
		for _, h := range handlers {
			if err := h.StopFunc(context.TODO()); err != nil {
				logs.GetLogger().Errorf("shutting down %s failed: %s", h.Component, err)
				continue
			}
			logs.GetLogger().Infof("%s shut down successfully ", h.Component)
		}

		logs.GetLogger().Warn("Graceful shutdown successful")

		close(out)
	}()

	signal.Reset(syscall.SIGTERM, syscall.SIGINT)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)
	return out
}

func ServeHttp(h http.Handler, name string, addr string) (StopFunc, error) {
	// Instantiate the server and start listening.
	srv := &http.Server{
		Addr:              addr,
		Handler:           h,
		ReadHeaderTimeout: 60 * time.Second,
	}

	go func() {
		certFile := conf.GetConfig().LOG.CrtFile
		keyFile := conf.GetConfig().LOG.KeyFile
		if _, err := os.Stat(certFile); err != nil {
			logs.GetLogger().Fatalf("need to manually generate the wss authentication certificate. error: %v", err)
			return
		}

		if err := srv.ListenAndServeTLS(certFile, keyFile); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logs.GetLogger().Fatalf("service: %s, listen: %s\n", name, err)
		}

	}()

	return srv.Shutdown, nil
}

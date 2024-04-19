package pkg

import (
	"context"
	"errors"
	"fmt"
	"github.com/filswan/go-mcs-sdk/mcs/api/common/logs"
	"github.com/gin-gonic/gin"
	"github.com/swanchain/go-computing-provider/conf"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
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

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		fmt.Fprintf(gin.DefaultWriter, "[GIN] %v | %3d | %13v | %15s | %s %s %s \n",
			time.Now().Format("2006-01-02T15:04:05.000Z0700"),
			c.Writer.Status(),
			time.Now().Sub(start),
			GetRealIP(c.Request),
			c.Request.Method,
			c.Request.URL.Path,
			c.Errors.String(),
		)
	}
}

func GetRealIP(r *http.Request) string {
	ipAddress := r.RemoteAddr
	if colonIndex := strings.LastIndex(ipAddress, ":"); colonIndex != -1 {
		ipAddress = ipAddress[:colonIndex]
	}

	if forwardedFor := r.Header.Get("X-Forwarded-For"); forwardedFor != "" {
		ipAddress = strings.TrimSpace(strings.Split(forwardedFor, ",")[0])
	}

	realIP := net.ParseIP(ipAddress)
	if realIP == nil {
		return "not found client ip"
	}
	return realIP.String()
}

package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/swanchain/go-computing-provider/internal/v1/service"
)

func CPApiV1(router *gin.RouterGroup) {
	router.GET("/host/info", service.GetServiceProviderInfo)
	router.POST("/lagrange/jobs", service.ReceiveJob)
	router.POST("/lagrange/jobs/redeploy", service.RedeployJob)
	router.DELETE("/lagrange/jobs", service.CancelJob)
	router.GET("/lagrange/cp", service.StatisticalSources)
	router.POST("/lagrange/jobs/renew", service.ReNewJob)
	router.GET("/lagrange/spaces/log", service.GetSpaceLog)
	router.POST("/lagrange/cp/proof", service.DoProof)

	router.GET("/cp", service.StatisticalSources)
	router.GET("/cp/info", service.GetCpInfo)
	router.POST("/cp/ubi", service.DoUbiTask)
	router.POST("/cp/receive/ubi", service.ReceiveUbiProof)

}

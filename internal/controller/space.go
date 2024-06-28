package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/swanchain/go-computing-provider/internal/v2/services"
)

type SpaceHandler struct {
	jobService services.JobService
}

func NewSpaceHandler(jobService services.JobService) *SpaceHandler {
	return &SpaceHandler{jobService: jobService}
}

func (sh *SpaceHandler) DeployJob(c *gin.Context) {

}

func (sh *SpaceHandler) RedeployJob(c *gin.Context) {

}

func (sh *SpaceHandler) DeleteJob(c *gin.Context) {

}

func (sh *SpaceHandler) extendJob(c *gin.Context) {

}

func (sh *SpaceHandler) GetLog(c *gin.Context) {

}

func (sh *SpaceHandler) WhiteList(c *gin.Context) {

}

func (sh *SpaceHandler) JobStatus(c *gin.Context) {

}

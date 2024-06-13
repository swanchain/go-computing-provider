package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/swanchain/go-computing-provider/internal/service"
)

type SpaceHandler struct {
	spaceService service.SpaceService
}

func NewSpaceHandler(spaceService service.SpaceService) *SpaceHandler {
	return &SpaceHandler{spaceService: spaceService}
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

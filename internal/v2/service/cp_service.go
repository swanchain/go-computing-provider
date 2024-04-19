package service

import (
	"github.com/gin-gonic/gin"
	logging "github.com/ipfs/go-log/v2"
	"github.com/swanchain/go-computing-provider/internal/pkg"
	"github.com/swanchain/go-computing-provider/internal/v2/models"
	"gorm.io/gorm"
	"net/http"
)

var servlog = logging.Logger("service")

type JobService struct {
	*gorm.DB
	k8sClient pkg.K8sService
}

func (js *JobService) doJob(c *gin.Context) {
	var job models.CreateJobReq
	if err := c.ShouldBindJSON(&job); err != nil {
		c.JSON(http.StatusBadRequest, CreateFailedRespWithCode(ReqParseJsonCode))
		return
	}

	servlog.Infof("Job received Data: %+v", job)

}

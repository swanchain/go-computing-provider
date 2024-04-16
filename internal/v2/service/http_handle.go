package service

import (
	"github.com/filswan/go-swan-lib/logs"
	"github.com/gin-gonic/gin"
	"github.com/swanchain/go-computing-provider/internal/v2/models"
	"net/http"
)

func CreateJob(c *gin.Context) {
	var job models.Job
	if err := c.ShouldBindJSON(&job); err != nil {
		c.JSON(http.StatusBadRequest, CreateFailedRespWithCode(ReqParseJsonCode))
		return
	}

	logs.GetLogger().Infof("Job received Data: %+v", job)

	service := NewJobService()
	service.doJob(job)
}

func ExtendJob(c *gin.Context) {

}

func TerminateJob(c *gin.Context) {

}

func GetCpResources(c *gin.Context) {

}

func GetJobInfo(c *gin.Context) {

}

type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg,omitempty"`
	Data any    `json:"data,omitempty"`
}

const (
	SuccessCode       = 200
	ReqParseJsonCode  = 300
	DataParseJsonCode = 301
	BadParamCode      = 400
)

var codeMsg = map[int]string{
	ReqParseJsonCode:  "request params parse to json failed",
	DataParseJsonCode: "data parse to json failed",
}

func CreateSuccessResp(data any) Result {
	return Result{
		Code: SuccessCode,
		Data: data,
	}
}

func CreateFailedRespWitBadParamMsg(msg string) Result {
	return Result{
		Code: BadParamCode,
		Msg:  msg,
	}
}

func CreateFailedRespWithCode(code int) Result {
	return Result{
		Code: code,
		Msg:  codeMsg[code],
	}
}

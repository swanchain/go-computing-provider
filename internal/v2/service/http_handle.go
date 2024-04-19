package service

import (
	"github.com/gin-gonic/gin"
)

func CreateJob(c *gin.Context) {

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

package common

type HttpResult struct {
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

func CreateSuccessResp(data any) HttpResult {
	return HttpResult{
		Code: SuccessCode,
		Data: data,
	}
}

func CreateFailedRespWithCode(code int, errMsg ...string) HttpResult {
	var msg string
	if len(errMsg) > 0 {
		msg = errMsg[0]
	} else {
		msg = codeMsg[code]
	}
	return HttpResult{
		Code: code,
		Msg:  msg,
	}
}

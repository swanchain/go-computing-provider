package pkg

import (
	libconstants "github.com/filswan/go-swan-lib/constants"
)

type BasicResponse struct {
	Status   string      `json:"status"`
	Code     int         `json:"code"`
	Data     interface{} `json:"data,omitempty"`
	Message  string      `json:"message,omitempty"`
	PageInfo *PageInfo   `json:"page_info,omitempty"`
}

type PageInfo struct {
	PageNumber       string `json:"page_number"`
	PageSize         string `json:"page_size"`
	TotalRecordCount string `json:"total_record_count"`
}

type MixedResponse struct {
	BasicResponse
	MixData struct {
		Success interface{} `json:"success"`
		Fail    interface{} `json:"fail"`
	} `json:"mix_data"`
}

func CreateSuccessResponse(_data interface{}) BasicResponse {
	return BasicResponse{
		Status: libconstants.SWAN_API_STATUS_SUCCESS,
		Data:   _data,
		Code:   SuccessCode,
	}
}

func CreateErrorResponse(code int, errMsg ...string) BasicResponse {
	var msg string
	if len(errMsg) == 0 {
		msg = codeMsg[code]
	} else {
		msg = errMsg[0]
	}
	return BasicResponse{
		Status:  libconstants.SWAN_API_STATUS_FAIL,
		Code:    code,
		Message: msg,
	}
}

const (
	SuccessCode = 200
	JsonError   = 400
	ServerError = 500

	SpaceSignatureError     = 6001
	ProofParamError         = 7001
	ProofReadLogError       = 7002
	ProofError              = 7003
	UbiTaskParamError       = 8001
	UbiTaskReadLogError     = 8002
	UbiTaskError            = 8003
	CheckResourcesError     = 9001
	CheckAvailableResources = 9002
)

var codeMsg = map[int]string{
	JsonError:           "An error occurred while converting to json",
	ServerError:         "Service failed",
	SpaceSignatureError: "Verify signature failed",

	ProofReadLogError: "An error occurred while read the log of proof",
	ProofError:        "An error occurred while executing the calculation task",

	CheckResourcesError:     "An error occurred while check resources available",
	CheckAvailableResources: "No resources available",
}

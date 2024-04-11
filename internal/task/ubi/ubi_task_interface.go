package ubi

import "github.com/swanchain/go-computing-provider/internal/models"

type TaskType struct {
}

var ubiTaskMap map[string]UbiTask

type UbiTask interface {
	doTask(models.UBITaskReq)
	ReceiveUbiProof(models.ReceiveProof)
}

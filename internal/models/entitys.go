package models

import (
	"encoding/json"
	"gorm.io/gorm"
)

const (
	FIL_C2_CPU512 = 1
	FIL_C2_CPU32G = 2
	FIL_C2_GPU512 = 3
	FIL_C2_GPU32G = 4
)

func UbiTaskTypeStr(typeInt int) string {
	var typeStr string
	switch typeInt {
	case FIL_C2_CPU512:
		typeStr = "fil-c2-512M"
	case FIL_C2_CPU32G:
		typeStr = "fil-c2-32G"
	case FIL_C2_GPU512:
		typeStr = "fil-c2-512M"
	case FIL_C2_GPU32G:
		typeStr = "fil-c2-32G"
	}
	return typeStr
}

const (
	TASK_RECEIVED_STATUS = iota + 1
	TASK_RUNNING_STATUS
	TASK_SUBMITTED_STATUS
	TASK_FAILED_STATUS
	TASK_VERIFIED_STATUS
	TASK_INVALID_STATUS
	TASK_VERIFYFAILED_STATUS
	TASK_REWARDED_STATUS
	TASK_TIMEOUT_STATUS
	TASK_REPEATED_STATUS
	TASK_NSC_STATUS
	TASK_UNKNOWN_STATUS
)

func TaskStatusStr(status int) string {
	var statusStr string
	switch status {
	case TASK_RECEIVED_STATUS:
		statusStr = "received"
	case TASK_RUNNING_STATUS:
		statusStr = "running"
	case TASK_SUBMITTED_STATUS:
		statusStr = "submitted"
	case TASK_FAILED_STATUS:
		statusStr = "failed"
	case TASK_VERIFIED_STATUS:
		statusStr = "verified"
	case TASK_REWARDED_STATUS:
		statusStr = "rewarded"
	case TASK_INVALID_STATUS:
		statusStr = "invalid"
	case TASK_TIMEOUT_STATUS:
		statusStr = "timeout"
	case TASK_VERIFYFAILED_STATUS:
		statusStr = "verifyFailed"
	case TASK_REPEATED_STATUS:
		statusStr = "repeated"
	case TASK_NSC_STATUS:
		statusStr = "NSC"
	case TASK_UNKNOWN_STATUS:
		statusStr = "unknown"
	}

	return statusStr
}

const (
	RESOURCE_TYPE_CPU = 0
	RESOURCE_TYPE_GPU = 1
)

const (
	TaskSequencer = 1
)

func GetResourceTypeStr(resourceType int) string {
	switch resourceType {
	case RESOURCE_TYPE_CPU:
		return "CPU"
	case RESOURCE_TYPE_GPU:
		return "GPU"
	}
	return ""
}

type TaskEntity struct {
	Id                 int64  `json:"id" gorm:"primaryKey;id"`
	Type               int    `json:"type" gorm:"type"`
	Name               string `json:"name" gorm:"name"`
	Contract           string `json:"contract" gorm:"contract"`
	ResourceType       int    `json:"resource_type" gorm:"resource_type"` // 1
	InputParam         string `json:"input_param" gorm:"input_param"`
	VerifyParam        string `json:"verify_param" gorm:"verify_param"`
	TxHash             string `json:"tx_hash" gorm:"tx_hash"`
	Status             int    `json:"status" gorm:"status"`
	CreateTime         int64  `json:"create_time" gorm:"create_time"`
	EndTime            int64  `json:"end_time" gorm:"end_time"`
	Error              string `json:"error" gorm:"error"`
	Deadline           int64  `json:"deadline"`
	CheckCode          string `json:"check_code"`
	BlockHash          string `json:"block_hash"`
	Sign               string `json:"sign"`
	Reward             string `json:"reward"`
	SequenceCid        string `json:"sequence_cid"`
	SettlementCid      string `json:"settlement_cid"`
	SequenceTaskAddr   string `json:"sequence_task_addr"`
	SettlementTaskAddr string `json:"settlement_task_addr"`
	Sequencer          int    `json:"sequencer" gorm:"default:-1"`
	Proof              string
}

func (task *TaskEntity) TableName() string {
	return "t_task"
}

const (
	All_FLAG         = -1
	UN_DELETEED_FLAG = 0
	DELETED_FLAG     = 1
)

const (
	POD_UNKNOWN_STATUS = 0
	POD_RUNNING_STATUS = 1
	POD_DELETE_STATUS  = 2
)

const (
	JOB_RECEIVED_STATUS   = 0
	JOB_DEPLOY_STATUS     = 1
	JOB_RUNNING_STATUS    = 2
	JOB_TERMINATED_STATUS = 3
	JOB_COMPLETED_STATUS  = 4
)

func GetJobStatus(status int) string {
	var statusStr string
	switch status {
	case JOB_RECEIVED_STATUS:
		statusStr = "received"
	case JOB_DEPLOY_STATUS:
		statusStr = "deploying"
	case JOB_RUNNING_STATUS:
		statusStr = "running"
	case JOB_TERMINATED_STATUS:
		statusStr = "terminated"
	case JOB_COMPLETED_STATUS:
		statusStr = "completed"
	default:
		statusStr = "unknown"
	}
	return statusStr
}

const (
	DEPLOY_RECEIVE_JOB = iota + 1
	DEPLOY_DOWNLOAD_SOURCE
	DEPLOY_UPLOAD_RESULT
	DEPLOY_BUILD_IMAGE
	DEPLOY_PUSH_IMAGE
	DEPLOY_PULL_IMAGE
	DEPLOY_TO_K8S
)

func GetDeployStatusStr(deployStatus int) string {
	var statusStr string
	switch deployStatus {
	case DEPLOY_DOWNLOAD_SOURCE:
		statusStr = "downloadSource"
	case DEPLOY_UPLOAD_RESULT:
		statusStr = "uploadResult"
	case DEPLOY_BUILD_IMAGE:
		statusStr = "buildImage"
	case DEPLOY_PUSH_IMAGE:
		statusStr = "pushImage"
	case DEPLOY_PULL_IMAGE:
		statusStr = "pullImage"
	case DEPLOY_TO_K8S:
		statusStr = "deployToK8s"
	}
	return statusStr
}

type JobEntity struct {
	Id              int64  `json:"id" gorm:"primaryKey;autoIncrement"`
	Source          string `json:"source" gorm:"source"` // market name
	Name            string `json:"name" gorm:"name"`
	SpaceUuid       string `json:"space_uuid" gorm:"space_uuid"`
	JobUuid         string `json:"job_uuid" gorm:"job_uuid"`
	TaskUuid        string `json:"task_uuid" gorm:"task_uuid"`
	ResourceType    string `json:"resource_type"  gorm:"resource_type"`
	SpaceType       int    `json:"space_type" gorm:"space_type"` // 0: public; 1: private
	SourceUrl       string `json:"source_url" gorm:"source_url"`
	Hardware        string `json:"hardware" gorm:"hardware"`
	Duration        int    `json:"duration" gorm:"duration"`
	DeployStatus    int    `json:"deploy_status" gorm:"deploy_status"`
	WalletAddress   string `json:"wallet_address"  gorm:"wallet_address"`
	ResultUrl       string `json:"result_url" gorm:"result_url"`
	RealUrl         string `json:"real_url" gorm:"real_url"`
	K8sDeployName   string `json:"k8s_deploy_name" gorm:"k8s_deploy_name"`
	K8sResourceType string `json:"k8s_resource_type" gorm:"k8s_resource_type"`
	NameSpace       string `json:"name_space" gorm:"name_space"`
	ImageName       string `json:"image_name" gorm:"image_name"`
	BuildLog        string `json:"build_log" gorm:"build_log"`
	ContainerLog    string `json:"container_log" gorm:"container_log"`
	Reward          string `json:"reward" gorm:"reward"`
	ExpireTime      int64  `json:"expire_time" gorm:"expire_time"`
	CreateTime      int64  `json:"create_time" gorm:"create_time"`
	Error           string `json:"error" gorm:"error"`
	DeleteAt        int    `json:"delete_at" gorm:"delete_at; default:0"` // 1 deleted
	IpWhiteList     string `json:"ip_white_list"`
	PodStatus       int    `json:"pod_status"`
	Status          int    `json:"status"`
	StartedBlock    uint64 `json:"started_block" gorm:"column:started_block;not null;default:0"`
	ScannedBlock    uint64 `json:"scanned_block" gorm:"column:scanned_block;not null;default:0"`
	EndedBlock      uint64 `json:"ended_block" gorm:"column:ended_block;not null;default:0"`
}

func (*JobEntity) TableName() string {
	return "t_job"
}

const (
	Task_TYPE_FIL_C2_512 = iota + 1
	Task_TYPE_MINING
	Task_TYPE_AI
	Task_TYPE_FIL_C2_32
	Task_TYPE_NODE_PORT
)

func TaskTypeStr(taskType int) string {
	var typeStr string
	switch taskType {
	case Task_TYPE_FIL_C2_512:
		typeStr = "Fil-C2-512M"
	case Task_TYPE_MINING:
		typeStr = "Mining"
	case Task_TYPE_AI:
		typeStr = "AI"
	case Task_TYPE_FIL_C2_32:
		typeStr = "Fil-C2-32G"
	case Task_TYPE_NODE_PORT:
		typeStr = "NodePort"
	}
	return typeStr
}

type CpInfoEntity struct {
	Id                 int64    `json:"id" gorm:"primaryKey;autoIncrement"`
	NodeId             string   `json:"node_id" gorm:"node_id"`
	OwnerAddress       string   `json:"owner_address" gorm:"owner_address"`
	Beneficiary        string   `json:"beneficiary" gorm:"beneficiary"`
	WorkerAddress      string   `json:"worker_address" gorm:"worker_address"`
	Version            string   `json:"version" gorm:"version"`
	ContractAddress    string   `json:"contract_address" gorm:"contract_address"`
	MultiAddressesJSON string   `gorm:"multi_addresses_json;type:text" json:"-"`
	TaskTypesJSON      string   `gorm:"task_types_json; type:text" json:"-"`
	CreateAt           string   `json:"create_at" gorm:"create_at"`
	UpdateAt           string   `json:"update_at" gorm:"update_at"`
	MultiAddresses     []string `json:"multi_addresses" gorm:"-"`
	TaskTypes          []uint8  `json:"task_types" gorm:"-"` // 1:Fil-C2-512M, 2:mining, 3: AI, 4:Fil-C2-32G

}

func (*CpInfoEntity) TableName() string {
	return "t_cp_info"
}

func (c *CpInfoEntity) BeforeSave(tx *gorm.DB) (err error) {
	if len(c.MultiAddresses) != 0 {
		if multiAddrBytes, err := json.Marshal(c.MultiAddresses); err == nil {
			c.MultiAddressesJSON = string(multiAddrBytes)
		} else {
			return err
		}
	}

	if len(c.TaskTypes) != 0 {
		intTaskTypes := make([]int, len(c.TaskTypes))
		for i, v := range c.TaskTypes {
			intTaskTypes[i] = int(v)
		}

		if taskTypesBytes, err := json.Marshal(intTaskTypes); err == nil {
			c.TaskTypesJSON = string(taskTypesBytes)
		} else {
			return err
		}
	}
	return nil
}

func (c *CpInfoEntity) AfterFind(tx *gorm.DB) (err error) {
	if err = json.Unmarshal([]byte(c.MultiAddressesJSON), &c.MultiAddresses); err != nil {
		return err
	}
	if err = json.Unmarshal([]byte(c.TaskTypesJSON), &c.TaskTypes); err != nil {
		return err
	}
	return nil
}

const (
	NetworkNetset               = "netset-2300e518e9ad45"
	NetworkGlobalSubnet         = "global-e59cad59af9c65"
	NetworkGlobalOutAccess      = "global-pd6sdo8cjd61yd"
	NetworkGlobalInAccess       = "global-01kls78xh7dk4n"
	NetworkGlobalNamespace      = "global-ao9kq72mjc0sl3"
	NetworkGlobalDns            = "global-s92ms87dl3j6do"
	NetworkGlobalPodInNamespace = "global-pod1namespace1"
)

var networkPolicyMap = []string{NetworkGlobalSubnet, NetworkGlobalOutAccess,
	NetworkGlobalInAccess, NetworkGlobalNamespace, NetworkGlobalDns}

func ExistResource(name string) bool {
	for _, s := range networkPolicyMap {
		if s == name {
			return true
		}
	}
	return false
}

type EcpJobEntity struct {
	Id            int64  `json:"id" gorm:"primaryKey;autoIncrement"`
	Uuid          string `json:"uuid" gorm:"uuid"`
	Name          string `json:"name" gorm:"name"`
	Image         string `json:"image" gorm:"image"`
	Env           string `json:"env" gorm:"env"`
	Status        string `json:"status"` // created|restarting|running|removing|paused|exited|dead
	ContainerName string `json:"container_name" gorm:"container_name"`
	CreateTime    int64  `json:"create_time" gorm:"create_time"`
	DeleteAt      int    `json:"delete_at" gorm:"delete_at; default:0"` // 1 deleted
}

func (*EcpJobEntity) TableName() string {
	return "t_ecp_job"
}

package models

import (
	"math/big"
	"time"
)

type BidStatus string

const (
	BidDisabledStatus    BidStatus = "bidding_disabled"
	BidEnabledStatus     BidStatus = "bidding_enabled"
	BidGpuDisabledStatus BidStatus = "bidding_gpu_disabled"

	ActiveStatus   string = "Active"
	InactiveStatus string = "Inactive"
)

type ComputingProvider struct {
	Name          string `json:"name"`
	NodeId        string `json:"node_id"`
	MultiAddress  string `json:"multi_address"`
	Autobid       int    `json:"autobid"`
	Status        string `json:"status"`
	PublicAddress string `json:"public_address"`
}

type JobData struct {
	UUID     string `json:"uuid"`
	Name     string `json:"name"`
	Status   string `json:"status"`
	Duration int    `json:"duration"`
	//Hardware      string `json:"hardware"`
	JobSourceURI                string `json:"job_source_uri"`
	JobResultURI                string `json:"job_result_uri"`
	StorageSource               string `json:"storage_source"`
	TaskUUID                    string `json:"task_uuid"`
	CreatedAt                   string `json:"created_at"`
	UpdatedAt                   string `json:"updated_at"`
	BuildLog                    string `json:"build_log"`
	ContainerLog                string `json:"container_log"`
	NodeIdJobSourceUriSignature string `json:"node_id_job_source_uri_signature"`
}

type Job struct {
	Uuid   string
	Status JobStatus
	Url    string
	Count  int
}

type JobStatus string

const (
	JobDownloadSource JobStatus = "downloadSource" // download file form job_resource_uri
	JobUploadResult   JobStatus = "uploadResult"   // upload task result to mcs
	JobBuildImage     JobStatus = "buildImage"     // build images
	JobPushImage      JobStatus = "pushImage"      // push image to registry
	JobPullImage      JobStatus = "pullImage"      // download file form job_resource_uri
	JobDeployToK8s    JobStatus = "deployToK8s"    // deploy image to k8s
)

type DeleteJobReq struct {
	CreatorWallet string `json:"creator_wallet"`
	SpaceName     string `json:"space_name"`
}

type SpaceJSON struct {
	Data struct {
		Files []SpaceFile `json:"files"`
		Owner struct {
			PublicAddress string `json:"public_address"`
		} `json:"owner"`
		Space struct {
			Uuid        string `json:"uuid"`
			Name        string `json:"name"`
			ActiveOrder struct {
				Config SpaceHardware `json:"config"`
			} `json:"activeOrder"`
		} `json:"space"`
	} `json:"data"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

type SpaceFile struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type SpaceHardware struct {
	Description  string `json:"description"`
	HardwareType string `json:"hardware_type"`
	Memory       int    `json:"memory"`
	Name         string `json:"name"`
	Vcpu         int    `json:"vcpu"`
}

type Resource struct {
	Cpu     Specification
	Memory  Specification
	Gpu     Specification
	Storage Specification
}

type Specification struct {
	Quantity int64
	Unit     string
}

type CacheSpaceDetail struct {
	WalletAddress string
	SpaceName     string
	SpaceUuid     string
	ExpireTime    int64
	JobUuid       string
	TaskType      string
	DeployName    string
	Hardware      string
	Url           string
	TaskUuid      string
}

type UBITaskReq struct {
	ID         int           `json:"id"`
	Name       string        `json:"name,omitempty"`
	Type       int           `json:"type"`
	ZkType     string        `json:"zk_type"`
	InputParam string        `json:"input_param"`
	Signature  string        `json:"signature"`
	Resource   *TaskResource `json:"resource"`
}

type TaskResource struct {
	CPU     string `json:"cpu"`
	GPU     string `json:"gpu"`
	Memory  string `json:"memory"`
	Storage string `json:"storage"`
}

type CacheUbiTaskDetail struct {
	TaskId     string `json:"task_id"`
	TaskType   string `json:"task_type"`
	ZkType     string `json:"zk_type"`
	Tx         string `json:"tx"`
	Status     string `json:"status"`
	Reward     string `json:"reward"`
	CreateTime string `json:"create_time"`
}

type Account struct {
	OwnerAddress   string
	NodeId         string
	MultiAddresses []string
	UbiFlag        uint8
	Beneficiary    struct {
		BeneficiaryAddress string
		Quota              *big.Int
		Expiration         *big.Int
	}
}

type TaskList []CacheUbiTaskDetail

func (t TaskList) Len() int {
	return len(t)
}

func (t TaskList) Less(i, j int) bool {
	timeI, _ := time.Parse("2006-01-02 15:04:05", t[i].CreateTime)
	timeJ, _ := time.Parse("2006-01-02 15:04:05", t[j].CreateTime)
	return timeI.Before(timeJ)
}

func (t TaskList) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

type ReceiveProof struct {
	TaskId    string `json:"task_id"`
	TaskType  string `json:"task_type"`
	Proof     string `json:"proof"`
	ZkType    string `json:"zk_type"`
	NameSpace string `json:"name_space"`
}
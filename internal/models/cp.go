package models

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
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
	UUID                        string `json:"uuid"`
	Name                        string `json:"name"`
	Duration                    int    `json:"duration"`
	JobSourceURI                string `json:"job_source_uri"`
	JobResultURI                string `json:"job_result_uri,omitempty"`
	StorageSource               string `json:"storage_source,omitempty"`
	TaskUUID                    string `json:"task_uuid"`
	BuildLog                    string `json:"build_log,omitempty"`
	ContainerLog                string `json:"container_log"`
	NodeIdJobSourceUriSignature string `json:"node_id_job_source_uri_signature,omitempty"`
	JobRealUri                  string `json:"job_real_uri,omitempty"`
}

type Job struct {
	Uuid   string
	Status int
	Url    string
}

type JobStatus string

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

type UBITaskReq struct {
	ID           int           `json:"id"`
	Name         string        `json:"name,omitempty"`
	Type         int           `json:"type"`
	InputParam   string        `json:"input_param"`
	VerifyParam  string        `json:"verify_param"`
	Signature    string        `json:"signature"`
	Resource     *TaskResource `json:"resource"`
	ResourceType int           `json:"resource_type"`
	DeadLine     int64         `json:"deadline"`
	CheckCode    string        `json:"check_code"`
}

type UbiC2Proof struct {
	TaskId    string `json:"task_id,omitempty"`
	TaskType  string `json:"task_type,omitempty"`
	Proof     string `json:"proof,omitempty"`
	ZkType    string `json:"zk_type,omitempty"`
	NameSpace string `json:"name_space,omitempty"`
}

type TaskResource struct {
	CPU     string `json:"cpu"`
	GPU     string `json:"gpu"`
	Memory  string `json:"memory"`
	Storage string `json:"storage"`
}

type Account struct {
	OwnerAddress   string
	NodeId         string
	MultiAddresses []string
	TaskTypes      []uint8 // 1:Fil-C2-512M, 2:Aleo, 3: AI, 4:Fil-C2-32G
	Beneficiary    string
	WorkerAddress  string
	Version        string
	Contract       string
}

type EcpCollateralInfo struct {
	CpAddress         string
	CollateralBalance string
	FrozenBalance     string
	Status            string
}

type FcpCollateralInfo struct {
	CpAddress        string
	AvailableBalance string
	LockedCollateral string
	Status           string
}

type EcpTaskInfo struct {
	TaskID               *big.Int
	TaskType             *big.Int
	ResourceType         *big.Int
	InputParam           string
	VerifyParam          string
	CpAccount            common.Address
	Proof                string
	Deadline             *big.Int
	TaskRegistryContract common.Address
	CheckCode            string
	Owner                common.Address
	Version              string
}

package models

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

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

type SpaceJsonWithNoData struct {
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
}

type SpaceJSON struct {
	Data    SpaceJsonWithNoData `json:"data"`
	Message string              `json:"message"`
	Status  string              `json:"status"`
}

type SpaceFile struct {
	Name         string `json:"name"`
	URL          string `json:"url"`
	SymmetricKey string `json:"symmetric_key"`
	Iv           string `json:"iv"`
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

type CollateralContractInfoForECP struct {
	CollateralToken string
	WithdrawDelay   int64
}

type CpCollateralInfoForECP struct {
	CpAddress         string
	CollateralBalance string
	FrozenBalance     string
	Status            string
}

type CpCollateralInfoForFCP struct {
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

type WithdrawRequest struct {
	Amount        string
	RequestBlock  int64
	WithdrawDelay int64
}

type TaskInfoOnChain struct {
	TaskUuid           string
	CpList             []string
	OwnerAddress       string
	Reward             *big.Int
	Collateral         *big.Int
	StartTimestamp     int64
	TerminateTimestamp int64
	Duration           int64
	TaskStatus         int
	CollateralStatus   int
}

type ResourcePrice struct {
	CpuPrice         float64            `json:"cpu_price"`
	MemoryPrice      float64            `json:"memory_price"`
	HdEphemeralPrice float64            `json:"hd_ephemeral_price"`
	HdPersHddPrice   float64            `json:"hd_pers_hdd_price"`
	HdPersSsdPrice   float64            `json:"hd_pers_ssd_price"`
	HdPersNvmePrice  float64            `json:"hd_pers_nvme_price"`
	GpuDefaultPrice  float64            `json:"gpu_default_price"`
	GpusPrice        map[string]float64 `json:"gpus_price"`
}

const (
	NOT_ASSIGNED = iota
	IN_PROGRESS
	COMPLETED
	TERMINATED
)

func JobOnChainStatus(status int) string {
	var str string
	switch status {
	case NOT_ASSIGNED:
		str = "not_assigned"
	case IN_PROGRESS:
		str = "in_progress"
	case COMPLETED:
		str = "completed"
	case TERMINATED:
		str = "terminated"
	}

	return str
}

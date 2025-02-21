package models

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type JobData struct {
	UUID                        string   `json:"uuid"`
	Name                        string   `json:"name"`
	Duration                    int      `json:"duration"`
	JobSourceURI                string   `json:"job_source_uri"`
	JobResultURI                string   `json:"job_result_uri,omitempty"`
	StorageSource               string   `json:"storage_source,omitempty"`
	TaskUUID                    string   `json:"task_uuid"`
	BuildLog                    string   `json:"build_log,omitempty"`
	ContainerLog                string   `json:"container_log"`
	NodeIdJobSourceUriSignature string   `json:"node_id_job_source_uri_signature,omitempty"`
	JobRealUri                  string   `json:"job_real_uri,omitempty"`
	JobType                     int      `json:"job_type"`  // 0: Standard job; 1: Custom job
	BidPrice                    string   `json:"bid_price"` // Amount users are willing to pay
	IpWhiteList                 []string `json:"ip_white_list"`
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
	Hardware     string `json:"hardware"` // Nvidia 3080  CPU only
	Memory       int64  `json:"memory"`   // unit bytes
	Name         string `json:"name"`
	Vcpu         int64  `json:"vcpu"`
	Storage      int64  `json:"storage"` // unit bytes
	Gpu          int64  `json:"gpu"`
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
	CPU       string `json:"cpu"`
	Gpu       string `json:"gpu"`
	GpuDetail struct {
		ProductName string `json:"product_name"`
		Count       string `json:"count"`
	} `json:"gpu_detail"`
	Memory  string `json:"memory"`
	Storage string `json:"storage"`
}

type Account struct {
	OwnerAddress   string
	NodeId         string
	MultiAddresses []string
	TaskTypes      []uint8 // 1:Fil-C2-512M, 2:mining, 3: AI, 4:Fil-C2-32G 5:NodePort 100:Exit
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
	CpuPrice         string            `json:"cpu_price"`
	MemoryPrice      string            `json:"memory_price"`
	HdEphemeralPrice string            `json:"hd_ephemeral_price"`
	HdPersHddPrice   string            `json:"hd_pers_hdd_price,omitempty"`
	HdPersSsdPrice   string            `json:"hd_pers_ssd_price,omitempty"`
	HdPersNvmePrice  string            `json:"hd_pers_nvme_price,omitempty"`
	GpuDefaultPrice  string            `json:"gpu_default_price"`
	GpusPrice        map[string]string `json:"gpus_price"`
	Pricing          bool              `json:"pricing"`
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

type FcpDeployImageReq struct {
	Uuid          string              `json:"uuid"`
	Name          string              `json:"name"`
	InstanceName  string              `json:"instance_name"`
	Duration      int                 `json:"duration"`
	Sign          string              `json:"sign"`
	WalletAddress string              `json:"wallet_address"`
	IpWhitelist   []string            `json:"ip_whitelist"`
	DeployConfig  DeployConfig        `json:"deploy_config"`
	Resource      K8sResourceForImage `json:"resource"`
	IpWhiteList   []string            `json:"ip_white_list"`
	BidPrice      string              `json:"bid_price"` // Amount users are willing to pay
	JobType       int                 `json:"job_type"`
	DeployType    int                 `json:"deploy_type"` // 0: field; 1: docker; 2: yaml
	DeployContent string              `json:"deploy_content"`
	HealthPath    string              `json:"health_path"` // deploy_type=1 or 2, used
}

type YamlContent struct {
	Version  string `json:"version"`
	Services struct {
		Image       string   `json:"image"`
		Cmd         []string `json:"cmd"`
		RunCommands []string `json:"run_commands"`
		Envs        []string `json:"envs"`
		ExposePort  []int    `json:"expose_port"`
	} `json:"services"`
}

type DeployConfig struct {
	Image       string            `json:"image"`
	HealthPath  string            `json:"health_path"`
	Envs        map[string]string `json:"envs"`
	Cmd         []string          `json:"cmd"`
	RunCommands []string          `json:"run_commands"`
	Ports       map[string][]int  `json:"ports"`
	WorkDir     string            `json:"work_dir"`
}

type K8sResourceForImage struct {
	Cpu       int64
	Memory    float64
	Storage   float64
	Gpus      []ReqGpu `json:"gpus"`
	FilC2Flag string
}

type ReqGpu struct {
	GpuModel string `json:"gpu_model"`
	GPU      int    `json:"count"`
}

type FcpDeployImageResp struct {
	UUID               string    `json:"uuid,omitempty"`
	ServiceUrl         string    `json:"service_url,omitempty"`
	HealthPath         string    `json:"health_path,omitempty"`
	Price              float64   `json:"price"`
	ServicePortMapping []PortMap `json:"service_port_mapping,omitempty"`
	ContainerLog       string    `json:"container_log"`
	BuildLog           string    `json:"build_log"`
}

// =========

type ZkTaskReq struct {
	Id          int           `json:"id"`
	Uuid        string        `json:"uuid,omitempty"`
	Name        string        `json:"name,omitempty"`
	TaskType    int           `json:"task_type"` // 1: fil-c2-512-cpu; 2:fil-c2-32-cpu; 3: fil-c2-512-gpu; 4: fil-c2-32-gpu; 5: mining
	InputParam  string        `json:"input_param"`
	VerifyParam string        `json:"verify_param"`
	Signature   string        `json:"signature"`
	Resource    *ResourceInfo `json:"resource"`
	DeadLine    int64         `json:"deadline"`
	CheckCode   string        `json:"check_code"`

	Image string            `json:"image,omitempty"`
	Cmd   []string          `json:"cmd,omitempty"`
	Ports map[string][]int  `json:"ports,omitempty"`
	Envs  map[string]string `json:"envs,omitempty"`
}

type ResourceInfo struct {
	CPU     int64 `json:"cpu"`
	Memory  int64 `json:"memory"`  // bytes
	Storage int64 `json:"storage"` // bytes
	Gpus    []struct {
		GPU      int    `json:"gpu"`
		GPUModel string `json:"gpu_model"`
	} `json:"gpus"`
}

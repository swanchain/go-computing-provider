package models

type ClusterResource struct {
	NodeId           string          `json:"node_id,omitempty"`
	CpAccountAddress string          `json:"cpAccount_address"`
	Region           string          `json:"region,omitempty"`
	ClusterInfo      []*NodeResource `json:"cluster_info"`
	NodeName         string          `json:"node_name,omitempty"`
	Runtime          string          `json:"runtime,omitempty"`
}

type NodeResource struct {
	MachineId string `json:"machine_id"`
	CpuName   string `json:"cpu_name"`
	Cpu       Common `json:"cpu"`
	Vcpu      Common `json:"vcpu"`
	Memory    Common `json:"memory"`
	Gpu       Gpu    `json:"gpu"`
	Storage   Common `json:"storage"`
}

type CollectNodeInfo struct {
	Gpu     Gpu    `json:"gpu"`
	CpuName string `json:"cpu_name"`
}
type Gpu struct {
	DriverVersion string      `json:"driver_version"`
	CudaVersion   string      `json:"cuda_version"`
	AttachedGpus  int         `json:"attached_gpus"`
	Details       []GpuDetail `json:"details"`
}

type GpuDetail struct {
	ProductName   string    `json:"product_name"`
	Status        GpuStatus `json:"status"`
	FbMemoryUsage Common    `json:"fb_memory_usage"`
	OriginalName  string    `json:"original_name"`
	Index         string    `json:"index"`
}

type Common struct {
	Total        string `json:"total"`
	Used         string `json:"used"`
	Free         string `json:"free"`
	RemainderNum int64  `json:"-"`
}

type ResourceStatus struct {
	Request  int64
	Capacity int64
}

type GpuStatus string

const (
	Occupied  GpuStatus = "occupied"
	Available GpuStatus = "available"
)

type ResourcePolicy struct {
	Cpu     CpuQuota   `json:"cpu"`
	Gpu     []GpuQuota `json:"gpu"`
	Memory  Quota      `json:"memory"`
	Storage Quota      `json:"storage"`
}

type CpuQuota struct {
	Quota int64 `json:"quota"`
}

type GpuQuota struct {
	Name  string `json:"name"`
	Quota int64  `json:"quota"`
}
type Quota struct {
	Quota int64  `json:"quota"`
	Unit  string `json:"unit"`
}

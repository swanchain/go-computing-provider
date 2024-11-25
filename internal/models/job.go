package models

type EcpImageJobReq struct {
	UUID       string            `json:"uuid,omitempty"`
	Name       string            `json:"name,omitempty"`
	Image      string            `json:"image,omitempty"`
	Cmd        []string          `json:"cmd"`
	Ports      []int             `json:"ports"`
	HealthPath string            `json:"health_path"`
	Envs       map[string]string `json:"envs,omitempty"`
	Resource   HardwareResource  `json:"resource"`
	Price      string            `json:"price"`
	Duration   int               `json:"duration"`
	JobType    int               `json:"job_type"` // 1 mining; 2: inference
}

type EcpImageResp struct {
	UUID               string    `json:"uuid,omitempty"`
	ServiceUrl         string    `json:"service_url,omitempty"`
	HealthPath         string    `json:"health_path,omitempty"`
	Price              float64   `json:"price"`
	ServicePortMapping []PortMap `json:"service_port_mapping,omitempty"`
}

type PortMap struct {
	ContainerPort int `json:"container_port"`
	ExternalPort  int `json:"external_port"`
}

type HardwareResource struct {
	CPU      int64  `json:"cpu"`
	Memory   int64  `json:"memory"`
	Storage  int64  `json:"storage"`
	GPU      int    `json:"gpu"`
	GPUModel string `json:"gpu_model"`
}

type EcpJobStatusResp struct {
	Uuid    string `json:"uuid"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type EcpInferenceReq struct {
	UUID       string            `json:"uuid,omitempty"`
	Name       string            `json:"name,omitempty"`
	Image      string            `json:"image,omitempty"`
	Cmd        []string          `json:"cmd"`
	Ports      []int             `json:"ports"`
	HealthPath string            `json:"health_path"`
	Envs       map[string]string `json:"envs,omitempty"`
	Resource   HardwareResource  `json:"resource"`
	Price      string            `json:"price"`
	Duration   int               `json:"duration"`
}

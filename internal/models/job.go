package models

type EcpJobCreateReq struct {
	UUID     string            `json:"uuid,omitempty"`
	Name     string            `json:"name,omitempty"`
	Image    string            `json:"image,omitempty"`
	Envs     map[string]string `json:"envs,omitempty"`
	Resource HardwareResource  `json:"resource"`
	Price    string            `json:"price"`
	Duration int               `json:"duration"`
}

type HardwareResource struct {
	CPU      int64  `json:"cpu"`
	Memory   int64  `json:"memory"`
	Storage  int64  `json:"storage"`
	GPU      int64  `json:"gpu"`
	GPUModel string `json:"gpu_model"`
}

type EcpJobStatusResp struct {
	Uuid   string `json:"uuid"`
	Status string `json:"status"`
}

package models

type EcpJobCreateReq struct {
	UUID     string            `json:"uuid"`
	Name     string            `json:"name"`
	Image    string            `json:"image"`
	Envs     map[string]string `json:"envs"`
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

type EcpJobDeleteReq struct {
	Uuid string `json:"uuid"`
}

package models

import "github.com/docker/docker/api/types/container"

type EcpImageJobReq struct {
	Uuid             string            `json:"uuid"`
	Name             string            `json:"name"`
	Price            string            `json:"price"`
	Duration         int               `json:"duration"`
	JobType          int               `json:"job_type"`
	HealthPath       string            `json:"health_path"`
	Sign             string            `json:"sign"`
	WalletAddress    string            `json:"wallet_address"`
	Resource         *HardwareResource `json:"resource"`
	YamlConfig       *YamlConfig       `json:"yaml_config"`
	DockerfileConfig *DockerfileConfig `json:"dockerfile_config"`
	EnvsForDb        []string
}

type HardwareResource struct {
	CPU      int64  `json:"cpu"`
	Memory   int64  `json:"memory"`
	Storage  int64  `json:"storage"`
	GPU      int    `json:"gpu"`
	GPUModel string `json:"gpu_model"`
}

type YamlConfig struct {
	Image string            `json:"image"`
	Cmd   []string          `json:"cmd"`
	Ports []int             `json:"ports"`
	Envs  map[string]string `json:"envs"`
}

type DockerfileConfig struct {
	BaseImage   string            `json:"base_image"`
	WorkDir     string            `json:"work_dir"`
	EnvVars     map[string]string `json:"env_vars"`
	RunCommands []string          `json:"run_commands"`
	ExposePorts []int             `json:"expose_ports"`
	StartCmd    []string          `json:"start_cmd"`
}

type DeployJobParam struct {
	JobType      int // 1 mining; 2: inference
	Uuid         string
	Name         string
	Image        string
	Cmd          []string
	Ports        []int
	HealthPath   string
	Envs         []string
	NeedResource container.Resources

	BuildImagePath string
	BuildImageName string
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

type EcpJobStatusResp struct {
	Uuid               string    `json:"uuid"`
	Status             string    `json:"status"`
	ServiceUrl         string    `json:"service_url,omitempty"`
	HealthPath         string    `json:"health_path,omitempty"`
	Price              float64   `json:"price"`
	ServicePortMapping []PortMap `json:"service_port_mapping,omitempty"`
	Message            string    `json:"message,omitempty"`
}

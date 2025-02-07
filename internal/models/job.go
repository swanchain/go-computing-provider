package models

import "github.com/docker/docker/api/types/container"

type EcpImageJobReq struct {
	Uuid          string            `json:"uuid,omitempty"`
	Name          string            `json:"name,omitempty"`
	Price         string            `json:"price,omitempty"`
	Duration      int               `json:"duration,omitempty"`
	JobType       int               `json:"job_type,omitempty"`
	HealthPath    string            `json:"health_path,omitempty"`
	Sign          string            `json:"sign,omitempty"`
	WalletAddress string            `json:"wallet_address,omitempty"`
	Resource      *HardwareResource `json:"resource,omitempty"`
	Image         string            `json:"image,omitempty"`
	Cmd           []string          `json:"cmd,omitempty"`
	Ports         map[string][]int  `json:"ports,omitempty"`
	Envs          map[string]string `json:"envs,omitempty"`
	RunCommands   []string          `json:"run_commands,omitempty"`
	ResourceUrl   string            `json:"resource_url,omitempty"`
	WorkDir       string            `json:"work_dir,omitempty"`
	IpWhiteList   []string          `json:"ip_white_list"`
	DeployType    int               `json:"deploy_type"` // 0: field; 1: dockerfile; 2: yaml
	DockerContent string            `json:"docker_content"`
	YamlContent   string            `json:"yaml_content"`
}

type HardwareResource struct {
	CPU      int64  `json:"cpu"`
	Memory   int64  `json:"memory"`
	Storage  int64  `json:"storage"`
	GPU      int    `json:"gpu"`
	GPUModel string `json:"gpu_model"`
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
	IpWhiteList    []string `json:"ip_white_list"`
}

type EcpImageResp struct {
	UUID               string    `json:"uuid,omitempty"`
	ServiceUrl         string    `json:"service_url,omitempty"`
	HealthPath         string    `json:"health_path,omitempty"`
	Price              float64   `json:"price"`
	ServicePortMapping []PortMap `json:"service_port_mapping,omitempty"`
	LogUrl             string    `json:"log_url"`
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

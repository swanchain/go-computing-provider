package models

type EcpJobCreateReq struct {
	Uuid     string            `json:"uuid"`
	Name     string            `json:"name"`
	Image    string            `json:"image"`
	EnvVar   map[string]string `json:"env_var"`
	Duration int               `json:"duration"`
	//Resource *TaskResource     `json:"resource"`
}

type EcpJobDeleteReq struct {
	Uuid string `json:"uuid"`
}

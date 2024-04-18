package models

type JobEntity struct {
	Id              uint   `json:"id" gorm:"primaryKey;id"`
	Source          string `json:"source" gorm:"source"`
	Name            string `json:"name" gorm:"name"`
	Type            int    `json:"type" gorm:"type"`
	SourceUrl       string `json:"source_url" gorm:"source_url"`
	Hardware        string `json:"hardware" gorm:"hardware"`
	DeployStatus    int    `json:"deploy_status" gorm:"deploy_status"`
	Status          int    `json:"status" gorm:"status"`
	ResultUrl       string `json:"result_url" gorm:"source_url"`
	K8sDeployName   string `json:"k8s_deploy_name" gorm:"k8s_deploy_name"`
	K8sResourceType string `json:"k8s_resource_type" gorm:"k8s_resource_type"`
	NameSpace       string `json:"name_space" gorm:"name_space"`
	ImageName       string `json:"image_name" gorm:"image_name"`
	BuildLog        string `json:"build_log" gorm:"build_log"`
	ContainerLog    string `json:"container_log" gorm:"container_log"`
	ExpireTime      int64  `json:"expire_time" gorm:"expire_time"`
	CreateTime      int64  `json:"create_time" gorm:"create_time"`
	Error           string `json:"error" gorm:"error"`
}

func (*JobEntity) TableName() string {
	return "Job"
}

package models

const (
	ZK_SOURCE_TYPE_CPU = 0
	ZK_SOURCE_TYPE_GPU = 1
)

const (
	ZK_TYPE_FIL_32  = "fil-c2-32G"
	ZK_TYPE_FIL_512 = "fil-c2-512M"
	ZK_TYPE_ALEO    = "aleo_proof"
)

func GetSourceTypeStr(resourceType int) string {
	switch resourceType {
	case ZK_SOURCE_TYPE_CPU:
		return "CPU"
	case ZK_SOURCE_TYPE_GPU:
		return "GPU"
	}
	return ""
}

type ZkTaskEntity struct {
	Id           uint    `json:"id" gorm:"primaryKey;id"`
	ZkType       string  `json:"zk_type" gorm:"zk_type"`
	Name         string  `json:"name" gorm:"name"`
	ResourceType int     `json:"resource_type" gorm:"resource_type"`
	InputParam   string  `json:"input_param" gorm:"input_param"`
	TxHash       string  `json:"tx_hash" gorm:"tx_hash"`
	Status       int     `json:"status" gorm:"status"`
	Reward       float64 `json:"reward" gorm:"reward"`
	CreateTime   int64   `json:"create_time" gorm:"create_time"`
	Error        string  `json:"error" gorm:"error"`
}

func (zk *ZkTaskEntity) TableName() string {
	return "Task"
}

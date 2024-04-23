package models

type CreateJobReq struct {
	UUID       string  `json:"uuid" binding:"required"`
	Name       string  `json:"name" binding:"required"`
	Duration   int     `json:"duration" binding:"required"`
	SourceURI  string  `json:"source_uri" binding:"required"`
	Entrypoint string  `json:"entrypoint"`
	Source     string  `json:"source" binding:"required"`
	Config     *Config `json:"config" binding:"required"`
	Sign       string  `json:"sign" binding:"required"`
	Payment    Payment `json:"pyment" binding:"required"`
}

type Config struct {
	Vcpu        int    `json:"vcpu"`
	Memory      int64  `json:"memory"`  //Byte
	Storage     int64  `json:"storage"` //Byte
	StorageType string `json:"storage_type"`
	GPU         int    `json:"gpu"`
	GPUModel    string `json:"gpu_model"`
}

type Payment struct {
	TxHash  string `json:"tx_hash"`
	ChainId int64  `json:"chain_id"`
}

package models

type Commit1Task struct {
	SectorNum  int64      `json:"SectorNum"`
	Phase1Out  string     `json:"Phase1Out"`
	SectorSize uint64     `json:"SectorSize"`
	Sid        SectorRef  `json:"Sid"`
	Ticket     string     `json:"Ticket"`
	Cids       SectorCids `json:"Cids"`
	Seed       Seed       `json:"seed"`
}

type SectorRef struct {
	ID struct {
		Miner  uint64 `json:"Miner"`
		Number uint64 `json:"Number"`
	} `json:"ID"`
	ProofType int64 `json:"ProofType"`
}

type SectorCids struct {
	Unsealed struct {
		Field1 string `json:"/"`
	} `json:"Unsealed"`
	Sealed struct {
		Field1 string `json:"/"`
	} `json:"Sealed"`
}

type Seed struct {
	Value string `json:"Value"`
	Epoch int    `json:"Epoch"`
}

type Commit2Proof struct {
	TaskUuid  string `json:"task_uuid"`
	CpAddress string `json:"cp_address"`
	NodeId    string `json:"node_id"`
	TaskId    string `json:"task_id"`
	TaskType  string `json:"task_type"`
	Proof     string `json:"proof"`
}

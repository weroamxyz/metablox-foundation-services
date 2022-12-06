package models

import "time"

type WorkloadReq struct {
	Identity Identity `json:"identity" binding:"required"`
	Qos      Qos      `json:"qos" binding:"required"`
	Tracks   string   `json:"tracks"`
}

type Identity struct {
	Validator *VerifiablePresentation `json:"validator" binding:"required"`
	Miner     *VerifiablePresentation `json:"miner" binding:"required"`
}

type WorkloadDTO struct {
	Identity Identity `json:"identity" binding:"required"`
	Qos      Qos      `json:"qos" binding:"required"`
	Tracks   string   `json:"tracks"`
}

// WorkloadRecord dao layer
type WorkloadRecord struct {
	Id           int64     `json:"id" db:"id"`
	Miner        string    `json:"miner" db:"miner" binding:"required"`
	Validator    string    `json:"validator" db:"validator" binding:"required"`
	Qos          string    `json:"qos" db:"qos" binding:"required"`
	Tracks       string    `json:"tracks" db:"tracks" binding:"required"`
	CredentialID string    `json:"credentialID" db:"credential_id"`
	Model        string    `json:"model" db:"model"`
	Serial       string    `json:"serial" db:"serial"`
	Name         string    `json:"name" db:"name"`
	CreateTime   time.Time `json:"createTime" db:"create_time"`
}

type WorkloadDTOReq struct {
	Did     string `json:"did" db:"did"`
	BizDate string `json:"bizDate" db:"bizDate"`
}

type Qos struct {
	Bandwidth  string `json:"bandwidth"`
	Rssi       string `json:"rssi"`
	PacketLose string `json:"packetLose"`
	Latency    string `json:"latency"`
	Nonce      string `json:"nonce" binding:"required"`
	Signature  string `json:"signature" binding:"required"`
}

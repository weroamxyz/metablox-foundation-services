package models

import "time"

type WorkloadReq struct {
	Identity Identity `json:"identity" binding:required`
	Qos      string   `json:"qos" binding:required`
	Tracks   string   `json:"tracks" binding:required`
}

type Identity struct {
	Validator *VerifiablePresentation `json:"validator" binding:required`
	Miner     *VerifiablePresentation `json:"miner" binding:required`
}

type WorkloadDTO struct {
	Identity Identity `json:"identity" binding:required`
	Qos      string   `json:"qos" binding:required`
	Tracks   string   `json:"tracks" binding:required`
}

// Workload dao layer
type Workload struct {
	Id           int64     `json:"id" db:"id"`
	Miner        string    `json:"identity" db:"identity" binding:required`
	Validator    string    `json:"validator" db:"validator" binding:required`
	Qos          string    `json:"qos" db:"qos" binding:required`
	Tracks       string    `json:"tracks" db:"Tracks" binding:required`
	CredentialID string    `json:"credentialID" db:"credential_id"`
	Model        string    `json:"model" db:"model"`
	Serial       string    `json:"serial" db:"serial"`
	Name         string    `json:name db:"name"`
	CreateTime   time.Time `json:"createTime" db:"create_time"`
}

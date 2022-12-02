package models

import "github.com/shopspring/decimal"

type MinerRewardReq struct {
	DID string `json:"did,omitempty" db:"DID"`
}

type NFTRewardReq struct {
	ChainId         string `json:"chainId,omitempty" db:"ChainId"`
	TokenId         string `json:"tokenId,omitempty" db:"tokenId"`
	ContractAddress string `json:"contractAddress,omitempty" db:"ContractAddress"`
}

type MinersReq struct {
	Latitude  decimal.Decimal `json:"latitude" form:"latitude"`
	Longitude decimal.Decimal `json:"longitude" form:"latitude"`
	Distance  decimal.Decimal `json:"distance" form:"distance"`
}

type MinersDTO struct {
	Latitude  decimal.Decimal `json:"latitude"`
	Longitude decimal.Decimal `json:"longitude"`
	Distance  decimal.Decimal `json:"distance"`
}

type MinersWithDistanceDTO struct {
	Distance decimal.Decimal `json:"distance" db:"distance"`
	MinerInfo
}

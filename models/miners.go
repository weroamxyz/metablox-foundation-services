package models

import (
	"github.com/shopspring/decimal"
	"time"
)

type MinerInfo struct {
	ID             string          `db:"ID" json:"id"`
	Name           string          `db:"Name" json:"name"`
	SSID           string          `db:"SSID" json:"ssid"`
	BSSID          string          `db:"BSSID" json:"bssid"`
	Longitude      decimal.Decimal `db:"Longitude" json:"longitude"`
	Latitude       decimal.Decimal `db:"Latitude" json:"latitude"`
	OnlineStatus   bool            `db:"OnlineStatus" json:"onlineStatus"`
	MiningPower    float64         `db:"MiningPower" json:"miningPower"`
	IsMinable      bool            `db:"IsMinable" json:"isMinable"`
	DID            string          `db:"DID" json:"did"`
	IsVirtual      bool            `db:"IsVirtual" json:"isVirtual"`
	DeviceName     string          `db:"DeviceName" json:"deviceName"`
	Location       string          `db:"Location" json:"location"`
	SignalStrength string          `db:"SignalStrength" json:"signalStrength"`
	CreateTime     *time.Time      `db:"CreateTime" json:"createTime"`
}

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
	Longitude decimal.Decimal `json:"longitude" form:"longitude"`
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
	IsVirtual  bool   `db:"IsVirtual" json:"-"`
	CreateTime string `db:"createTime" json:"createTime"`
}

type MinerDetailReq struct {
	BSSID string `json:"bssid" form:"bssid" binding:"required"`
}

type MinerDetailReqDTO struct {
	BSSID string `json:"bssid"`
}

type MinerInfoDTO struct {
	ID             string          `db:"ID" json:"id"`
	Name           string          `db:"Name" json:"Name"`
	SSID           string          `db:"SSID" json:"SSID"`
	BSSID          string          `db:"longitude" json:"bssid"`
	Longitude      decimal.Decimal `db:"Longitude" json:"longitude"`
	Latitude       decimal.Decimal `db:"latitude" json:"latitude"`
	Availability   bool            `db:"availability" json:"Availability"`
	RewardEarned   decimal.Decimal `db:"rewardEarned" json:"RewardEarned"`
	MiningPower    float64         `db:"MiningPower" json:"miningPower"`
	IsMinable      bool            `db:"IsMinable" json:"isMinable"`
	DID            string          `db:"DID" json:"did"`
	DeviceName     string          `db:"DeviceName" json:"deviceName"`
	Address        string          `db:"address" json:"Address"`
	SignalStrength string          `db:"SignalStrength" json:"SignalStrength"`
	CreateTime     *time.Time      `db:"CreateTime" json:"createTime"`
}

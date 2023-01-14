package models

import (
	"github.com/shopspring/decimal"
	"time"
)

type RewardsRecord struct {
	Id             int64           `json:"id" db:"id"`
	BizDate        *time.Time      `json:"bizDate" db:"biz_date"`
	Did            string          `json:"did" db:"did"`
	OwnerAddress   string          `json:"owner_address" db:"owner_address"`
	UserType       string          `json:"user_type" db:"user_type"` // App,NFT
	Rewards        decimal.Decimal `json:"rewards" db:"rewards"`
	IsWithdrawn    bool            `json:"is_withdrawn" db:"is_withdrawn"`
	CreateTime     *time.Time      `json:"create_time" db:"create_time"`
	WithdrawalTime *time.Time      `json:"withdrawalTime" db:"withdrawal_time"`
}

package models

import "github.com/shopspring/decimal"

type AppRewardsPageReq struct {
	BizDate string `json:"bizDate"`
}

type AppRewardsPageDTO struct {
	BizDate      string          `json:"bizDate" db:"bizDate"`
	TotalRewards decimal.Decimal `json:"totalRewards" db:"totalRewards"`
	IsCollected  bool            `json:"isCollected" db:"isCollected"`
}

type AppTotalRewardsDTO struct {
	BizDate        string          `json:"bizDate" db:"bizDate"`
	TotalRewards   decimal.Decimal `json:"totalRewards" db:"totalRewards"`
	TotalWithdrawn decimal.Decimal `json:"totalWithdrawn" db:"totalWithdrawn"`
}

package models

import (
	"github.com/shopspring/decimal"
)

type Page struct {
	PageNum  uint64 `json:"pageNum" form:"pageNum,default=1"`
	PageSize uint64 `json:"pageSize" form:"pageSize,default=10"`
}

func (s *Page) PageInfo() (offset uint64, limit uint64) {
	if s.PageNum > 0 && s.PageSize > 0 {
		offset = (s.PageNum - 1) * s.PageSize
		limit = s.PageSize
	} else {
		offset = 0
		limit = 10
	}
	return
}

type AppRewardsPageReq struct {
	BizDate string `json:"bizDate"`
	Did     string `json:"did" binding:"required" form:"did"`
	Page
}

type AppRewardsPageReqDTO struct {
	AppRewardsPageReq
	UserType string
}

type AppRewardsPageDTO struct {
	Id          int64           `json:"id" db:"id"`
	BizDate     string          `json:"bizDate" db:"bizDate"`
	Rewards     decimal.Decimal `json:"rewards" db:"rewards"`
	IsWithdrawn bool            `json:"isWithdrawn" db:"isWithdrawn"`
}

type AppTotalRewardsReq struct {
	Did string `json:"did" form:"did" binding:"required"`
}

type AppTotalRewardsReqDTO struct {
	AppTotalRewardsReq
	UserType string
}

type AppTotalRewardsDTO struct {
	LatestWithdrawalTime string `json:"latestWithdrawalTime" db:"latestWithdrawalTime"`
	RewardsBalance       string     `json:"rewardsBalance" db:"rewardsBalance"`
	TotalWithdrawn       string     `json:"totalWithdrawn" db:"totalWithdrawn"`
}

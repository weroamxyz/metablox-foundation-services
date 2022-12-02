package service

import (
	"github.com/MetaBloxIO/metablox-foundation-services/models"
	"github.com/shopspring/decimal"
)

func GetAppRewardsPage(bizDate string) ([]*models.AppRewardsPageDTO, int64, error) {
	dtos := make([]*models.AppRewardsPageDTO, 0)
	dtos = append(dtos, &models.AppRewardsPageDTO{
		BizDate:      "2022-12-02",
		TotalRewards: decimal.NewFromFloat(10.89),
		IsCollected:  true,
	})
	dtos = append(dtos, &models.AppRewardsPageDTO{
		BizDate:      "2022-12-01",
		TotalRewards: decimal.NewFromFloat(8.56),
		IsCollected:  true,
	})
	return dtos, 2, nil
}

func GetAppTotalRewards() (*models.AppTotalRewardsDTO, error) {
	return &models.AppTotalRewardsDTO{
		BizDate:        "2022-12-01",
		TotalRewards:   decimal.Decimal{},
		TotalWithdrawn: decimal.Decimal{},
	}, nil
}

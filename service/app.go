package service

import (
	"github.com/MetaBloxIO/metablox-foundation-services/dao"
	"github.com/MetaBloxIO/metablox-foundation-services/models"
)

func GetAppRewardsPage(dto *models.AppRewardsPageReqDTO) ([]*models.AppRewardsPageDTO, int64, error) {
	return dao.SelectRewardRecordPage(dto)
}

func GetAppTotalRewards(dto *models.AppTotalRewardsReqDTO) (*models.AppTotalRewardsDTO, error) {
	return dao.SelectAppTotalRewards(dto)
}

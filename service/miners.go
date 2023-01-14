package service

import (
	"github.com/MetaBloxIO/metablox-foundation-services/dao"
	"github.com/MetaBloxIO/metablox-foundation-services/models"
)

func GetNearbyMinersList(dto *models.MinersDTO) ([]*models.MinersWithDistanceDTO, error) {

	return dao.SelectNearbyMinersList(dto)
}

func GetMinerDetailByBSSID(reqDto *models.MinerDetailReqDTO) (*models.MinerInfoDTO, error) {
	return dao.GetMinerInfo(reqDto)
}

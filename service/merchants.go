package service

import (
	"github.com/MetaBloxIO/metablox-foundation-services/dao"
	"github.com/MetaBloxIO/metablox-foundation-services/models"
)

func GetNearbyMerchantsList(dto *models.NearbyMerchantsReqDTO) (merchantsArr []*models.MerchantsInfo, err error) {

	return dao.GetNearbyMerchantsList(dto)
}

func GetMerchantDetailById(dto *models.MerchantDetailReqDTO) (merchantInfo *models.MerchantDetailDTO, err error) {

	return dao.GetMerchantDetail(dto)
}

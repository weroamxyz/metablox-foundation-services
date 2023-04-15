package service

import (
	"github.com/MetaBloxIO/metablox-foundation-services/dao"
	"github.com/MetaBloxIO/metablox-foundation-services/models"
)

func GetNearbyMerchantsList(dto *models.MerchantsReqDTO) (merchantsArr []*models.NearbyMerchantsDTO, err error) {

	return merchantsArr, nil
}

func GetMerchantDetailById(dto *models.MerchantDetailReqDTO) (merchantInfo *models.MerchantDetailDTO, err error) {

	return dao.GetMerchantDetail(dto)
}

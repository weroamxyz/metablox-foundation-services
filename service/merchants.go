package service

import (
	"github.com/MetaBloxIO/metablox-foundation-services/dao"
	"github.com/MetaBloxIO/metablox-foundation-services/models"
)

func GetNearbyMerchantsList(dto *models.MerchantsDTO) (merchantsArr []*models.MerchantsWithDistanceDTO, err error) {

	var merchantInfo *models.MerchantsWithDistanceDTO
	// Query the DB here using distance, get the Merchant ID and get the merchantInfo filled up in loops.
	merchantsArr = append(merchantsArr, merchantInfo)

	return merchantsArr, nil
}

func GetMerchantDetailById(dto *models.MerchantDetailDTO) (merchantInfo *models.MerchantsWithDistanceDTO, err error) {

	return dao.GetMerchantDetail(dto)
}

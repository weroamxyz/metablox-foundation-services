package dao

import "github.com/MetaBloxIO/metablox-foundation-services/models"

func GetMerchantDetail(dto *models.MerchantDetailReqDTO) (merchantDetail *models.MerchantDetailDTO, err error) {

	merchantDetail.MerchantId = dto.MerchantId

	return merchantDetail, nil
}

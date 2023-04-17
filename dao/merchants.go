package dao

import (
	"encoding/json"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/MetaBloxIO/metablox-foundation-services/comm/consts"
	"github.com/MetaBloxIO/metablox-foundation-services/models"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

func GetMerchantDetail(dto *models.MerchantDetailReqDTO) (merchantDetail *models.MerchantDetailDTO, err error) {

	merchantDetail.MerchantId = dto.MerchantId

	return merchantDetail, nil
}

func GetNearbyMerchantsList(dto *models.NearbyMerchantsReqDTO) (merchantList []*models.MerchantsInfo, err error) {

	if dto.Latitude.IsZero() || dto.Longitude.IsZero() {
		return nil, errors.New("both longitude and latitude are required")
	}

	// max 50km
	if dto.Distance.IsZero() || dto.Distance.GreaterThan(decimal.NewFromFloat(consts.MaxDistance)) {
		dto.Distance = decimal.NewFromFloat(consts.MaxDistance)
	}
	bytes, _ := json.Marshal(dto)
	fmt.Println(string(bytes))

	sql := squirrel.Select(` *,unix_timestamp(CreateTime) createTime,
	ROUND(
    IFNULL(6378.138 * 2 * ASIN(
      SQRT(
        POW(
          SIN(
            (
              ` + dto.Latitude.String() + ` * PI() / 180 - Latitude * PI() / 180
            ) / 2
          ), 2
        ) + COS(` + dto.Latitude.String() + ` * PI() / 180) * COS(Latitude * PI() / 180) * POW(
          SIN(
            (
              ` + dto.Longitude.String() + `* PI() / 180 - Longitude * PI() / 180
            ) / 2
          ), 2
        )
      )
    ),0),2) AS distance`).From("test_merchant_info").OrderBy(" distance ASC")

	sql = sql.Having("distance<=?", dto.Distance)
	var list = make([]*models.MerchantsInfo, 0)
	var rows *sqlx.Rows

	sqlStr, args, err := sql.ToSql()
	if err != nil {
		return list, err
	}
	rows, err = SqlDB.Queryx(sqlStr, args...)

	if err != nil {
		return list, err
	}
	defer rows.Close()

	for rows.Next() {
		merchantInfoDTO := &models.NearbyMerchantsDTO{}
		err = rows.StructScan(&merchantInfoDTO)
		if err != nil {
			return nil, err
		}
		businessHour := make([]string, 2)
		businessHour[0] = merchantInfoDTO.OpenHour
		businessHour[1] = merchantInfoDTO.CloseHour
		// Fill up other info here
		// Merchant contacts, like address, email, phone etc...
		contactExp := squirrel.Select("*").From("test_merchant_contact").Limit(1)
		contactExp = contactExp.Where("find_in_set(?,merchant_id)", merchantInfoDTO.MerchantId)
		sql, args, err := contactExp.ToSql()
		if err != nil {
			return nil, err
		}

		merchantContact := &merchantInfoDTO.ContactDTO
		err = SqlDB.Get(merchantContact, sql, args...)
		if err != nil {
			return nil, err
		}
		// Merchant resources files, like banner, logo etc...
		resourceExp := squirrel.Select("*").From("test_merchant_resource").Limit(1)
		resourceExp = resourceExp.Where("find_in_set(?,merchant_id)", merchantInfoDTO.MerchantId)
		sql, args, err = resourceExp.ToSql()
		if err != nil {
			return nil, err
		}

		merchantResource := &merchantInfoDTO.ResourceDTO
		err = SqlDB.Get(merchantResource, sql, args...)
		if err != nil {
			return nil, err
		}
		// Move data in sub DTOs to the main DTO
		merchantInfoDTO.Photo.Small = merchantResource.Small
		merchantInfoDTO.Photo.Medium = merchantResource.Medium
		merchantInfoDTO.Photo.Large = merchantResource.Large

		merchantInfoDTO.Contact.Email = merchantContact.Email
		merchantInfoDTO.Contact.Website = merchantContact.Website
		merchantInfoDTO.Contact.Phone = merchantContact.Phone
		merchantInfoDTO.Address.StreetName = merchantContact.StreetName
		merchantInfoDTO.Address.City = merchantContact.City
		merchantInfoDTO.Address.Province = merchantContact.Province
		merchantInfoDTO.Address.Postcode = merchantContact.Postcode
		merchantInfoDTO.Address.Country = merchantContact.Country

		// Transfer the DTO to the return struct
		var merchantInfo models.MerchantsInfo
		merchantInfoDTOBytes, err := json.Marshal(merchantInfoDTO)
		err = json.Unmarshal(merchantInfoDTOBytes, &merchantInfo)

		list = append(list, &merchantInfo)
	}

	return merchantList, err
}

package models

import "github.com/shopspring/decimal"

type MerchantsInfo struct {
	MerchantId string   `json:"merchantId"`
	Name       string   `json:"name"`
	Photo      struct { // Merchant Banner Images link for different sizes
		Small  string `json:"small"`
		Medium string `json:"medium"`
		Large  string `json:"large"`
	} `json:"photo"`
	Logo         string   `json:"logo"`         // Merchant Portrait/Icon Image link
	Category     string   `json:"category"`     // Business Type, eg. "BubbleTea","Bakery","Cafe"
	BusinessHour []string `json:"businessHour"` // Opening hours for the store, eg. 9am to 8pm will be ["9:00", "20:00"]
	TempClose    bool     `json:"tempClose"`    // Indication if the store is closed, overriding the normal business hours, should be updated by Merchant
	Contact      struct {
		Email   string `json:"email"`
		Website string `json:"website"`
		Phone   string `json:"phone"`
	} `json:"contact"`
	Address struct {
		StreetNumber string `json:"streetNumber"` // 888 Alpha Street
		City         string `json:"city"`         // Vancouver
		Province     string `json:"province"`     // BC
		Postcode     string `json:"postcode"`     // V1A2G3
		Country      string `json:"country"`      // Canada
	} `json:"address"`
	Points      int             `json:"points"`      // Points that can be earned by check-in, currenly flat 10 for all the store
	FreeParking bool            `json:"freeParking"` // If the store offers free-parking :)
	Latitude    decimal.Decimal `json:"latitude"`
	Longitude   decimal.Decimal `json:"longitude"`
	Distance    decimal.Decimal `json:"distance"`
}

type MerchantsWithDistanceDTO struct {
	Distance   decimal.Decimal `json:"distance" db:"distance"`
	InfoStruct MerchantsInfo
	CreateTime string `db:"createTime" json:"createTime"`
}

type MerchantsReq struct {
	Latitude  decimal.Decimal `json:"latitude" form:"latitude"`
	Longitude decimal.Decimal `json:"longitude" form:"longitude"`
	Distance  decimal.Decimal `json:"distance" form:"distance"`
}

type MerchantsDTO struct {
	Latitude  decimal.Decimal `json:"latitude"`
	Longitude decimal.Decimal `json:"longitude"`
	Distance  decimal.Decimal `json:"distance"`
}

type MerchantDetailReq struct {
	MerchantId string `json:"merchantId"`
}

type MerchantDetailDTO struct {
	MerchantId string `db:"merchantId" json:"merchantId"`
}

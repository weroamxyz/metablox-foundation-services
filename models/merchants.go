package models

import "github.com/shopspring/decimal"

type MerchantsInfo struct {
	MerchantId string   `json:"merchantId" db:"merchant_id"`
	Name       string   `json:"name" db:"name"`
	Photo      struct { // Merchant Banner Images link for different sizes
		Small  string `json:"small" db:"banner_small"`
		Medium string `json:"medium" db:"banner_medium"`
		Large  string `json:"large" db:"banner_large"`
	} `json:"photo"`
	Logo         string   `json:"logo" db:"logo"`            // Merchant Portrait/Icon Image link
	Category     string   `json:"category" db:"category"`    // Business Type, eg. "BubbleTea","Bakery","Cafe"
	BusinessHour []string `json:"businessHour"`              // Opening hours for the store, eg. 9am to 8pm will be ["9:00", "20:00"]
	TempClose    bool     `json:"tempClose" db:"temp_close"` // Indication if the store is closed, overriding the normal business hours, should be updated by Merchant
	Contact      struct {
		Email   string `json:"email" db:"email"`
		Website string `json:"website" db:"website"`
		Phone   string `json:"phone" db:"phone"`
	} `json:"contact"`
	Address struct {
		StreetName string `json:"streetName" db:"street_name"` // 888 Alpha Street
		City       string `json:"city" db:"city"`              // Vancouver
		Province   string `json:"province" db:"province"`      // BC
		Postcode   string `json:"postcode" db:"postcode"`      // V1A2G3
		Country    string `json:"country" db:"country"`        // Canada
	} `json:"address"`
	Points      int             `json:"points" db:"points"`            // Points that can be earned by check-in, currenly flat 10 for all the store
	FreeParking bool            `json:"freeParking" db:"free_parking"` // If the store offers free-parking :)
	Latitude    decimal.Decimal `json:"latitude" db:"latitude"`
	Longitude   decimal.Decimal `json:"longitude" db:"longitude"`
	Distance    decimal.Decimal `json:"distance,omitempty" db:"distance"`
}

type NearbyMerchantsReq struct {
	Latitude  decimal.Decimal `json:"latitude" form:"latitude"`
	Longitude decimal.Decimal `json:"longitude" form:"longitude"`
	Distance  decimal.Decimal `json:"distance" form:"distance"`
}

type NearbyMerchantsReqDTO struct {
	Latitude  decimal.Decimal `json:"latitude" form:"latitude"`
	Longitude decimal.Decimal `json:"longitude" form:"longitude"`
	Distance  decimal.Decimal `json:"distance" form:"distance"`
}

type MerchantDetailReq struct {
	MerchantId string `json:"merchantId"`
}

type MerchantDetailReqDTO struct {
	MerchantId string `json:"merchantId"`
}

type MerchantDetailDTO struct {
	MerchantsInfo
}

type NearbyMerchantsDTO struct {
	MerchantsInfo
	OpenHour   string `json:"openHour" db:"open_hour"`
	CloseHour  string `json:"closeHour" db:"close_hour"`
	ContactDTO struct {
		Email      string `json:"email" db:"email"`
		Website    string `json:"website" db:"website"`
		Phone      string `json:"phone" db:"phone"`
		StreetName string `json:"streetName" db:"street_name"` // 888 Alpha Street
		City       string `json:"city" db:"city"`              // Vancouver
		Province   string `json:"province" db:"province"`      // BC
		Postcode   string `json:"postcode" db:"postcode"`      // V1A2G3
		Country    string `json:"country" db:"country"`        // Canada
	} `json:"contactDTO"`
	ResourceDTO struct {
		Small  string `json:"small" db:"banner_small"`
		Medium string `json:"medium" db:"banner_medium"`
		Large  string `json:"large" db:"banner_large"`
		Logo   string `json:"logo" db:"logo"`
	} `json:"resourceDTO"`
	Distance decimal.Decimal `json:"distance" db:"distance"`
}

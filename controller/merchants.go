package controller

import (
	"github.com/MetaBloxIO/metablox-foundation-services/comm/requtil"
	"github.com/MetaBloxIO/metablox-foundation-services/models"
	"github.com/MetaBloxIO/metablox-foundation-services/service"
	"github.com/gin-gonic/gin"
)

func GetNearbyMerchantsListHandler(c *gin.Context) {

	req, err := requtil.ShouldBindQuery[models.MerchantsReq](c)
	if err != nil {
		return
	}

	if req.Longitude.IsZero() || req.Longitude.IsZero() {
		ResponseErrorWithMsg(c, CodeError, "both longitude and latitude are required")
		return
	}

	list, err := service.GetNearbyMerchantsList(&models.MerchantsReqDTO{
		Distance:  req.Distance,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
	})

	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	ResponseSuccess(c, list)
}

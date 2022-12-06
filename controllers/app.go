package controllers

import (
	"github.com/MetaBloxIO/metablox-foundation-services/comm/consts"
	"github.com/MetaBloxIO/metablox-foundation-services/comm/requtil"
	"github.com/MetaBloxIO/metablox-foundation-services/models"
	"github.com/MetaBloxIO/metablox-foundation-services/service"
	"github.com/gin-gonic/gin"
)

func GetAppRewardsPageHandler(c *gin.Context) {

	req, err := requtil.ShouldBindQuery[models.AppRewardsPageReq](c)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	m := &models.AppRewardsPageReqDTO{
		AppRewardsPageReq: *req,
		UserType:          consts.AppUser,
	}
	list, total, err := service.GetAppRewardsPage(m)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}
	ResponseSuccessWithPageData(c, list, total)
}

func GetAppTotalRewardsHandler(c *gin.Context) {

	req, err := requtil.ShouldBindQuery[models.AppTotalRewardsReq](c)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	data, err := service.GetAppTotalRewards(&models.AppTotalRewardsReqDTO{
		AppTotalRewardsReq: *req,
		UserType:           consts.AppUser,
	})
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}
	ResponseSuccess(c, data)
}

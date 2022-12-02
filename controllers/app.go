package controllers

import (
	"github.com/MetaBloxIO/metablox-foundation-services/service"
	"github.com/gin-gonic/gin"
)

func GetAppRewardsPageHandler(c *gin.Context) {

	value := c.Query("bizDate")

	list, total, err := service.GetAppRewardsPage(value)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	ResponseSuccessWithPage(c, list, total)
}
func GetAppTotalRewardsHandler(c *gin.Context) {

	data, err := service.GetAppTotalRewards()
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}
	ResponseSuccess(c, data)
}

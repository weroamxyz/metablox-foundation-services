package controllers

import "C"
import (
	"github.com/MetaBloxIO/metablox-foundation-services/comm/requtil"
	"github.com/MetaBloxIO/metablox-foundation-services/models"
	"github.com/MetaBloxIO/metablox-foundation-services/service"
	"github.com/gin-gonic/gin"
)

func WorkloadValidateHandler(c *gin.Context) {

	req, err := requtil.ShouldBindJSON[models.WorkloadReq](c)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	dto := &models.WorkloadDTO{
		Identity: req.Identity,
		Qos:      req.Qos,
		Tracks:   req.Tracks,
	}

	err = service.WorkloadValidate(dto)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}
	ResponseSuccessWithMsg(c, "success")
}

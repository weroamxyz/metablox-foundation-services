package controller

import (
	"github.com/MetaBloxIO/metablox-foundation-services/comm/requtil"
	"github.com/MetaBloxIO/metablox-foundation-services/models"
	"github.com/MetaBloxIO/metablox-foundation-services/service"
	"github.com/gin-gonic/gin"
)

func EmailSubmissionhandler(c *gin.Context) {
	req, err := requtil.ShouldBindJSON[models.EmailSubmission](c)
	if err != nil {
		return
	}

	err = service.HandleEmailSubmission(req)

	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	ResponseSuccessWithMsg(c, "success")
}

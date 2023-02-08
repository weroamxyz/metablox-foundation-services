package controller

import (
	"github.com/MetaBloxIO/metablox-foundation-services/comm/requtil"
	"github.com/MetaBloxIO/metablox-foundation-services/models"
	"github.com/MetaBloxIO/metablox-foundation-services/resources"
	"github.com/MetaBloxIO/metablox-foundation-services/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetWifiUserInfoHandler(c *gin.Context) {
	vp, err := requtil.ShouldBindJSON[models.VerifiablePresentation](c)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}

	info, err := service.GetWifiUserInfo(vp)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}
	ResponseSuccess(c, info)
}

func GetWifiCertFileHandler(c *gin.Context) {

	c.FileFromFS(resources.WifiServerFileName, http.FS(resources.EmbedFS))
}

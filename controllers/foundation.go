package controllers

import "C"
import (
	"github.com/MetaBloxIO/metablox-foundation-services/service"
	"github.com/gin-gonic/gin"
)

func WorkProofHandler(c *gin.Context) {

	hash, err := RegisterDIDForUser(c)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}
	service.WorkProof()
	ResponseSuccessWithMsgAndData(c, "DID has been successfully registered to registry", hash)
}

func GetProfitHandler(c *gin.Context) {
	did := c.Query("did")

	total, err := service.GetProfit(did)
	if err != nil {
		ResponseErrorWithMsg(c, CodeError, err.Error())
		return
	}
	ResponseSuccess(c, gin.H{"total": total.String()})
}

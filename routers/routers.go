package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/metabloxDID/controllers"
)

func Setup() {
	r := gin.New()

	r.POST("/registry/storedoc", controllers.SendDocToRegistryHandler)

	r.POST("/vc/issue/wifi/:did", controllers.IssueWifiVCHandler)
	r.POST("/vc/renew/wifi/:did", controllers.RenewVCHandler)
	r.POST("/vc/revoke/wifi/:did", controllers.RevokeVCHandler)

	r.POST("/vc/issue/mining/:did", controllers.IssueMiningVCHandler)
	r.POST("/vc/renew/mining/:did", controllers.RenewVCHandler)
	r.POST("/vc/revoke/mining/:did", controllers.RevokeVCHandler)

	r.GET("/minerlist/:did", controllers.GetMinerListHandler)

	r.GET("/nonce", controllers.GenerateNonceHandler)

	r.GET("/testing/signatures/:message", controllers.GenerateTestSignatures) //todo: don't leave this active in any release version as it is only for testing
	r.Run(":8888")
}

package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/MetaBloxIO/metablox-foundation-services/controllers"
)

func Setup() {
	r := gin.New()

	r.POST("/registry/storedoc", controllers.SendDocToRegistryHandler)

	r.POST("/vc/wifi/issue", controllers.IssueWifiVCHandler)
	r.POST("/vc/wifi/renew", controllers.RenewVCHandler)
	r.POST("/vc/wifi/revoke", controllers.RevokeVCHandler)

	r.POST("/vc/mining/issue", controllers.IssueMiningVCHandler)
	r.POST("/vc/mining/renew", controllers.RenewVCHandler)
	r.POST("/vc/mining/revoke", controllers.RevokeVCHandler)

	r.GET("/minerlist", controllers.GetMinerListHandler)

	r.GET("/nonce", controllers.GenerateNonceHandler)

	r.GET("/testing/signatures/:message", controllers.GenerateTestSignatures) //todo: don't leave this active in any release version as it is only for testing
	r.POST("/testing/assignissuer", controllers.AssignIssuer)
	r.POST("/testing/updatevc", controllers.SetVCAttribute)
	r.POST("/testing/readvcchanged", controllers.ReadVCChangedEvents)
	r.Run(":8888")
}

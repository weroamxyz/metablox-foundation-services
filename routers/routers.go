package routers

import (
	"github.com/MetaBloxIO/metablox-foundation-services/controllers"
	"github.com/gin-gonic/gin"
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

	r.POST("/vc/staking/issue", controllers.IssueStakingVCHandler)
	r.POST("/vc/staking/renew", controllers.RenewVCHandler)
	r.POST("/vc/staking/revoke", controllers.RevokeVCHandler)

	r.GET("/minerlist", controllers.GetMinerListHandler)

	r.GET("/nonce", controllers.GenerateNonceHandler)

	r.GET("/pubkey", controllers.GetIssuerPublicKeyHandler)

	r.GET("/testing/signatures/:message", controllers.GenerateTestSignatures) //todo: don't leave this active in any release version as it is only for testing
	r.POST("/testing/assignissuer", controllers.AssignIssuer)
	r.POST("/testing/updatevc", controllers.SetVCAttribute)
	r.POST("/testing/readvcchanged", controllers.ReadVCChangedEvents)
	r.POST("/testing/signpresentation", controllers.GenerateTestingPresentationSignatures)

	r.Run(":8888")
}

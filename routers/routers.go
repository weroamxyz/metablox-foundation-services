package routers

import (
	"github.com/MetaBloxIO/metablox-foundation-services/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func Setup() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	r.POST("/registry/storedoc", controllers.SendDocToRegistryHandler)
	r.POST("/registry/storedid", controllers.RegisterDIDHandler)

	r.POST("/vc/wifi/issue", controllers.IssueWifiVCHandler)
	r.POST("/vc/wifi/renew", controllers.RenewVCHandler)
	r.POST("/vc/wifi/revoke", controllers.RevokeVCHandler)
	r.GET("/vc/wifi/userInfo", controllers.GetWifiUserInfoHandler)
	r.GET("/vc/wifi/certFile", controllers.GetWifiCertFileHandler)

	r.POST("/vc/mining/issue", controllers.IssueMiningVCHandler)
	r.POST("/vc/mining/renew", controllers.RenewVCHandler)
	r.POST("/vc/mining/revoke", controllers.RevokeVCHandler)

	r.POST("/vc/staking/issue", controllers.IssueStakingVCHandler)
	r.POST("/vc/staking/renew", controllers.RenewVCHandler)
	r.POST("/vc/staking/revoke", controllers.RevokeVCHandler)

	r.GET("/nonce", controllers.GenerateNonceHandler)

	r.GET("/pubkey", controllers.GetIssuerPublicKeyHandler)

	r.GET("/testing/signatures/:message", controllers.GenerateTestSignatures) //todo: don't leave this active in any release version as it is only for testing
	r.POST("/testing/assignissuer", controllers.AssignIssuer)
	r.POST("/testing/updatevc", controllers.SetVCAttribute)
	r.POST("/testing/readvcchanged", controllers.ReadVCChangedEvents)
	r.POST("/testing/signpresentation", controllers.GenerateTestingPresentationSignatures)

	r.POST("/workload/validate", controllers.WorkloadValidateHandler)
	r.GET("/miners", controllers.GetNearbyMinersListHandler)
	r.GET("/app/rewardsPage", controllers.GetAppRewardsPageHandler)
	r.GET("/app/totalRewards", controllers.GetAppTotalRewardsHandler)

	r.Run(":8888")
}

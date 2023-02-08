package router

import (
	"github.com/MetaBloxIO/metablox-foundation-services/comm/log"
	"github.com/MetaBloxIO/metablox-foundation-services/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var deprecatedAPIs = map[string]bool{
	"/registry/storedid": true,
	"/vc/wifi/userInfo":  true,
	"/vc/wifi/certFile":  true,
	"/vc/wifi/issue":     true,
	"/vc/wifi/renew":     true,
	"/vc/wifi/revoke":    true,
	"/vc/mining/issue":   true,
	"/vc/mining/renew":   true,
	"/vc/mining/revoke":  true,
	"/nonce":             true,
	"/pubkey":            true,
	"/workload/validate": true,
	"/miners":            true,
	"/miner/getByBssid":  true,
	"/miner/detail":      true,
	"/miner/heartbeat":   true,
	"/app/rewardsPage":   true,
	"/app/totalRewards":  true}

func ForceUpgradeTips() gin.HandlerFunc {

	return func(c *gin.Context) {
		if deprecatedAPIs[c.Request.URL.Path] {
			c.JSON(http.StatusOK, gin.H{
				"code": 400,
				"msg":  "please upgrade your app version",
				"data": nil,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

func Setup() {
	r := gin.New()
	r.Use(gin.LoggerWithWriter(log.GetLogWriter()), gin.RecoveryWithWriter(log.GetLogWriter()))
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.Use(ForceUpgradeTips())

	v1 := r.Group("/foundation/v1")

	v1.POST("/registry/storedid", controller.RegisterDIDHandler)

	v1.POST("/vc/wifi/issue", controller.IssueWifiVCHandler)
	v1.POST("/vc/wifi/renew", controller.RenewVCHandler)
	v1.POST("/vc/wifi/revoke", controller.RevokeVCHandler)
	v1.POST("/vc/wifi/userInfo", controller.GetWifiUserInfoHandler)
	v1.GET("/vc/wifi/certFile", controller.GetWifiCertFileHandler)
	v1.POST("/vc/mining/issue", controller.IssueMiningVCHandler)
	v1.POST("/vc/mining/renew", controller.RenewVCHandler)
	v1.POST("/vc/mining/revoke", controller.RevokeVCHandler)

	v1.GET("/nonce", controller.GenerateNonceHandler)
	v1.GET("/pubkey", controller.GetIssuerPublicKeyHandler)

	v1.POST("/workload/validate", controller.WorkloadValidationHandler)
	v1.GET("/miners", controller.GetNearbyMinersListHandler)
	v1.GET("/miner/getByBssid", controller.GetMinerDetailHandler)
	v1.GET("/miner/detail", controller.GetMinerDetailHandler)
	v1.POST("/miner/heartbeat", controller.HeartbeatHandler)
	v1.GET("/app/rewardsPage", controller.GetAppRewardsPageHandler)
	v1.GET("/app/totalRewards", controller.GetAppTotalRewardsHandler)

	r.Run(":8888")
}

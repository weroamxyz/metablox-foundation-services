package routers

import (
	"github.com/MetaBloxIO/metablox-foundation-services/comm/log"
	"github.com/MetaBloxIO/metablox-foundation-services/controllers"
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

	v1.POST("/registry/storedid", controllers.RegisterDIDHandler)

	v1.POST("/vc/wifi/issue", controllers.IssueWifiVCHandler)
	v1.POST("/vc/wifi/renew", controllers.RenewVCHandler)
	v1.POST("/vc/wifi/revoke", controllers.RevokeVCHandler)
	v1.POST("/vc/wifi/userInfo", controllers.GetWifiUserInfoHandler)
	v1.GET("/vc/wifi/certFile", controllers.GetWifiCertFileHandler)
	v1.POST("/vc/mining/issue", controllers.IssueMiningVCHandler)
	v1.POST("/vc/mining/renew", controllers.RenewVCHandler)
	v1.POST("/vc/mining/revoke", controllers.RevokeVCHandler)

	v1.GET("/nonce", controllers.GenerateNonceHandler)
	v1.GET("/pubkey", controllers.GetIssuerPublicKeyHandler)

	v1.POST("/workload/validate", controllers.WorkloadValidationHandler)
	v1.GET("/miners", controllers.GetNearbyMinersListHandler)
	v1.GET("/miner/getByBssid", controllers.GetMinerDetailHandler)
	v1.GET("/miner/detail", controllers.GetMinerDetailHandler)
	v1.POST("/miner/heartbeat", controllers.HeartbeatHandler)
	v1.GET("/app/rewardsPage", controllers.GetAppRewardsPageHandler)
	v1.GET("/app/totalRewards", controllers.GetAppTotalRewardsHandler)

	r.Run(":8888")
}

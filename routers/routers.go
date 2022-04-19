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

	r.GET("/testing/signatures/:message", controllers.GenerateTestSignatures)
	r.Run(":8888")
}

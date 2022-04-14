package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/metabloxDID/controllers"
)

func Setup() {
	r := gin.New()

	r.POST("/registry/storedoc", controllers.SendDocToRegistryHandler)

	r.POST("/vc/issue/wifi", controllers.IssueWifiVCHandler)
	r.POST("/vc/renew/wifi", controllers.RenewVCHandler)
	r.POST("/vc/revoke/wifi", controllers.RevokeVCHandler)

	r.POST("/vc/issue/mining", controllers.IssueMiningVCHandler)
	r.POST("/vc/renew/mining", controllers.RenewVCHandler)
	r.POST("/vc/revoke/mining", controllers.RevokeVCHandler)
	r.Run(":8888")
}

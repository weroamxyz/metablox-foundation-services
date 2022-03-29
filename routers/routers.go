package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/metabloxDID/controllers"
)

func Setup() {
	r := gin.New()

	r.POST("/registry/storedoc", controllers.SendDocToRegistryHandler)

	r.POST("/vc/issue", controllers.IssueVCHandler)

	r.Run(":8888")
}

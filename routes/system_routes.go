package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/naufal/simba-qr-app/controllers"
)

func SystemRoute(router *gin.Engine) {
	systemGroup := router.Group("/v1/system")
	{
		systemGroup.POST("/read-key", controllers.UploadQRCode)
		systemGroup.POST("/generate-key", controllers.CreateSystem)
		systemGroup.DELETE("/:id", controllers.DeleteSystem)
	}
}

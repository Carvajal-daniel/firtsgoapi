package router

import (
	"github.com/Carvajal-daniel/firtsgoapi/handlers"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{

		v1.GET("/opening", handlers.ShowHandlerOpening)

		v1.POST("/opening", handlers.CreateHandlerOpening)

		v1.DELETE("/opening", handlers.DeleteHandlerOpening)

		v1.PUT("/opening", handlers.UpdateHandlerOpening)

		v1.GET("/openings", handlers.ListHandlerOpening)
	}

}

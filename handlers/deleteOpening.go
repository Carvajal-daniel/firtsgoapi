package handlers

import "github.com/gin-gonic/gin"

func DeleteHandlerOpening(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Delete Opening",
	})
}

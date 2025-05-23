package handlers

import "github.com/gin-gonic/gin"

func ListHandlerOpening(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "List Opening",
	})
}

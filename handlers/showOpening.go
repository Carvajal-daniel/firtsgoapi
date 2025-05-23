package handlers

import "github.com/gin-gonic/gin"

func ShowHandlerOpening(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Show Opening",
	})
}

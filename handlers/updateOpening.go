package handlers

import "github.com/gin-gonic/gin"

func UpdateHandlerOpening(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Update Opening",
	})
}

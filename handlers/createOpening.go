package handlers

import "github.com/gin-gonic/gin"

func CreateHandlerOpening(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Post Opening",
	})
}

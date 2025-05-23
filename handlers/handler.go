package handlers

import "github.com/gin-gonic/gin"

func CreateHandlerOpening(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Post Opening",
	})
}
func ShowHandlerOpening(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Show Opening",
	})
}
func DeleteHandlerOpening(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Delete Opening",
	})
}
func UpdateHandlerOpening(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Update Opening",
	})
}
func ListHandlerOpening(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "List Opening",
	})
}

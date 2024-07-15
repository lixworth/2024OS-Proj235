package services

import "github.com/gin-gonic/gin"

func LoginHandler(context *gin.Context) {
	context.JSON(200, gin.H{
		"msg": "hello world",
	})
}

package services

import (
	"github.com/gin-gonic/gin"
	"ota-updater/internal/common"
)

type LoginRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

func LoginHandler(context *gin.Context) {
	var req LoginRequest
	err := context.ShouldBindBodyWithJSON(&req)
	if err != nil {
		context.JSON(200, common.Response(1, nil))
		return
	}

	context.JSON(200, gin.H{
		"msg": "hello world",
	})
}

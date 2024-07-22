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
	res := common.Response(1, nil)
	if req.Account == "root" && req.Password == "passwordroot" { // 临时
		res.Error = 0
		res.Data = struct {
			Token string `json:"token"`
		}{
			Token: "token",
		}
	} else {
		res.Error = 1
	}
	context.JSON(200, res)
}

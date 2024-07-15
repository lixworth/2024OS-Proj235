package internal

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ota-updater/internal/common"
	"ota-updater/internal/flag"
	"ota-updater/services"
)

func RegisterRouter(router *gin.Engine) {
	// frontend: client ui
	router.Static("/_nuxt", flag.GetFlags().NuxtOutput+"/_nuxt")
	router.StaticFile("/favicon.ico", flag.GetFlags().NuxtOutput+" /favicon.ico")
	router.StaticFile("/", flag.GetFlags().NuxtOutput+"/index.html")

	// backend: api
	router.POST("/api/auth", services.LoginHandler)
	router.GET("/api/info", services.InfoHandler)                   // 系统信息
	router.GET("/api/check-status", services.CheckStatusHandler)    // 检测更新状态
	router.POST("/api/update", services.UpdateHandler)              // 提交更新任务到最新版本
	router.POST("/api/check-version", services.CheckVersionHandler) // 检测更新
	router.POST("/api/manual-update", services.ManualUpdateHandler) // 手动更新
}

func InitHTTPServer() {
	router := gin.Default()
	if common.IsRelease() {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	RegisterRouter(router)
	router.NoRoute(func(context *gin.Context) {
		context.Redirect(http.StatusFound, "/#/") // hash router mode
	})
	err := router.Run(flag.GetFlags().ListenAddress)
	if err != nil {
		panic(err)
	}
}

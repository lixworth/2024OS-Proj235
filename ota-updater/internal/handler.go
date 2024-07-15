package internal

import (
	"github.com/gin-gonic/gin"
	"ota-updater/internal/flag"
	"ota-updater/services"
)

func RegisterRouter(router *gin.Engine) {
	// frontend: client ui
	router.Static("/_nuxt", flag.GetFlags().NuxtOutput+"/_nuxt")
	router.StaticFile("/favicon.ico", flag.GetFlags().NuxtOutput+" /favicon.ico")
	router.StaticFile("/", flag.GetFlags().NuxtOutput+"/index.html")

	// backend: api
	router.POST("/auth", services.LoginHandler)
	router.GET("/info", services.InfoHandler)                   // 系统信息
	router.GET("/check-status", services.CheckStatusHandler)    // 检测更新状态
	router.POST("/update", services.UpdateHandler)              // 提交更新任务到最新版本
	router.POST("/check-version", services.CheckVersionHandler) // 检测更新
	router.POST("/manual-update", services.ManualUpdateHandler) // 手动更新
}

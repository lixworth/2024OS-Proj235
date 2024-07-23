package internal

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"ota-updater/internal/common"
	"ota-updater/internal/flag"
	"ota-updater/internal/storage"
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

func CheckInitial() bool {
	res, err := storage.GetValue("initial")
	fmt.Println(res)
	if err != nil {
		return false
	}
	return true
}

func Bootstrap() {
	// 检测初次启动
	if CheckInitial() {
		err := storage.SetValue("initial", "okay")
		if err != nil {
			panic(err)
		}
		hash, err := common.PasswordHash("passwordroot")
		if err != nil {
			panic(err)
		}
		err = storage.SetValue("password", hash)
		if err != nil {
			panic(err)
		}
	}
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

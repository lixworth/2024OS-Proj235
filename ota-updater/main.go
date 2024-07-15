package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"ota-updater/internal"
	"ota-updater/internal/flag"
	"ota-updater/internal/storage"
)

func main() {
	storage.Test()
	err := flag.InitFlagsHandler()
	if err != nil {
		panic(err)
	}
	_, err = os.Stat(flag.GetFlags().NuxtOutput)
	if err != nil {
		fmt.Println("请先构建 ClientUI 项目")
		panic(err)
	}
	router := gin.Default()
	internal.RegisterRouter(router)
	router.NoRoute(func(context *gin.Context) {
		context.Redirect(http.StatusFound, "/#/")
	})
	err = router.Run(":9301")
	if err != nil {
		panic(err)
	}
	return
}

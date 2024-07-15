package main

import (
	"fmt"
	"os"
	"ota-updater/internal"
	"ota-updater/internal/flag"
	"ota-updater/internal/storage"
)

func CheckEnvironment() {
	_, err := os.Stat(flag.GetFlags().NuxtOutput)
	if err != nil {
		fmt.Println("请先构建 ClientUI 项目")
		panic(err)
	}
}

func main() {
	err := flag.InitFlagsHandler()
	if err != nil {
		panic(err)
	}

	CheckEnvironment()
	storage.InitBadgerDB()
	internal.InitHTTPServer()
	return
}

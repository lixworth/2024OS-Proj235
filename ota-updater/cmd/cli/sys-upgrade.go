package main

import (
	"flag"
	"ota-updater/internal/logger"
)

type Flag struct {
	NuxtOutput    string
	StoragePath   string
	ListenAddress string
}

var flags Flag

func main() {
	logger.InitLogger()
	logger.GetLogger().Info("SystemUpgrade", "OTA-UPGRADE CLI Project")

	flag.Parse()
	flags = Flag{}
}

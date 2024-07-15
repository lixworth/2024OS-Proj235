package flag

import (
	"flag"
	"ota-updater/internal/common"
)

type Flag struct {
	NuxtOutput    string
	StoragePath   string
	ListenAddress string
}

var flags Flag

func InitFlagsHandler() error {
	var nuxtDefaultPath string
	if common.IsRelease() {
		nuxtDefaultPath = "./client-ui"
	} else {
		nuxtDefaultPath = "./../client-ui/.output/public"
	} // todo 将静态文件打包进升级程序本体
	NuxtOutput := flag.String("nuxt_output", nuxtDefaultPath, "Nuxt Output directory")
	StoragePath := flag.String("storage_path", "./runtime/storage", "DB Storage Path")
	ListenAddress := flag.String("listen_address", "127.0.0.1:9301", "HTTP Server Listen Address")
	flag.Parse()

	flags = Flag{
		NuxtOutput:    *NuxtOutput,
		StoragePath:   *StoragePath,
		ListenAddress: *ListenAddress,
	}
	return nil
}

func GetFlags() Flag {
	return flags
}

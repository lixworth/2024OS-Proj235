package flag

import (
	"flag"
)

type Flag struct {
	NuxtOutput  string
	StoragePath string
}

var flags Flag

func InitFlagsHandler() error {
	NuxtOutput := flag.String("nuxt_output", "./../client-ui/.output/public", "Nuxt Output directory")
	StoragePath := flag.String("storage_path", "./storage", "DB Storage Path")
	flag.Parse()

	flags = Flag{
		NuxtOutput:  *NuxtOutput,
		StoragePath: *StoragePath,
	}
	return nil
}

func GetFlags() Flag {
	return flags
}

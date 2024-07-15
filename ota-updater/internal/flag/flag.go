package flag

import (
	"flag"
)

type Flag struct {
	RuntimeDir string
	ConfigFile string
	NuxtOutput string
}

var flags Flag

func InitFlagsHandler() error {
	NuxtOutput := flag.String("nuxt_output", "./../client-ui/.output/public", "Nuxt Output directory")
	RuntimeDir := flag.String("runtime_dir", "./runtime/", "Runtime directory")
	ConfigFile := flag.String("config_file", "./config.toml", "Config file path")
	flag.Parse()

	flags = Flag{
		RuntimeDir: *RuntimeDir,
		ConfigFile: *ConfigFile,
		NuxtOutput: *NuxtOutput,
	}
	return nil
}

func GetFlags() Flag {
	return flags
}

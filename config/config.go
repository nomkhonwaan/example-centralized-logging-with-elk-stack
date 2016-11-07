package config

import (
	"flag"
	"log"

	"github.com/spf13/viper"
)

type StresserConfigType struct {
	ConfigName string
	ConfigPath string
	ConfigType string
	Debug      bool
}

var StresserConfig = StresserConfigType{}

func init() {
	flag.StringVar(&(StresserConfig.ConfigName), "config-name", "config", "Configuration filename, default is [config]")
	flag.StringVar(&(StresserConfig.ConfigPath), "config-path", ".", "Configuration path, default is [.]")
	flag.StringVar(&(StresserConfig.ConfigType), "config-type", "yaml", "Configuration type, default is [yaml]")
	flag.BoolVar(&(StresserConfig.Debug), "debug", false, "Enable debug mode, default is [0]")
	flag.Parse()

	viper.SetConfigType(StresserConfig.ConfigType)
	viper.AddConfigPath(StresserConfig.ConfigPath)
	viper.AddConfigPath("$HOME/.stresser")
	viper.SetConfigName(StresserConfig.ConfigName)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
}

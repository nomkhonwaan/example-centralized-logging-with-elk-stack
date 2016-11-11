package config

import (
	"flag"
	"log"

	"github.com/spf13/viper"
)

type configuration struct {
	ConfigName string
	ConfigPath string
	ConfigType string
	Debug      bool
}

var LoggerConfig = configuration{}

func init() {
	flag.StringVar(&(LoggerConfig.ConfigName), "config-name", "config", "Configuration filename, default is [config]")
	flag.StringVar(&(LoggerConfig.ConfigPath), "config-path", ".", "Configuration path, default is [.]")
	flag.StringVar(&(LoggerConfig.ConfigType), "config-type", "yaml", "Configuration type, default is [yaml]")
	flag.BoolVar(&(LoggerConfig.Debug), "debug", false, "Enable debug mode, default is [0]")
	flag.Parse()

	viper.SetConfigType(LoggerConfig.ConfigType)
	viper.AddConfigPath(LoggerConfig.ConfigPath)
	viper.AddConfigPath("$HOME/.logger")
	viper.SetConfigName(LoggerConfig.ConfigName)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
}

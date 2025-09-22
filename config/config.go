package config

import (
	"flag"
	"runtime"

	"github.com/spf13/viper"
)

type Config struct {
	Http HttpConfig
}

type HttpConfig struct {
	Port int64
}

// 定义全局变量
var config Config

var configFilePath = flag.String("config", "", "config file path")

func LoadConfig() error {
	flag.Parse()
	if *configFilePath == "" {
		if runtime.GOOS == "windows" {
			*configFilePath = ".\\config\\config.yaml"
		} else {
			*configFilePath = "./config/config.yaml"
		}
	}
	return load(*configFilePath)
}

func Get() *Config {
	return &config
}

func load(filePath string) error {
	viper.SetConfigFile(filePath)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return err
	}

	return nil
}

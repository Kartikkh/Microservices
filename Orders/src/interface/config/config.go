package config

import (
	"github.com/Orders/src/util"
)

var config *Config

func Init() (*Config, error) {
	configFileName := "config.json"
	dirName := "files/orders"
	config = &Config{}
	err := util.ReadJsonFile(dirName, configFileName, config)
	if err != nil {
		return config, err
	}
	return config, nil
}

func GetConfig() *Config {
	return config
}

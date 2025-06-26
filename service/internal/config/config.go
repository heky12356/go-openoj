package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port string `mapstructure:"port"`
	} `mapstructure:"server"`

	Database struct {
		Host     string `mapstructure:"host"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Name     string `mapstructure:"name"`
	} `mapstructure:"database"`
}

var config Config

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../../../config/")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("read config file error: %s \n", err))
	}

	if err := viper.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("unmarshal config file error: %s \n", err))
	}
}

func GetConfig() Config {
	return config
}

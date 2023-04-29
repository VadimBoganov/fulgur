package config

import (
	"github.com/spf13/viper"
)

type Config struct{
	DatabasePath string `mapstructure:"databasePath"`
	Logger LoggerConfig `mapstructure:"logger"`
}

type LoggerConfig struct{
	Path string `mapstructure:"path"`
	FileName string `mapstructure:"fileName"`
}

func (c *Config) InitConfig(path string) error { 
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	if err := viper.ReadInConfig(); err != nil{
		return err
	}

	return viper.Unmarshal(&c)
}
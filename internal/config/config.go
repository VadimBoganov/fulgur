package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DatabasePath  string       `mapstructure:"databasePath"`
	FTPPath       string       `mapstructure:"ftpPath"`
	LocalFilePath string       `mapstructure:"localFilePath"`
	FtpUrl        string       `mapstructure:"ftpUrl"`
	Logger        LoggerConfig `mapstructure:"logger"`
	Server        ServerConfig `mapstructure:"server"`
}

type LoggerConfig struct {
	Path     string `mapstructure:"path"`
	FileName string `mapstructure:"fileName"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
}

var config Config

func InitConfig(path string) error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return viper.Unmarshal(&config)
}

func GetConfig() *Config {
	return &config
}

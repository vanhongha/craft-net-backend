package config

import (
	"github.com/spf13/viper"

	"craftnet/internal/util"
)

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type Config struct {
	Database DatabaseConfig
}

var AppConfig *Config

func LoadConfig() {
	viper.SetConfigName("config") // config file name
	viper.SetConfigType("yaml")   // config extension
	viper.AddConfigPath("../../") // config path

	// read config file
	if err := viper.ReadInConfig(); err != nil {
		util.GetLogger().LogErrorWithMsgAndError("Error reading config file", err)
	}

	// decode config into struct
	AppConfig = &Config{}
	if err := viper.Unmarshal(AppConfig); err != nil {
		util.GetLogger().LogErrorWithMsgAndError("Unable to decode config into struct", err)
	}

	util.GetLogger().LogInfo("Config loaded successfully")
}
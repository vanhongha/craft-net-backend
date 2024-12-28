package config

import (
	"craftnet/internal/util"
	"fmt"

	"github.com/spf13/viper"
)

type JwtAuthConfig struct {
	Secret string `mapstructure:"secret"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
	SSLMode  string `mapstructure:"sslmode"`
}

type Config struct {
	Jwt      JwtAuthConfig  `mapstructure:"jwt"`
	Database DatabaseConfig `mapstructure:"database"`
}

var AppConfig *Config

func GetAppConfig() *Config {
	if AppConfig == nil {
		LoadConfig("../")
	}
	if AppConfig == nil {
		util.GetLogger().LogErrorWithMsg("Please load config first", true)
	}
	return AppConfig
}

func LoadConfig(path string) {
	viper.SetConfigName("config") // config file name
	viper.SetConfigType("yaml")   // config extension
	viper.AddConfigPath(path)     // config path

	// read config file
	if err := viper.ReadInConfig(); err != nil {
		util.GetLogger().LogErrorWithMsgAndError("Error reading config file", err, false)
	}

	// decode config into struct
	AppConfig = &Config{}
	if err := viper.Unmarshal(AppConfig); err != nil {
		util.GetLogger().LogErrorWithMsgAndError("Unable to decode config into struct", err, false)
	}

	fmt.Println(222)

	util.GetLogger().LogInfo("Config loaded successfully")
}

func GetJwtSecret() string {
	return GetAppConfig().Jwt.Secret
}

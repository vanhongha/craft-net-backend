package config

import (
	"craftnet/internal/util"

	"github.com/samber/lo"
	"github.com/spf13/viper"
)

type AwsConfig struct {
	Region string `mapstructure:"region"`
	Bucket string `mapstructure:"bucket"`
}

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
	Aws      AwsConfig      `mapstructure:"aws"`
	Jwt      JwtAuthConfig  `mapstructure:"jwt"`
	Database DatabaseConfig `mapstructure:"database"`
}

var AppConfig *Config

func GetAppConfig() *Config {
	if AppConfig == nil {
		LoadConfig("../../")
	}
	if AppConfig == nil {
		util.GetLogger().LogErrorWithMsg("Please load config first", true)
	}
	return AppConfig
}

func LoadConfig(path string) {
	if !lo.IsNil(AppConfig) {
		return
	}

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

	util.GetLogger().LogInfo("Config loaded successfully")
}

func GetJwtSecret() string {
	return GetAppConfig().Jwt.Secret
}

func GetAwsRegion() string {
	return GetAppConfig().Aws.Region
}

func GetAwsBucket() string {
	return GetAppConfig().Aws.Bucket
}

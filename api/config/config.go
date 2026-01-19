package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	App      AppConfig      `mapstructure:"app"`
	Database DatabaseConfig `mapstructure:"database"`
	Redis    RedisConfig    `mapstructure:"redis"`
}

type AppConfig struct {
	Env  string `mapstructure:"env"`
	Port int    `mapstructure:"port"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
	Charset  string `mapstructure:"charset"`
	Loc      string `mapstructure:"loc"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

var AppConfigData Config

func LoadConfig() {
	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "dev"
	}

	configName := "config"
	if env == "prod" {
		configName = "config.prod"
	}

	viper.SetConfigName(configName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config") // Optional: look in config dir

	// Environment variable overrides
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	if err := viper.Unmarshal(&AppConfigData); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	fmt.Printf("Config loaded: %s\n", configName)
}

package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
    Database    database    `mapstructure:",squash"`
    HTTP        http        `mapstructure:",squash"`
    AppEnv      string      `mapstructure:"app_env"`
}

type database struct {
    Domain     string      `mapstructure:"database_domain"`
    Name    string      `mapstructure:"database_name"`
}

type http struct {
    Port string     `mapstructure:"http_port"`
}

func New() (*Config, error) {
	viper.SetDefault("HTTP_PORT", "9999")
	viper.SetDefault("DATABASE_DOMAIN", "mongodb:mongo")
	viper.SetDefault("DATABASE_NAME", "user_database")
	viper.SetDefault("APP_ENV", "local")

	viper.SetConfigType("env")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Println("config: .env file not found")
	}

	cfg := &Config{}
	if err := viper.Unmarshal(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

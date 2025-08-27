package config

import "github.com/caarlos0/env/v11"

type AppConfig struct {
	JWTConfig
	DatabaseConfig
}

func NewConfig() (AppConfig, error) {
	config := AppConfig{}
	err := env.Parse(&config)
	return config, err
}

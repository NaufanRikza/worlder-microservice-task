package config

import "github.com/caarlos0/env/v11"

type Config struct {
	DatabaseConfig
	MqttConfig
}

func NewConfig() (Config, error) {
	config := Config{}
	err := env.Parse(&config)
	return config, err
}

package config

import (
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	JWTConfig
	DatabaseConfig
	MqttConfig
}

func NewConfig() (AppConfig, error) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := AppConfig{}
	err = env.Parse(&config)
	return config, err
}

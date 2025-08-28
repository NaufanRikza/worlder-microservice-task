package config

import (
	"github.com/caarlos0/env/v11"
)

type Config struct {
	MqttConfig
	AppConfig
	JWTConfig
}

func NewConfig() (Config, error) {

	// err := godotenv.Load("../.env")
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	config := Config{}
	err := env.Parse(&config)
	return config, err
}

package config

type JWTConfig struct {
	SecretKey string `env:"JWT_SECRET"`
}

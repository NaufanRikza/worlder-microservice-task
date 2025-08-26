package config

type AppConfig struct {
	DataGenerationFrequency uint `env:"DATA_GENERATION_FREQUENCY" envDefault:"1000"`
}

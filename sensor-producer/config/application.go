package config

type AppConfig struct {
	DataGenerationFrequency uint `env:"DATA_GENERATION_FREQUENCY"`
	ProducerHTTPPort        uint `env:"PRODUCER_HTTP_PORT"`
}

package config

type AppConfig struct {
	DataGenerationFrequency uint   `env:"DATA_GENERATION_FREQUENCY"`
	ProducerHTTPPort        uint   `env:"PRODUCER_HTTP_PORT"`
	SensorType              string `env:"SENSOR_TYPE"`
	SensorID                int    `env:"SENSOR_ID"`
}

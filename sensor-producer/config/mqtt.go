package config

import "fmt"

type MqttConfig struct {
	Host     string `env:"MQTT_HOST" envDefault:"localhost"`
	Port     string `env:"MQTT_PORT" envDefault:"1883"`
	User     string `env:"MQTT_USER"`
	Pass     string `env:"MQTT_PASS"`
	ClientID string `env:"MQTT_CLIENT_ID"`
	Topic    string `env:"MQTT_TOPIC"`
}

func (m MqttConfig) GetBrokerURL() string {
	return fmt.Sprintf(
		"tcp://%s:%s",
		m.Host, m.Port,
	)
}

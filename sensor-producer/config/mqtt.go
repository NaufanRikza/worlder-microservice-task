package config

import "fmt"

type MqttConfig struct {
	Host     string `env:"MQTT_HOST" envDefault:"localhost"`
	Port     string `env:"MQTT_PORT" envDefault:"1883"`
	User     string `env:"MQTT_USER" envDefault:"naufanrikza"`
	Pass     string `env:"MQTT_PASS" envDefault:"Tornado090699"`
	ClientID string `env:"MQTT_CLIENT_ID" envDefault:"sensor-producer"`
	Topic    string `env:"MQTT_TOPIC" envDefault:"sensor/data"`
}

func (m MqttConfig) GetBrokerURL() string {
	return fmt.Sprintf(
		"tcp://%s:%s",
		m.Host, m.Port,
	)
}

package cmd

import (
	"sensor-producer/config"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func StartMQTTClient(config config.MqttConfig) mqtt.Client {
	// Initialize MQTT client
	opts := mqtt.NewClientOptions()
	opts.AddBroker(config.GetBrokerURL()) // TCP MQTT
	opts.SetClientID(config.ClientID)     // Unique client ID
	opts.SetUsername(config.User)         // Username
	opts.SetPassword(config.Pass)         // Password

	mqttClient := mqtt.NewClient(opts) // Initialize client

	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	return mqttClient
}

package cmd

import (
	"fmt"
	"os"
	"sensor-consumer/config"
	"sensor-consumer/core/usecase"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func StartMQTTClient(config config.MqttConfig) mqtt.Client {
	// Initialize MQTT client
	opts := mqtt.NewClientOptions()
	opts.AddBroker(config.GetBrokerURL()) // TCP MQTT
	hostname, _ := os.Hostname()
	clientID := fmt.Sprintf("%s-%s", config.ClientID, hostname)
	fmt.Println("MQTT Client ID:", clientID)
	opts.SetClientID(clientID)    // Unique client ID
	opts.SetUsername(config.User) // Username
	opts.SetPassword(config.Pass) // Password

	mqttClient := mqtt.NewClient(opts) // Initialize client

	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	return mqttClient
}

func StartMQTTSubscriber(sensorUseCase usecase.SensorUsecase) error {
	return sensorUseCase.Subscribe()
}

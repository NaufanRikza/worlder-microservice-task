package cmd

import (
	"context"
	"encoding/json"
	"log"
	"math/rand"
	"sensor-producer/config"
	"sensor-producer/core/interface/dto"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func StartMQTTClient(config config.MqttConfig, freqChannel chan uint, ctx context.Context) {
	// Initialize MQTT client
	opts := mqtt.NewClientOptions()
	opts.AddBroker(config.Host)       // TCP MQTT
	opts.SetClientID(config.ClientID) // Unique client ID
	opts.SetUsername(config.User)     // Username
	opts.SetPassword(config.Pass)     // Password

	mqttClient := mqtt.NewClient(opts) // Initialize client

	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	GenerateSensorData(mqttClient, config.Topic, freqChannel, ctx)
}

func GenerateSensorData(mqttClient mqtt.Client, topic string, freqChannel chan uint, ctx context.Context) {
	// Generate and publish sensor data in certain timing
	ticker := time.NewTicker(1000 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Generate sensor data
			min := 10.0
			max := 100.0
			sensorData := dto.SensorData{
				SensorValue: rand.Float64()*(max-min) + min, // Random value between min and max
				ID1:         "T",
				ID2:         0,
				Timestamp:   time.Now().UTC().Format(time.RFC3339),
			}

			// Publish sensor data
			payload, err := json.Marshal(sensorData)
			if err != nil {
				log.Fatal(err)
			}

			token := mqttClient.Publish(topic, 0, false, payload)
			if token.Wait() && token.Error() != nil {
				log.Fatal(token.Error())
			}

		case freq := <-freqChannel:
			// Update the ticker frequency
			ticker.Stop()
			ticker = time.NewTicker(time.Duration(freq) * time.Millisecond)

		case <-ctx.Done():
			// Stop the MQTT client
			mqttClient.Disconnect(250)
			return
		}
	}
}

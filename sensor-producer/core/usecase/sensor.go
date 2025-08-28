package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/rand/v2"
	"sensor-producer/core/entity"
	"sensor-producer/core/infrastructure"
	"time"
)

type sensorUsecase struct {
	Publisher   infrastructure.Publisher
	Topic       string
	FreqChannel chan uint
	InitialFreq uint
}

type SensorUsecase interface {
	Start(ctx context.Context, sensorType string, sensorTypeName string, sensorID int)
	ChangeFrequency(freq uint) error
}

func NewSensorUsecase(publisher infrastructure.Publisher, topic string, initialFreq uint, freqChannel chan uint) SensorUsecase {
	return &sensorUsecase{
		Publisher:   publisher,
		Topic:       topic,
		InitialFreq: initialFreq,
		FreqChannel: freqChannel,
	}
}

func (s *sensorUsecase) Start(ctx context.Context, sensorType string, sensorTypeName string, sensorID int) {
	// Generate and publish sensor data in certain timing
	ticker := time.NewTicker(time.Duration(s.InitialFreq) * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Generate sensor data
			fmt.Println("Generating sensor data...")
			min := 10.0
			max := 100.0
			value := rand.Float64()*(max-min) + min
			value = math.Round(value*100) / 100

			sensorData := entity.SensorData{
				SensorValue: value,
				SensorType:  sensorTypeName,
				ID1:         sensorType,
				ID2:         sensorID,
				Timestamp:   time.Now().UTC().Format(time.RFC3339),
			}

			// Publish sensor data
			payload, err := json.Marshal(sensorData)
			if err != nil {
				log.Fatal(err)
			}

			s.Publisher.Publish(s.Topic, payload)

		case freq := <-s.FreqChannel:
			// Update the ticker frequency
			fmt.Println("Changing frequency to:", freq, "ms")
			ticker.Stop()
			ticker = time.NewTicker(time.Duration(freq) * time.Millisecond)

		case <-ctx.Done():
			fmt.Println("Stopping sensor data generation...")
			// Stop the MQTT client
			s.Publisher.Disconnect()
			return
		}
	}
}

func (s *sensorUsecase) ChangeFrequency(freq uint) error {
	select {
	case s.FreqChannel <- freq:
		return nil
	default:
		return fmt.Errorf("failed to change frequency")
	}
}

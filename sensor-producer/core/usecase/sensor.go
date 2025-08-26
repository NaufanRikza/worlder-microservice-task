package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
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
	Start(ctx context.Context)
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

func (s *sensorUsecase) Start(ctx context.Context) {
	// Generate and publish sensor data in certain timing
	ticker := time.NewTicker(time.Duration(s.InitialFreq) * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Generate sensor data
			min := 10.0
			max := 100.0
			sensorData := entity.SensorData{
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

			s.Publisher.Publish(s.Topic, payload)

		case freq := <-s.FreqChannel:
			// Update the ticker frequency
			ticker.Stop()
			ticker = time.NewTicker(time.Duration(freq) * time.Millisecond)

		case <-ctx.Done():
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

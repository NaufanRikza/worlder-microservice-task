package usecase

import (
	"context"
	"encoding/json"
	"log"
	"sensor-consumer/core/dto"
	"sensor-consumer/core/entity"
	mqtt_consumer "sensor-consumer/core/infrastructure/consumer"
	"sensor-consumer/core/repository"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type sensorUsecase struct {
	sensorRepo repository.SensorRepository
	consumer   mqtt_consumer.Consumer
}

type SensorUsecase interface {
	HandleMessage(client mqtt.Client, msg mqtt.Message)
	Subscribe() error
	GetSensorData(ctx context.Context, req dto.SensorRequest) ([]dto.SensorDataResult, error)
	DeleteSensorData(ctx context.Context, req dto.DeleteSensorRequest) error
	UpdateSensorData(ctx context.Context, req dto.UpdateSensorRequest, body dto.UpdateSensorBody) error
}

func NewSensorUsecase(sensorRepo repository.SensorRepository, consumer mqtt_consumer.Consumer) SensorUsecase {
	return &sensorUsecase{
		sensorRepo: sensorRepo,
		consumer:   consumer,
	}
}

func (s *sensorUsecase) GetSensorData(ctx context.Context, req dto.SensorRequest) ([]dto.SensorDataResult, error) {
	sensorData, err := s.sensorRepo.Get(ctx, req)
	return sensorData, err
}

func (s *sensorUsecase) DeleteSensorData(ctx context.Context, req dto.DeleteSensorRequest) error {
	return s.sensorRepo.Delete(ctx, req)
}

func (s *sensorUsecase) UpdateSensorData(ctx context.Context, req dto.UpdateSensorRequest, body dto.UpdateSensorBody) error {
	updatedData := entity.SensorData{
		SensorValue: body.SensorValue,
	}
	updatedData.ID = req.ID
	return s.sensorRepo.Update(ctx, req, updatedData)
}

func (s *sensorUsecase) HandleMessage(client mqtt.Client, msg mqtt.Message) {
	var sensorData entity.SensorData
	if err := json.Unmarshal(msg.Payload(), &sensorData); err != nil {
		log.Printf("Error unmarshalling sensor data: %v", err)
		return
	}

	ctx := context.Background()
	err := s.sensorRepo.Create(ctx, sensorData)
	if err != nil {
		log.Printf("Error saving sensor data: %v", err)
	}
}

func (s *sensorUsecase) Subscribe() error {
	return s.consumer.Consume(s.HandleMessage)
}

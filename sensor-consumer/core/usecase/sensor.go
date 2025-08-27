package usecase

import (
	"context"
	"sensor-consumer/core/entity"
	"sensor-consumer/core/repository"
)

type sensorUsecase struct {
	sensorRepo repository.SensorRepository
}

type SensorUsecase interface {
	GetSensorData(ctx context.Context, req repository.SensorRequest) (entity.SensorData, error)
	DeleteSensorData(ctx context.Context, id uint64) error
	UpdateSensorData(ctx context.Context, id uint64, body repository.UpdateSensorBody) error
}

func NewSensorUsecase(sensorRepo repository.SensorRepository) SensorUsecase {
	return &sensorUsecase{
		sensorRepo: sensorRepo,
	}
}

func (s *sensorUsecase) GetSensorData(ctx context.Context, req repository.SensorRequest) (entity.SensorData, error) {
	sensorData, err := s.sensorRepo.Get(ctx, req)
	return sensorData, err
}

func (s *sensorUsecase) DeleteSensorData(ctx context.Context, id uint64) error {
	return s.sensorRepo.Delete(ctx, id)
}

func (s *sensorUsecase) UpdateSensorData(ctx context.Context, id uint64, body repository.UpdateSensorBody) error {
	updatedData := entity.SensorData{
		SensorValue: body.SensorValue,
	}
	updatedData.ID = id
	return s.sensorRepo.Update(ctx, updatedData)
}

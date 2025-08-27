package usecase

import (
	"sensor-consumer/core/entity"
	"sensor-consumer/core/repository"
)

type sensorUsecase struct {
	sensorRepo repository.SensorRepository
}

type SensorUsecase interface {
	GetSensorData(repository.SensorRequest) (entity.SensorData, error)
	DeleteSensorData(id uint64) error
	UpdateSensorData(id uint64, body repository.UpdateSensorBody) error
}

func NewSensorUsecase(sensorRepo repository.SensorRepository) SensorUsecase {
	return &sensorUsecase{
		sensorRepo: sensorRepo,
	}
}

func (s *sensorUsecase) GetSensorData(req repository.SensorRequest) (entity.SensorData, error) {
	sensorData, err := s.sensorRepo.Get(req)
	return sensorData, err
}

func (s *sensorUsecase) DeleteSensorData(id uint64) error {
	return s.sensorRepo.Delete(id)
}

func (s *sensorUsecase) UpdateSensorData(id uint64, body repository.UpdateSensorBody) error {
	updatedData := entity.SensorData{
		SensorValue: body.SensorValue,
	}
	updatedData.ID = id
	return s.sensorRepo.Update(updatedData)
}

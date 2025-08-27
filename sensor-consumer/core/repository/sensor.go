package repository

import (
	"fmt"
	"sensor-consumer/core/entity"

	"gorm.io/gorm"
)

type sensorRepository struct {
	DB *gorm.DB
}

type SensorRepository interface {
	Get(SensorRequest) (entity.SensorData, error)
	Delete(id uint64) error
	Update(sensor entity.SensorData) error
}

func NewSensorRepository(db *gorm.DB) SensorRepository {
	return &sensorRepository{
		DB: db,
	}
}

func (r *sensorRepository) Get(req SensorRequest) (entity.SensorData, error) {
	var sensor entity.SensorData
	db := r.DB.Model(&sensor).Table(sensor.TableName())
	db = db.Where("id1 = ?", req.ID1)
	db = db.Where("id2 = ?", req.ID2)
	if req.TimeStart != nil && !req.TimeStart.IsZero() {
		if req.TimeEnd != nil && !req.TimeEnd.IsZero() {
			db.Where("timestamp BETWEEN ? AND ?", req.TimeStart, req.TimeEnd)
		} else {
			db.Where("timestamp >= ?", req.TimeStart)
		}
	}

	offset := int((req.Page - 1) * req.Limit)
	orderby := fmt.Sprintf("%s %s", req.Sort, req.Order)

	err := db.Order(orderby).Offset(offset).Limit(int(req.Limit)).Find(&sensor).Error

	return sensor, err
}

func (r *sensorRepository) Delete(id uint64) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		return tx.Model(&entity.SensorData{}).Table(entity.SensorData{}.TableName()).Delete(&entity.SensorData{}, id).Error
	})
}

func (r *sensorRepository) Update(sensor entity.SensorData) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		return tx.Model(&sensor).Table(sensor.TableName()).Updates(sensor).Error
	})
}

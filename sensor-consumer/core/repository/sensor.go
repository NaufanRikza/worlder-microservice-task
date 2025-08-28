package repository

import (
	"context"
	"fmt"
	"sensor-consumer/core/dto"
	"sensor-consumer/core/entity"

	"gorm.io/gorm"
)

type sensorRepository struct {
	DB *gorm.DB
}

type SensorRepository interface {
	Get(ctx context.Context, req dto.SensorRequest) ([]entity.SensorData, error)
	Delete(ctx context.Context, id uint64) error
	Update(ctx context.Context, sensor entity.SensorData) error
	Create(ctx context.Context, sensor entity.SensorData) error
}

func NewSensorRepository(db *gorm.DB) SensorRepository {
	return &sensorRepository{
		DB: db,
	}
}

func (r *sensorRepository) Get(ctx context.Context, req dto.SensorRequest) ([]entity.SensorData, error) {
	var sensor []entity.SensorData
	db := r.DB.WithContext(ctx).Model(&sensor).Table(entity.SensorData{}.TableName())
	db = db.Where("id1 = ?", req.ID1)
	db = db.Where("id2 = ?", req.ID2)
	if req.TimeStart != nil && !req.TimeStart.IsZero() {
		if req.TimeEnd != nil && !req.TimeEnd.IsZero() {
			db.Where("timestamp BETWEEN ? AND ?", req.TimeStart, req.TimeEnd)
		} else {
			db.Where("timestamp >= ?", req.TimeStart)
		}
	}

	offset := int((req.Page - 1) * req.Length)
	orderby := fmt.Sprintf("%s %s", req.Sort, req.Order)

	err := db.Order(orderby).Offset(offset).Limit(int(req.Length)).Find(&sensor).Error

	return sensor, err
}

func (r *sensorRepository) Delete(ctx context.Context, id uint64) error {
	return r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Model(&entity.SensorData{}).Table(entity.SensorData{}.TableName()).Delete(&entity.SensorData{}, id).Error
	})
}

func (r *sensorRepository) Update(ctx context.Context, sensor entity.SensorData) error {
	return r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Model(&sensor).Table(sensor.TableName()).Updates(sensor).Error
	})
}

func (r *sensorRepository) Create(ctx context.Context, sensor entity.SensorData) error {
	return r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Model(&sensor).Table(sensor.TableName()).Create(&sensor).Error
	})
}

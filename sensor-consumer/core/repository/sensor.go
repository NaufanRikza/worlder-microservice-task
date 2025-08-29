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
	Get(ctx context.Context, req dto.SensorRequest) ([]dto.SensorDataResult, error)
	Delete(ctx context.Context, req dto.DeleteSensorRequest) error
	Update(ctx context.Context, req dto.UpdateSensorRequest, sensor entity.SensorData) error
	Create(ctx context.Context, sensor entity.SensorData) error
}

func NewSensorRepository(db *gorm.DB) SensorRepository {
	return &sensorRepository{
		DB: db,
	}
}

func (r *sensorRepository) Get(ctx context.Context, req dto.SensorRequest) ([]dto.SensorDataResult, error) {
	var sensor []dto.SensorDataResult
	db := r.DB.Debug().
		WithContext(ctx).
		Select("sensor_value, sensor_type, id1, id2, timestamp").
		Model(&sensor).
		Table(entity.SensorData{}.TableName())

	if req.TimeStart != nil && !req.TimeStart.IsZero() {
		if req.TimeEnd != nil && !req.TimeEnd.IsZero() {
			db = db.Where("timestamp BETWEEN ? AND ?", req.TimeStart, req.TimeEnd)
		} else {
			db = db.Where("timestamp >= ?", req.TimeStart)
		}
	}

	if req.ID1 != "" {
		db = db.Where("id1 = ?", req.ID1)
	}

	if req.ID2 > 0 {
		db = db.Where("id2 = ?", req.ID2)
	}

	offset := int((req.Page - 1) * req.Length)
	orderby := fmt.Sprintf("%s %s", req.Sort, req.Order)

	err := db.Order(orderby).Offset(offset).Limit(int(req.Length)).Find(&sensor).Error

	return sensor, err
}

func (r *sensorRepository) Delete(ctx context.Context, req dto.DeleteSensorRequest) error {
	return r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		db := tx.Debug().Model(&entity.SensorData{}).
			Table(entity.SensorData{}.TableName())

		if req.ID != 0 {
			db = db.Where("id = ?", req.ID)
		}

		if req.TimeStart != nil && !req.TimeStart.IsZero() {
			if req.TimeEnd != nil && !req.TimeEnd.IsZero() {
				db = db.Where("timestamp BETWEEN ? AND ?", req.TimeStart, req.TimeEnd)
			} else {
				db = db.Where("timestamp >= ?", req.TimeStart)
			}
		}

		if req.ID1 != "" {
			db = db.Where("id1 = ?", req.ID1)
		}

		if req.ID2 > 0 {
			db = db.Where("id2 = ?", req.ID2)
		}

		return db.Delete(&entity.SensorData{}).Error
	})
}

func (r *sensorRepository) Update(ctx context.Context, req dto.UpdateSensorRequest, sensor entity.SensorData) error {
	return r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		db := tx.Debug().Model(&sensor).Table(sensor.TableName())

		if req.ID != 0 {
			db = db.Where("id = ?", req.ID)
		}

		if req.TimeStart != nil && !req.TimeStart.IsZero() {
			if req.TimeEnd != nil && !req.TimeEnd.IsZero() {
				db = db.Where("timestamp BETWEEN ? AND ?", req.TimeStart, req.TimeEnd)
			} else {
				db = db.Where("timestamp >= ?", req.TimeStart)
			}
		}
		if req.ID1 != "" {
			db = db.Where("id1 = ?", req.ID1)
		}

		if req.ID2 > 0 {
			db = db.Where("id2 = ?", req.ID2)
		}

		return db.Updates(&sensor).Error
	})
}

func (r *sensorRepository) Create(ctx context.Context, sensor entity.SensorData) error {
	return r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Model(&sensor).Table(sensor.TableName()).Create(&sensor).Error
	})
}

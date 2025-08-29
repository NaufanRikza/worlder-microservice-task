package entity

import (
	"time"

	"gorm.io/gorm"
)

type SensorData struct {
	DefaultAttribute
	SensorValue float64   `json:"sensor_value" gorm:"column:sensor_value"`
	SensorType  string    `json:"sensor_type" gorm:"column:sensor_type"`
	ID1         string    `json:"id1" gorm:"column:id1"`
	ID2         int       `json:"id2" gorm:"column:id2"`
	Timestamp   time.Time `json:"timestamp" gorm:"column:timestamp"`
}

func (SensorData) TableName() string {
	return "sensor_data"
}

func (s *SensorData) BeforeCreate(tx *gorm.DB) (err error) {
	s.Timestamp = s.Timestamp.UTC()
	return
}

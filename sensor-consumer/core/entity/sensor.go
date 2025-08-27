package entity

import "time"

type SensorData struct {
	DefaultAttribute
	SensorValue float64   `json:"sensor_value" gorm:"column:sensor_value"`
	ID1         string    `json:"id1" gorm:"column:id1"`
	ID2         int       `json:"id2" gorm:"column:id2"`
	Timestamp   time.Time `json:"timestamp" gorm:"column:timestamp"`
}

func (SensorData) TableName() string {
	return "sensor_data"
}

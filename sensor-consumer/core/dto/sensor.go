package dto

import "time"

type SensorRequest struct {
	Page      uint       `json:"page" query:"page" validate:"required"`
	Length    uint       `json:"length" query:"length" validate:"required"`
	Sort      string     `json:"sort" query:"sort" validate:"required"`
	Order     string     `json:"order" query:"order" validate:"required"`
	ID1       string     `json:"id1" query:"id1"`
	ID2       uint       `json:"id2" query:"id2"`
	TimeStart *time.Time `json:"time_start" query:"time_start" time_format:"2006-01-02T15:04:05Z07:00"`
	TimeEnd   *time.Time `json:"time_end" query:"time_end" time_format:"2006-01-02T15:04:05Z07:00"`
}

type UpdateSensorRequest struct {
	ID1       string     `json:"id1" query:"id1"`
	ID2       uint64     `json:"id2" query:"id2"`
	ID        uint64     `json:"id" query:"id"`
	TimeStart *time.Time `json:"time_start" query:"time_start" time_format:"2006-01-02T15:04:05Z07:00"`
	TimeEnd   *time.Time `json:"time_end" query:"time_end" time_format:"2006-01-02T15:04:05Z07:00"`
}

type UpdateSensorBody struct {
	SensorValue float64 `json:"sensor_value" validate:"required,gt=0"`
}

type DeleteSensorRequest struct {
	ID1       string     `json:"id1" query:"id1"`
	ID2       uint64     `json:"id2" query:"id2"`
	ID        uint64     `json:"id" query:"id"`
	TimeStart *time.Time `json:"time_start" query:"time_start" time_format:"2006-01-02T15:04:05Z07:00"`
	TimeEnd   *time.Time `json:"time_end" query:"time_end" time_format:"2006-01-02T15:04:05Z07:00"`
}

type SensorDataResult struct {
	SensorValue float64   `json:"sensor_value"`
	ID1         string    `json:"id1"`
	ID2         int       `json:"id2"`
	Timestamp   time.Time `json:"timestamp"`
}

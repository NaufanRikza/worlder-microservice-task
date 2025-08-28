package dto

import "time"

type SensorRequest struct {
	Page      uint       `json:"page" query:"page" validate:"required"`
	Length    uint       `json:"length" query:"length" validate:"required"`
	Limit     uint       `json:"limit" query:"limit" validate:"required"`
	Sort      string     `json:"sort" query:"sort" validate:"required"`
	Order     string     `json:"order" query:"order" validate:"required"`
	ID1       uint       `json:"id1" query:"id1" validate:"required"`
	ID2       uint       `json:"id2" query:"id2"`
	TimeStart *time.Time `json:"time_start" query:"time_start" time_format:"2006-01-02T15:04:05Z07:00"`
	TimeEnd   *time.Time `json:"time_end" query:"time_end" time_format:"2006-01-02T15:04:05Z07:00"`
}

type UpdateSensorBody struct {
	SensorValue float64 `json:"sensor_value" validate:"required"`
}

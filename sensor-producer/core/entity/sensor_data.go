package entity

type SensorData struct {
	SensorValue float64 `json:"sensor_value"`
	SensorType  string  `json:"sensor_type"`
	ID1         string  `json:"id1"`
	ID2         int     `json:"id2"`
	Timestamp   string  `json:"timestamp"`
}

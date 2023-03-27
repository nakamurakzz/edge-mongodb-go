package sensor

import "time"

type SensorData struct {
	Temperature  float64 `json:"temperature"`
	Humidity     float64 `json:"humidity"`
	timeStamp    int64
	isPublished  bool
	isPublishing bool
}

func NewSensorData(temperature float64, humidity float64) *SensorData {
	timeStamp := time.Now().Unix()

	return &SensorData{
		Temperature:  temperature,
		Humidity:     humidity,
		timeStamp:    timeStamp,
		isPublished:  false,
		isPublishing: false,
	}
}

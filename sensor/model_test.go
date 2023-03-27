package sensor

import "testing"

func TestNewSensorData(t *testing.T) {
	temperature := 20.0
	humidity := 30.0

	sensorData := NewSensorData(temperature, humidity)

	if sensorData.Temperature != temperature {
		t.Errorf("temperature is not same: %f", sensorData.Temperature)
	}

	if sensorData.Humidity != humidity {
		t.Errorf("humidity is not same: %f", sensorData.Humidity)
	}

	if sensorData.isPublished != false {
		t.Errorf("isPublished is not false: %t", sensorData.isPublished)
	}

	if sensorData.isPublishing != false {
		t.Errorf("isPublishing is not false: %t", sensorData.isPublished)
	}
}

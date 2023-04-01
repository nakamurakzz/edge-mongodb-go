package sensor

import (
	"encoding/csv"
	"os"
	"strconv"
)

type FileRepository struct {
}

func NewFileRepository() *FileRepository {
	return &FileRepository{}
}

func (f *FileRepository) GetSensorData() (sensorData *[]SensorData, err error) {
	filePath := "./sensor_data.csv"

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	csvReader.Comma = ','

	var sensorDataList []SensorData
	for {
		record, err := csvReader.Read()
		if err != nil {
			break
		}

		temperature, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			return nil, err
		}
		humidity, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			return nil, err
		}
		sensorDataList = append(sensorDataList, SensorData{
			Temperature: temperature,
			Humidity:    humidity,
		})
	}

	return &sensorDataList, nil
}

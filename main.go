package main

import (
	"flag"
	"log"
	"time"

	"github.com/nakamurakzz/edge-mongodb-go/sensor"
)

func main() {
	log.Println("start main")

	flag.Parse()

	sensorUsecase := sensor.NewSensorUsecase()
	defer sensorUsecase.MongoDB.Close()
	sensorDataRepository := sensor.NewFileRepository()

	// 30秒おきにsensorUsecase.Registerを実行する
	ticker := time.Tick(2 * time.Second)

	for range ticker {
		// sensorDataを取得
		sensorData, err := sensorDataRepository.GetSensorData()
		if err != nil {
			log.Println(err.Error())
			continue
		}
		// sensorDataをローカルに登録
		sensorUsecase.Register(sensorData)

		// sensorDataをpublish
		sensorUsecase.Publish()
	}

	log.Println("end main")
}

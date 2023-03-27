package sensor

import (
	"log"
)

type SensorUsecase struct {
	mongoDB *Mongodb
}

func NewSensorUsecase() *SensorUsecase {
	// mongodbに接続
	mongoDB, err := NewMongodb()
	if err != nil {
		panic(err)
	}
	return &SensorUsecase{
		mongoDB: mongoDB,
	}
}

func (s *SensorUsecase) Register(temperature *float64, humidity *float64) {
	log.Println("start register")

	// documentを作成
	document := NewSensorData(*temperature, *humidity)
	err := s.mongoDB.CreateDocument(*document)
	if err != nil {
		panic(err)
	}
	// TODO: deferはここでいい？
	defer s.mongoDB.Close()

	log.Println("end register")
}

func (s *SensorUsecase) Publish() {
	log.Println("start publish")
	// documentを取得
	// TimeStreamにpublish
	// documentを削除

	log.Println("end publish")
}

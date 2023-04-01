package sensor

import (
	"log"
)

type SensorUsecase struct {
	MongoDB *Mongodb
}

func NewSensorUsecase() *SensorUsecase {
	// mongodbに接続
	mongoDB, err := NewMongodb()
	if err != nil {
		panic(err)
	}
	return &SensorUsecase{
		MongoDB: mongoDB,
	}
}

func (s *SensorUsecase) Register(sensorData *[]SensorData) {
	log.Println("start register")

	// documentを作成
	// TODO: N回ループしないように修正する
	for _, data := range *sensorData {
		document := NewSensorData(data.Temperature, data.Humidity)
		err := s.MongoDB.CreateDocument(*document)
		if err != nil {
			panic(err)
		}
	}

	log.Println("end register")
}

func (s *SensorUsecase) Publish() {
	log.Println("start publish")
	// documentを取得
	// isPublishがfalseのものを取得
	document, err := s.MongoDB.FindDocument(SensorDataFindQuery{
		IsPublish: false,
	})

	if err != nil {
		panic(err)
	}

	// TimeStreamにpublish

	// documentを削除
	err = s.MongoDB.DeleteDocument(document)
	if err != nil {
		panic(err)
	}
	log.Println("end publish")
}

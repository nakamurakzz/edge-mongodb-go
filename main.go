package main

import (
	"flag"

	"github.com/nakamurakzz/edge-mongodb-go/sensor"
)

func main() {

	// フラグ変数を定義します
	temperature := flag.Float64("temperature", 0, "Temperature value")
	humidity := flag.Float64("humidity", 0, "Humidity value")

	// コマンドライン引数を解析します
	flag.Parse()

	// mongodbに接続
	mongoDB, err := sensor.NewMongodb()
	if err != nil {
		panic(err)
	}
	defer mongoDB.Close()

	// documentを作成
	document := sensor.NewSensorData(*temperature, *humidity)
	err = mongoDB.CreateDocument(*document)
}

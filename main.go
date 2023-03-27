package main

import (
	"flag"
	"log"

	"github.com/nakamurakzz/edge-mongodb-go/sensor"
)

func main() {
	log.Println("start main")

	// register or publish
	mode := flag.String("mode", "none", "reg or pub")

	// TODO: Validation
	temperature := flag.Float64("temperature", 0, "Temperature value")
	humidity := flag.Float64("humidity", 0, "Humidity value")

	flag.Parse()

	sensorUsecase := sensor.NewSensorUsecase()

	switch *mode {
	case "reg":
		sensorUsecase.Register(temperature, humidity)
	case "pub":
		sensorUsecase.Publish()
	default:
		log.Println("mode is not reg or pub")
	}

	log.Println("end main")
}

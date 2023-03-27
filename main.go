package main

import (
	"github.com/nakamurakzz/edge-mongodb-go/repository"
)

func main() {
	// mongodbに接続
	mongoDB, err := repository.NewMongodb()
	if err != nil {
		panic(err)
	}
	defer mongoDB.Close()

	// documentを作成
	err = mongoDB.CreateDocument("test1", repository.Content{
		Temperature: 20.0,
		Humidity:    30.0,
	})
}

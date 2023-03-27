package repository

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongodb struct {
	Client *mongo.Client
	dbName string
}

// シングルトンとして実装する
var mongodb *Mongodb

func NewMongodb() (*Mongodb, error) {
	if mongodb != nil {
		return mongodb, nil
	}

	uri := "mongodb://localhost:27017"

	log.Println("mongodb connect")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	mongodb = &Mongodb{
		Client: client,
		dbName: "db2",
	}

	return mongodb, nil
}

type Content struct {
	Temperature float64
	Humidity    float64
}

func (m *Mongodb) CreateDocument(collection string, document Content) error {
	log.Println("create document")
	_, err := m.Client.Database(m.dbName).Collection(collection).InsertOne(context.Background(), document)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (m *Mongodb) Close() {
	log.Println("mongodb close")
	m.Client.Disconnect(context.Background())
}

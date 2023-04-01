package sensor

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SensorDataFindQuery struct {
	IsPublish bool
}

type Mongodb struct {
	Client     *mongo.Client
	dbName     string
	collection string
}

var mongodb *Mongodb

func NewMongodb(uri string) (*Mongodb, error) {
	if mongodb != nil {
		return mongodb, nil
	}

	log.Println("mongodb connect")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	mongodb = &Mongodb{
		Client:     client,
		dbName:     "sensor_db",
		collection: "sensor",
	}

	return mongodb, nil
}

func (m *Mongodb) CreateDocument(document SensorData) error {
	log.Println("create document")

	log.Println(document)
	_, err := m.Client.Database(m.dbName).Collection(m.collection).InsertOne(context.Background(), document)

	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (m *Mongodb) FindDocument(query SensorDataFindQuery) (*[]SensorData, error) {
	log.Println("find document")

	documents, err := m.Client.Database(m.dbName).Collection(m.collection).Find(context.Background(), query)

	fmt.Println("------------------")
	fmt.Println(*documents)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	var result []SensorData
	for documents.Next(context.Background()) {
		var document SensorData
		err := documents.Decode(&document)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		result = append(result, document)
	}

	return &result, nil
}

func (m *Mongodb) DeleteDocument(document *[]SensorData) error {
	log.Println("delete document")

	for _, data := range *document {
		_, err := m.Client.Database(m.dbName).Collection(m.collection).DeleteOne(context.Background(), data)
		if err != nil {
			log.Println(err.Error())
			return err
		}
	}

	return nil
}

func (m *Mongodb) UpdateDocument(document *[]SensorData) error {
	log.Println("update document")

	for _, data := range *document {
		_, err := m.Client.Database(m.dbName).Collection(m.collection).UpdateOne(context.Background(), data, data)
		if err != nil {
			log.Println(err.Error())
			return err
		}
	}

	return nil
}

func (m *Mongodb) Close() {
	log.Println("mongodb close")
	m.Client.Disconnect(context.Background())
}

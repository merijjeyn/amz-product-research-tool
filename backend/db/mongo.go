package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var database *mongo.Database
var client *mongo.Client

type MongoEntry struct {
	Key  string
	Data map[string]interface{}
}

func ConnectToMongoDB() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(fmt.Errorf("db/mongo.connect: error connecting to mongo db:\n%v", err))
	}

	database = client.Database("amz_prod_res")
}

func DisconnectFromMongoDB() {
	err := client.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
}

func SaveData_MDB(collectionName string, key string, data map[string]interface{}) error {
	if database == nil {
		fmt.Println("db nil")
	}
	collection := database.Collection(collectionName)

	dataToInsert := MongoEntry{
		Key:  key,
		Data: data,
	}

	_, err := collection.InsertOne(context.Background(), dataToInsert)
	if err != nil {
		return fmt.Errorf("db/mongo.SaveData_MDB: error inserting data into entry:\n%v", err)
	}

	return nil
}

package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBSet() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://154.208.140.222:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err1 := client.Connect(ctx)
	if err1 != nil {
		log.Fatal(err1)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println("Failed to connect to MongoDB :(")
		return nil
	}
	fmt.Println("Successfully connected to MongoDB.")
	return client
}

var client *mongo.Client = DBSet()

func ImageData() *mongo.Collection {
	var productCollection *mongo.Collection = client.Database("image-data").Collection("data")
	return productCollection
}

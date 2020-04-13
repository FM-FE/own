package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const IP = "localhost"
const mongoURL = "mongodb://" + IP + ":27017"

func GetClient() (client *mongo.Client) {
	clientOptions := options.Client().ApplyURI(mongoURL)
	client, e := mongo.Connect(context.TODO(), clientOptions)
	if e != nil {
		fmt.Print("Failed to connect mongo!")
		fmt.Print(e.Error())
		return
	}
	return client
}

func CreateAndInsert() {
	client := GetClient()
	collection := client.Database("test-database").Collection("test-collection")
	fmt.Print("make database and collection")
	collection.InsertOne(context.Background(), bson.D{
		{"item", "canvas"},
		{"qty", 100},
		{"tags", bson.A{"cotton"}},
		{"size", bson.D{
			{"h", 28},
			{"w", 35.5},
			{"uom", "cm"},
		}},
	})

}

func main() {
	CreateAndInsert()
}

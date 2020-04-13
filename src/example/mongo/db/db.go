package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

const IP = "localhost"
const mongoURL = "mongodb://" + IP + ":27017"

func GetClient(url string) (client *mongo.Client) {
	var clientOptions *options.ClientOptions
	if url == "" {
		clientOptions = options.Client().ApplyURI(mongoURL)
	} else {
		clientOptions = options.Client().ApplyURI(url)
	}

	client, e := mongo.Connect(context.TODO(), clientOptions)
	if e != nil {
		log.Println("Failed to connect mongo!")
		log.Println(e.Error())
		return
	}
	return client
}

func GetDatabase(url string, db string) (database *mongo.Database) {
	return GetClient(url).Database(db)
}

func GetCollection(url string, db string, coll string) (collection *mongo.Collection) {
	return GetClient(url).Database(db).Collection(coll)
}

func CreateAndInsert(w http.ResponseWriter, r *http.Request) {
	client := GetClient("")
	collection := client.Database("testdatabase").Collection("testcollection")
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

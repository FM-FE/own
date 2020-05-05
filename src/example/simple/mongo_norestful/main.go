package mongo_norestful

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

//
//const IP = "localhost"
//const mongoURL = "mongodb://" + IP + ":27017"
//
//func GetClient() (client *mongo.Client) {
//	clientOptions := options.Client().ApplyURI(mongoURL)
//	client, e := mongo.Connect(context.TODO(), clientOptions)
//	if e != nil {
//		fmt.Print("Failed to connect mongo!")
//		fmt.Print(e.Error())
//		return
//	}
//	return client
//}
//
//func CreateAndInsert() {
//	client := GetClient()
//	collection := client.Database("test-database").Collection("test-collection")
//	fmt.Print("make database and collection")
//	collection.InsertOne(context.Background(), bson.D{
//		{"item", "canvas"},
//		{"qty", 100},
//		{"tags", bson.A{"cotton"}},
//		{"size", bson.D{
//			{"h", 28},
//			{"w", 35.5},
//			{"uom", "cm"},
//		}},
//	})
//
//}
//

type test struct {
	Name string `json:"name" bson:"-"`
	Age  string `json:"-" bson:"age"`
}

func main() {
	log.Println("hello world")
	// CreateAndInsert()
	test := test{
		Name: "test",
		Age:  "23",
	}
	json_out, _ := json.Marshal(test)
	log.Println(string(json_out))

	bson_out, _ := bson.Marshal(test)
	log.Println(string(bson_out))

}

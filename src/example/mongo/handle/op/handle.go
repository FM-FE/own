package op

import (
	"context"
	"encoding/json"
	"example/mongo/db"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"time"
)

type Operation struct {
	Name        string    `json:"name" bson:"name"`
	Description string    `json:"description" bson:"description"`
	Step        []string  `json:"step" bson:"step"`
	Time        time.Time `json:"time" bson:"time"`
	Frequency   Frequency `json:"frequency" bson:"frequency"`

	ProgressBar float32 `json:"progress_bar" bson:"progress_bar"`
	Achieved    bool    `json:"achieved" bson:"achieved"`

	Weight float32 `json:"weight" bson:"weight"` // from 0 - 100

	Atom bool `json:"atom" bson:"atom"`
}

type Frequency struct {
	Oneshot   bool   `json:"oneshot" bson:"oneshot"`
	Frequency string `json:"frequency" bson:"Frequency"`
}

type CommonResponse struct {
	Result string `json:"result"`
	Error  string `json:"error"`
}

type InsertOperationResponse struct {
	CommonResponse
	InsertOneID primitive.ObjectID `json:"insert_one_id"`
}

func InsertOperation(w http.ResponseWriter, r *http.Request) {
	log.Println("in CreateOperation")
	var rsp InsertOperationResponse
	defer func() {
		buf, e := json.Marshal(&rsp)
		if e != nil {
			w.WriteHeader(500)
		}
		w.Write([]byte(buf))
	}()

	client := db.GetClient("")
	c := client.Database("test_db").Collection("test_coll")
	log.Println("Get Collection")

	// operation should be decoded from request
	operation := Operation{
		Name:        "op",
		Description: "des",
		Step:        []string{"first", "second"},
		Time:        time.Now(),
		Frequency: Frequency{
			Oneshot:   false,
			Frequency: "monthly",
		},
		ProgressBar: 63.2,
		Achieved:    false,
		Weight:      94.3,
		Atom:        false,
	}

	insertOneResult, e := c.InsertOne(context.TODO(), operation)
	if e != nil {
		log.Println(e.Error())
		rsp.Result = "ERROR"
		rsp.Error = e.Error()
		return
	}
	log.Println(insertOneResult.InsertedID)

	rsp.Result = "OK"
	rsp.InsertOneID = insertOneResult.InsertedID.(primitive.ObjectID)
}

func UpdateOperation(w http.ResponseWriter, r *http.Request) {
	client := db.GetClient("")
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

func ReadOperation(w http.ResponseWriter, r *http.Request) {
	client := db.GetClient("")
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

func DeleteOperation(w http.ResponseWriter, r *http.Request) {
	client := db.GetClient("")
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

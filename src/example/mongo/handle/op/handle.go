package op

import (
	"context"
	"encoding/json"
	"example/mongo/db"
	"example/mongo/handle/op/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"time"
)

type Operation struct {
	Name        string    `json:"name,omitempty" bson:"name,omitempty"`
	Description string    `json:"description,omitempty" bson:"description,omitempty"`
	Step        []string  `json:"step,omitempty" bson:"step,omitempty"`
	Time        time.Time `json:"time,omitempty" bson:"time,omitempty"`
	Frequency   Frequency `json:"frequency,omitempty" bson:"frequency,omitempty"`

	ProgressBar float32 `json:"progress_bar,omitempty" bson:"progress_bar,omitempty"`
	Achieved    bool    `json:"achieved,omitempty" bson:"achieved,omitempty"`

	Weight float32 `json:"weight,omitempty" bson:"weight,omitempty"` // from 0 - 100

	Atom bool `json:"atom,omitempty" bson:"atom,omitempty"`
}

type Frequency struct {
	Oneshot   bool   `json:"oneshot,omitempty" bson:"oneshot,omitempty"`
	Frequency string `json:"frequency,omitempty" bson:"Frequency,omitempty"`
}

type CommonResponse struct {
	Result string `json:"result"`
	Error  string `json:"error,omitempty"`
}

// Insert
type InsertOperationResponse struct {
	CommonResponse
	InsertOneID primitive.ObjectID `json:"insert_one_id"`
}

// Find
type FindOperationRequest struct {
	Operation Operation
	Limit     int64 `json:"limit" bson:"-"`
}

type FindOperationResponse struct {
	CommonResponse
	Operations []Operation `json:"operations"`
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
		utils.ErrorToResponse(&rsp.CommonResponse, e)
		return
	}
	log.Println(insertOneResult.InsertedID)

	rsp.Result = "OK"
	rsp.InsertOneID = insertOneResult.InsertedID.(primitive.ObjectID)
}

func UpdateOperation(w http.ResponseWriter, r *http.Request) {
	log.Println("in UpdateOperation")
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
	log.Println(c)

}

func FindOperation(w http.ResponseWriter, r *http.Request) {
	log.Println("in UpdateOperation")
	var rsp FindOperationResponse
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
	log.Println(c)

	reqbody, e := utils.AnalyzeRequest(r)
	if e != nil {
		utils.ErrorToResponse(&rsp.CommonResponse, e)
		return
	}

	var req FindOperationRequest
	filter, e := utils.JsonToBson(reqbody, req)
	if e != nil {
		utils.ErrorToResponse(&rsp.CommonResponse, e)
		return
	}

	findOptions := options.Find()
	if req.Limit != -1 {
		findOptions.SetLimit(req.Limit)
	}
	cursor, e := c.Find(context.Background(), filter, findOptions)
	if e != nil {
		utils.ErrorToResponse(&rsp.CommonResponse, e)
		return
	}

	for cursor.Next(context.Background()) {
		var operation Operation
		e = cursor.Decode(&operation)
		if e != nil {
			utils.ErrorToResponse(&rsp.CommonResponse, e)
			return
		}
		log.Println(operation)
		rsp.Operations = append(rsp.Operations, operation)
	}
	if err := cursor.Err(); err != nil {
		utils.ErrorToResponse(&rsp.CommonResponse, e)
		return
	}
	cursor.Close(context.Background())
	rsp.Result = "OK"
}

func DeleteOperation(w http.ResponseWriter, r *http.Request) {
	log.Println("in UpdateOperation")
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
	log.Println(c)

}

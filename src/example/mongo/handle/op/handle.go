package op

import (
	"context"
	"encoding/json"
	"example/mongo/db"
	"example/mongo/handle/op/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"time"
)

// Insert
type InsertOperationResponse struct {
	utils.CommonResponse
	InsertOneID primitive.ObjectID `json:"insert_one_id"`
}

// Find
type FindOperationRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type FindOperationResponse struct {
	utils.CommonResponse
	Operations []utils.Operation `json:"operations"`
}

// Update
type UpdateOperationRequest struct {
	Filter struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"filter"`
	Update struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"update"`
}

// Delete
type DeleteOperationRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func InsertOperation(w http.ResponseWriter, r *http.Request) {
	log.Println("in InsertOperation")
	log.Println(">> This is a {!most!} new test")
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
	op := utils.Operation{
		Name:        "op",
		Description: "des",
		Step:        []string{"first", "second"},
		Time:        time.Now(),
		Frequency: utils.Frequency{
			Oneshot:   false,
			Frequency: "monthly",
		},
		ProgressBar: 63.2,
		Achieved:    false,
		Weight:      94.3,
		Atom:        false,
	}

	insertOneResult, e := c.InsertOne(context.TODO(), op)
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

	var req UpdateOperationRequest
	e := utils.StructRequest(r, &req)
	if e != nil {
		log.Println("ERROR")
		utils.ErrorToResponse(&rsp.CommonResponse, e)
		return
	}

	result, e := c.UpdateMany(
		context.Background(),
		bson.D{
			{req.Filter.Key, req.Filter.Value},
		},
		bson.D{
			{"$set", bson.D{
				{req.Update.Key, req.Update.Value},
			}},
		},
	)
	if e != nil {
		log.Println("ERROR")
		utils.ErrorToResponse(&rsp.CommonResponse, e)
		return
	}
	log.Println(result)
	rsp.Result = "OK"

}

func UpdateOneOperation(w http.ResponseWriter, r *http.Request) {
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

	var req UpdateOperationRequest
	e := utils.StructRequest(r, &req)
	if e != nil {
		log.Println("ERROR")
		utils.ErrorToResponse(&rsp.CommonResponse, e)
		return
	}
	log.Printf("req is %+v", req)

	result, e := c.UpdateOne(
		context.Background(),
		bson.D{
			{req.Filter.Key, req.Filter.Value}},
		bson.D{
			{"$set", bson.D{
				{req.Update.Key, req.Update.Value},
			}},
		},
	)
	if e != nil {
		log.Println("ERROR")
		utils.ErrorToResponse(&rsp.CommonResponse, e)
		return
	}
	log.Printf("update result is: %+v", result)
	rsp.Result = "OK"
}

func FindOperation(w http.ResponseWriter, r *http.Request) {
	log.Println("in FindOperation")
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

	reqbody, e := utils.AnalyzeRequest(r)
	if e != nil {
		utils.ErrorToResponse(&rsp.CommonResponse, e)
		return
	}

	var req FindOperationRequest
	e = json.Unmarshal(reqbody, &req)
	if e != nil {
		log.Println("ERROR")
		utils.ErrorToResponse(&rsp.CommonResponse, e)
		return
	}
	filter := bson.D{{req.Key, req.Value}}

	cursor, e := c.Find(context.Background(), filter)
	if e != nil {
		utils.ErrorToResponse(&rsp.CommonResponse, e)
		return
	}

	var operation utils.Operation
	for cursor.Next(context.Background()) {
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
	e = cursor.Close(context.Background())
	if e != nil {
		log.Println("ERROR")
		utils.ErrorToResponse(&rsp.CommonResponse, e)
		return
	}
	rsp.Result = "OK"
}

func FindOneOperation(w http.ResponseWriter, r *http.Request) {
	log.Println("in FindOneOperation")
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

	reqbody, e := utils.AnalyzeRequest(r)
	if e != nil {
		log.Println("ERROR")
		utils.ErrorToResponse(&rsp.CommonResponse, e)
		return
	}

	var req FindOperationRequest
	e = json.Unmarshal(reqbody, &req)
	if e != nil {
		log.Println("ERROR")
		utils.ErrorToResponse(&rsp.CommonResponse, e)
		return
	}
	filter := bson.D{{req.Key, req.Value}}

	var op utils.Operation
	e = c.FindOne(context.Background(), filter).Decode(&op)
	if e != nil {
		log.Println("ERROR")
		utils.ErrorToResponse(&rsp.CommonResponse, e)
		return
	}
	log.Println(op)
	rsp.Operations = append(rsp.Operations, op)
	rsp.Result = "OK"
}

func DeleteOneOperation(w http.ResponseWriter, r *http.Request) {
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

	var req DeleteOperationRequest
	e := utils.StructRequest(r, &req)
	if e != nil {
		log.Println("ERROR")
		utils.ErrorToResponse(&rsp.CommonResponse, e)
		return
	}
	log.Printf("req is %+v", req)

	result, e := c.DeleteOne(
		context.Background(),
		bson.D{
			{req.Key, req.Value}},
	)

	if e != nil {
		log.Println("ERROR")
		utils.ErrorToResponse(&rsp.CommonResponse, e)
		return
	}
	log.Printf("update result is: %+v", result)
	log.Printf("delete count is: %d", result.DeletedCount)
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

	var req DeleteOperationRequest
	e := utils.StructRequest(r, &req)
	if e != nil {
		log.Println("ERROR")
		utils.ErrorToResponse(&rsp.CommonResponse, e)
		return
	}
	log.Printf("req is %+v", req)

	result, e := c.DeleteMany(
		context.Background(),
		bson.D{
			{req.Key, req.Value}},
	)

	if e != nil {
		log.Println("ERROR")
		utils.ErrorToResponse(&rsp.CommonResponse, e)
		return
	}
	log.Printf("update result is: %+v", result)
	log.Printf("delete count is: %d", result.DeletedCount)
	rsp.Result = "OK"
}

func DeleteDatabase(w http.ResponseWriter, r *http.Request) {
	log.Println("in DeleteDatabase")
	var rsp utils.CommonResponse
	defer func() {
		buf, e := json.Marshal(&rsp)
		if e != nil {
			w.WriteHeader(500)
		}
		w.Write([]byte(buf))
	}()

	client := db.GetClient("")
	e := client.Database("test_db").Drop(context.Background())
	if e != nil {
		log.Println("ERROR")
		utils.ErrorToResponse(&rsp, e)
		return
	}
	rsp.Result = "OK"
}

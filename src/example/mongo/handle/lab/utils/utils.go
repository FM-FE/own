package utils

import (
	"encoding/json"
	"example/mongo/handle/op"
	"io/ioutil"
	"log"
	"net/http"
)

func InsertOperation(rsp *op.InsertOperationResponse, i int) error {
	// channel
	channel := make(chan int)
	i = 1

	insertChannel := make(chan op.InsertOperationResponse)

	client := http.Client{}
	request, e := http.NewRequest("GET", "http://localhost:7460/operation/insert", nil)
	if e != nil {
		rsp.Result = "ERROR"
		rsp.Error = e.Error()
		log.Println(e.Error())
		channel <- i
		insertChannel <- *rsp
		return e
	}
	resp, e := client.Do(request)
	if e != nil {
		rsp.Result = "ERROR"
		rsp.Error = e.Error()
		log.Println(e.Error())
		channel <- i
		insertChannel <- *rsp
		return e
	}
	responseJSON, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		rsp.Result = "ERROR"
		rsp.Error = e.Error()
		log.Println(e.Error())
		channel <- i
		insertChannel <- *rsp
		return e
	}
	log.Println(string(responseJSON))

	var response op.InsertOperationResponse
	e = json.Unmarshal(responseJSON, &response)
	if e != nil {
		rsp.Result = "ERROR"
		rsp.Error = e.Error()
		log.Println(e.Error())
		insertChannel <- *rsp
		channel <- i
		return e
	}

	channel <- i
	return nil
}

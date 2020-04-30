package utils

import (
	"bytes"
	"encoding/json"
	"example/mongo/handle/op"
	"go.mongodb.org/mongo-driver/bson"
	"io/ioutil"
	"log"
	"net/http"
)

func JsonToBson(json_in []byte, target interface{}) (bson_out []byte, e error) {
	log.Println(string(json_in))
	e = json.Unmarshal(json_in, target)
	if e != nil {
		return nil, e
	}
	bson_out, e = bson.Marshal(target)
	if e != nil {
		return nil, e
	}
	log.Println(string(bson_out))
	return bson_out, nil
}

func AnalyzeRequest(r *http.Request) (body []byte, e error) {
	if r == nil {
		return
	}
	log.Println("In AnalyzeRequest")
	body, e = ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body)) // make sure request.body always has value
	if e != nil {
		return
	}
	log.Println(string(body))
	return
}

func ErrorToResponse(rsp *op.CommonResponse, e error) {
	log.Println(">>> ERROR OCCUR !!!")
	log.Println(">>>", e.Error())
	rsp.Result = "ERROR"
	rsp.Error = e.Error()
}

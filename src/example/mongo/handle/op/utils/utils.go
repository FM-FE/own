package utils

import (
	"bytes"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"io/ioutil"
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

func JsonToBson(json_in []byte, target interface{}) (bson_out []byte, e error) {
	log.Println(string(json_in))
	e = json.Unmarshal(json_in, target)
	if e != nil {
		return nil, e
	}
	log.Printf("target struct is: %+v", target)
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
	log.Println("reqbody is: ", string(body))
	return
}

func ErrorToResponse(rsp *CommonResponse, e error) {
	log.Println(">")
	log.Println(">>> ERROR OCCUR !!!")
	log.Println(">>>>> ", e.Error())
	log.Println(">>> ERROR OCCUR !!!")
	log.Println(">")
	rsp.Result = "ERROR"
	rsp.Error = e.Error()
}

package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type CommonResponse struct {
	Result string `json:"result"`
	Error  string `json:"error,omitempty"`
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

func StructRequest(r *http.Request, req interface{}) (e error) {
	if r == nil {
		return
	}
	log.Println("In StructRequest")
	reqbody, e := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	r.Body = ioutil.NopCloser(bytes.NewBuffer(reqbody)) // make sure request.body always has value
	if e != nil {
		return
	}
	log.Println("reqbody is: ", string(reqbody))
	e = json.Unmarshal(reqbody, &req)
	if e != nil {
		log.Println("ERROR")
		return e
	}
	log.Printf("req is: %+v", req)
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

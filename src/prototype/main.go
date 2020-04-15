package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"prototype/restful"
)

func main() {
	r := mux.NewRouter()

	// test hello
	r.HandleFunc("/hello", restful.Hello).Methods("GET")

	// forward task
	r.HandleFunc("/task", restful.NewTask).Methods("POST")

	e := http.ListenAndServe(":647", r)
	if e != nil {
		return
	}
}

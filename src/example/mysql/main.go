package main

import (
	"example/mysql/db"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	log.Println("In mysql_basic")
	router := mux.NewRouter()

	// sample
	router.HandleFunc("/first-mysql", db.ListTask).Methods("POST")
	// op & lab

	svr := http.Server{
		Addr: ":7461",
		Handler: handlers.CORS(
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE"}),
			handlers.AllowedOrigins([]string{"*"}))(router),
	}

	e := svr.ListenAndServe()
	if e != nil {
		log.Println(e.Error())
		return
	}
}

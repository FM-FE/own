package main

import (
	"example/mongo/db"
	"example/mongo/handle/lab"
	"example/mongo/handle/op"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	log.Println("In curd")
	router := mux.NewRouter()

	// sample
	router.HandleFunc("/first-mongo", db.CreateAndInsert).Methods("POST")
	// op & lab
	router.HandleFunc("/operation/insert", op.InsertOperation).Methods("GET")
	router.HandleFunc("/operation/insert/{_times}", lab.MultipleInsertOperation).Methods("GET")
	router.HandleFunc("/operation/find", op.FindOperation).Methods("POST")
	router.HandleFunc("/operation/find/one", op.FindOneOperation).Methods("POST")
	router.HandleFunc("/operation/update/one", op.UpdateOneOperation).Methods("POST")
	router.HandleFunc("/operation/update", op.UpdateOperation).Methods("POST")
	router.HandleFunc("/operation/delete/one", op.DeleteOneOperation).Methods("DELETE")
	router.HandleFunc("/operation/delete", op.DeleteOperation).Methods("DELETE")
	router.HandleFunc("/operation/database", op.DeleteDatabase).Methods("DELETE")

	svr := http.Server{
		Addr: ":7460",
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

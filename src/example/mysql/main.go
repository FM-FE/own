package main

import (
	"example/mysql/db"
	board "example/mysql/db/message_board"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	log.Println("In mysql_basic")
	router := mux.NewRouter()

	// task sample
	router.HandleFunc("/first-mysql", db.ListTask).Methods("POST")
	router.HandleFunc("/task/insert", db.InsertTask).Methods("POST")

	// message board
	// login & register
	router.HandleFunc("/login", board.Login).Methods("POST")
	router.HandleFunc("/register", board.Login).Methods("POST")

	// article 
	router.HandleFunc("/article", board.Login).Methods("POST")

	// comment
	router.HandleFunc("/comment", board.Login).Methods("POST")

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

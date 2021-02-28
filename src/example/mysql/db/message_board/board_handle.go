package message_board

import (
	"encoding/json"
	"example/mysql/db/utils"
	"log"
	"net/http"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRequest struct {
	User
}

// login
func Login(w http.ResponseWriter, r *http.Request) {
	log.Println("In Login")
	var rsp utils.CommonResponse
	defer func() {
		buf, e := json.Marshal(&rsp)
		if e != nil {
			w.WriteHeader(500)
		}
		w.Write([]byte(buf))
	}()

	rsp.Result = "OK"
}

func LoginCheck(user User) (e error) {
	log.Println("In LoginCheck")
	
	return
}

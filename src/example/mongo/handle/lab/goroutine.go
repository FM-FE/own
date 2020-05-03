package lab

import (
	"encoding/json"
	"example/mongo/handle/lab/utils"
	"example/mongo/handle/op"
	op_utils "example/mongo/handle/op/utils"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type MultipleInsertResponse struct {
	op_utils.CommonResponse
	InsertSlice []op.InsertOperationResponse `json:"insert_slice"`
}

// TODO var op here as global

func MultipleInsertOperation(w http.ResponseWriter, r *http.Request) {
	log.Println("in MultipleInsertOperation")
	var rsp MultipleInsertResponse
	defer func() {
		buf, e := json.Marshal(&rsp)
		if e != nil {
			w.WriteHeader(500)
		}
		w.Write([]byte(buf))
	}()

	timeStr := mux.Vars(r)["_times"]
	times, e := strconv.Atoi(timeStr)
	if e != nil {
		log.Println(e.Error())
		rsp.Result = "ERROR"
		rsp.Error = e.Error()
		return
	}
	log.Println("times: ", times)

	var insertSlice []op.InsertOperationResponse
	channel := make(chan int)
	insertChannel := make(chan op.InsertOperationResponse)

	for i := 0; i < times; i++ {
		go func() {
			var singleresp op.InsertOperationResponse
			e := utils.InsertOperation(&singleresp, channel, insertChannel)
			if e != nil {
				return
			}
		}()
	}

	for i := 0; i < times; i++ {
		log.Println(<-channel)
		log.Println(<-insertChannel)
	}
	log.Println("=========new add=======")
	log.Println("%+v\n", rsp.InsertSlice)
	log.Println("%+v\n", insertSlice)

	rsp.Result = "OK"
}

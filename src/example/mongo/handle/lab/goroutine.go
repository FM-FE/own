package lab

import (
	"encoding/json"
	"example/mongo/handle/lab/utils"
	"example/mongo/handle/op"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type MultipleInsertResponse struct {
	op.CommonResponse
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
			e := utils.InsertOperation(&singleresp, i)
			if e != nil {
				rsp.InsertSlice = append(rsp.InsertSlice, singleresp)
				insertSlice = append(insertSlice, singleresp)
				return
			}
		}()
	}

	close(channel)
	for i := 0; i < times; i++ {
		if <-channel == 1 {
			log.Println(<-channel)
			log.Println(<-insertChannel)
		}
	}
	log.Println("%+v\n", rsp.InsertSlice)
	log.Println("%+v\n", insertSlice)

	rsp.Result = "OK"
}

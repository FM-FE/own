package db

import (
	"encoding/json"
	"example/mysql/db/mysql_common"
	"example/mysql/db/utils"
	"fmt"
	"log"
	"net/http"
)

type ListTaskResponse struct {
	utils.CommonResponse
	Tasks []mysql_common.Task `json:"tasks"`
}

func ListTask(w http.ResponseWriter, r *http.Request) {
	log.Println("ListTask")
	var rsp ListTaskResponse
	defer func() {
		buf, e := json.Marshal(&rsp)
		if e != nil {
			w.WriteHeader(500)
		}
		w.Write([]byte(buf))
	}()

	db, e := mysql_common.GetMySQLDatabase(mysql_common.DB)
	if e != nil {
		utils.ErrorToResponse(&rsp.CommonResponse, e)
		return
	}
	defer db.Close()

	sqlcmd := fmt.Sprintf("SELECT * FROM %s", mysql_common.Table)
	log.Println("sqlcmd is: " + sqlcmd)
	result, e := db.Query(sqlcmd)
	if e != nil {
		log.Println(e.Error())
		utils.ErrorToResponse(&rsp.CommonResponse, e)
		return
	}
	var task mysql_common.Task
	for result.Next() {
		e := result.Scan(&task.Id, &task.Subject)
		if e != nil {
			utils.ErrorToResponse(&rsp.CommonResponse, e)
			return
		}
		rsp.Tasks = append(rsp.Tasks, task)
	}
	rsp.Result = "OK"
}

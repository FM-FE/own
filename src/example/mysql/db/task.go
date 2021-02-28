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

type InsertRequest struct {
	mysql_common.Task
}

// list
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

	log.Println(">>>>>>>>>>>>>>>>")
	log.Println(r.Host)
	log.Println(r.URL)
	log.Println(r.Method)
	
	log.Println(">>>>>>>>>>>>>>>>")

	e, tasks := ListTaskByMySQL()
	if e != nil {
		utils.ErrorToResponse(&rsp.CommonResponse, e)
		return
	}
	rsp.Tasks = tasks
	rsp.Result = "OK"
}

func ListTaskByMySQL() (e error, tasks []mysql_common.Task) {
	db, e := mysql_common.GetMySQLDatabase(mysql_common.DefaultServer)
	if e != nil {
		return
	}
	defer db.Close()
	sqlcmd := fmt.Sprintf("SELECT * FROM %s", mysql_common.Table)
	log.Println("sqlcmd is: " + sqlcmd)
	result, e := db.Query(sqlcmd)
	if e != nil {
		log.Println(e.Error())
		return
	}
	var task mysql_common.Task
	for result.Next() {
		err := result.Scan(&task.Id, &task.Subject)
		if err != nil {
			return
		}
		tasks = append(tasks, task)
	}
	return
}

// insert
func InsertTask(w http.ResponseWriter, r *http.Request) {
	log.Println("ListTask")
	var rsp ListTaskResponse
	defer func() {
		buf, e := json.Marshal(&rsp)
		if e != nil {
			w.WriteHeader(500)
		}
		w.Write([]byte(buf))
	}()

	var req InsertRequest
	e := utils.StructRequest(r, &req)
	if e != nil {
		utils.ErrorToResponse(&rsp.CommonResponse, e)
		return
	}

	e = InsertTaskByMySQL(req.Task)
	if e != nil {
		utils.ErrorToResponse(&rsp.CommonResponse, e)
		return
	}

	e, tasks := ListTaskByMySQL()
	if e != nil {
		log.Println(e.Error())
		utils.ErrorToResponse(&rsp.CommonResponse, e)
		return
	}
	rsp.Tasks = tasks
	rsp.Result = "OK"
}

func InsertTaskByMySQL(task mysql_common.Task) (e error) {
	db, e := mysql_common.GetMySQLDatabase(mysql_common.DefaultServer)
	if e != nil {
		log.Println(e.Error())
		return
	}
	defer db.Close()

	sqlcmd := fmt.Sprintf("INSERT INTO %s VALUES ( %d , '%s' )", mysql_common.Table, task.Id, task.Subject)
	log.Println("sqlcmd is: " + sqlcmd)
	result, e := db.Query(sqlcmd)
	if e != nil {
		log.Println(e.Error())
		return
	}
	// close result if used transactions
	result.Close()
	return
}

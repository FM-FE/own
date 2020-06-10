package mysql_common

import (
	"database/sql"
	"fmt"
	_ "github.com/mysql"
	"os"
)

var Username = "root"
var Password = os.Getenv("MYSQL_ROOT_PASSWORD")
var IP = "localhost"
var Port = ":3306"

var DB = "TEST"
var Table = "tasks"

type Task struct {
	Id      int    `json:"task_id"`
	Subject string `json:"subject"`
}

func GetMySQLDatabase(dbname string) (db *sql.DB, e error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s%s)/%s", Username, Password, IP, Port, dbname)
	db, e = sql.Open("mysql", dataSourceName)
	return
}

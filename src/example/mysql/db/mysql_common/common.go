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

var DefaultServer = MysqlServer{
	Username: Username,
	Password: Password,
	IP:       IP,
	Port:     Port,
	DB:       DB,
	Table:    Table,
}

type MysqlServer struct {
	Username string `json:"username"`
	Password string `json:"password"`
	IP       string `json:"ip"`
	Port     string `json:"port"`
	DB       string `json:"db"`
	Table    string `json:"table"`
}

type Task struct {
	Id      int    `json:"task_id"`
	Subject string `json:"subject"`
}

func GetMySQLDatabase(s MysqlServer) (db *sql.DB, e error) {
	// format by mysql connection
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s%s)/%s", s.Username, s.Password, s.IP, s.Port, s.DB)
	db, e = sql.Open("mysql", dataSourceName)
	return
}

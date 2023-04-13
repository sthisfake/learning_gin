package database

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func StartDb() {
	var err error
	db , err = sql.Open("mysql" , os.Getenv("DSN"))
	if err != nil {
		panic("failed to connect to db")
	}
}

func CloseDb(){
	 db.Close()
}

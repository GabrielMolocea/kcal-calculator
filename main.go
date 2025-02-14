package main

import (
	"database/sql"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Food struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	Calories int `json:"calories"`
}

func  main() {
	cfg := mysql.Config{
		User: os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net: "tcp",
		Addr: "127.0.0.1:3306",
		DBName: "kcal_calculator",
}

}
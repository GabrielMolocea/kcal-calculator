package database

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func InitDB() {
    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    cfg := mysql.Config{
        User:   os.Getenv("DBUSER"),
        Passwd: os.Getenv("DBPASS"),
        Net:    "tcp",
        Addr:   "127.0.0.1:3306",
        DBName: "kcal_calculator",
    }

    DB, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatalf("Error opening database: %v", err)
    }

    pingErr := DB.Ping()
    if pingErr != nil {
        log.Fatalf("Error pinging database: %v", pingErr)
    }
    log.Println("Connected to database")
}
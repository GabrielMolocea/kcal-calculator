package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)
var db *sql.DB


func main() {
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

	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatalf("Error pinging database: %v", pingErr)
	}
	fmt.Println("Connected to database")

	// Run the SQL script
	// err = runSQLScript(db, "kcal_calculator.sql")
	// if err != nil {
	//     log.Fatalf("Error running SQL script: %v", err)
	// }
	// fmt.Println("SQL script executed successfully")

	router := gin.Default()

	// Enable CORS
	router.Use(cors.Default())

	router.GET("/foods", getFoods)
	router.GET("/targetkcal", getTargetKcal)
	router.POST("/newfood", postFoods)
	router.POST("/targetkcal", postTargetKcal)
	router.POST("/resttargetkcal", restTargetKcal)

	router.Run("localhost:8080") // Change the port number here
}

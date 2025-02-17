package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

type Food struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Calories int64  `json:"calories"`
}

func getFoods(c *gin.Context) {
	var foods []Food
	rows, err := db.Query("SELECT id, name, calories FROM foods")
	if err != nil {
		log.Printf("Error querying database: %v", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var food Food
		if err := rows.Scan(&food.ID, &food.Name, &food.Calories); err != nil {
			log.Printf("Error scanning row: %v", err)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		foods = append(foods, food)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error with rows: %v", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, foods)
}

func postFoods(c *gin.Context) {
	var newFood Food

	if err := c.BindJSON(&newFood); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	result, err := db.Exec("INSERT INTO foods (name, calories) VALUES (?, ?)", newFood.Name, newFood.Calories)
	if err != nil {
		log.Printf("Error inserting into database: %v", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error getting last insert ID: %v", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	newFood.ID = id
	c.IndentedJSON(http.StatusCreated, newFood)
}

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

	router := gin.Default()

	// Enable CORS
	router.Use(cors.Default())

	router.GET("/foods", getFoods)
	router.POST("/newfood", postFoods)

	router.Run("localhost:8080") // Change the port number here
}

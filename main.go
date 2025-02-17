package main

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Food struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	Calories int `json:"calories"`
}


func getFoods(c *gin.Context) {
	rows, err := db.Query("SELECT * FROM foods")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	defer rows.Close()
}

func postFoods(c *gin.Context) {
	var newFood Food

	if err := c.BindJSON(&newFood); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	_, err := db.Exec("INSERT INTO foods (name, calories) VALUES (?, ?)", newFood.Name, newFood.Calories)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

}
func  main() {
	cfg := mysql.Config{
		User: os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net: "tcp",
		Addr: "127.0.0.1:3306",
		DBName: "kcal_calculator",
}

	router := gin.Default()
	router.GET("/foods", getFoods)
	router.POST("/foods", postFoods)
	

}
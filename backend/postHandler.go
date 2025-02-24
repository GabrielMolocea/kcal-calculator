package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Food struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Calories int64  `json:"calories"`
}

type TotalCalories struct {
	Total int64 `json:"total"`
}

type TargetKcal struct {
	ID         int64 `json:"id"`
	TargetKcal int64 `json:"targetKcal"`
}


func postTargetKcal(c *gin.Context) {
	var targetKcal TargetKcal
	if err := c.BindJSON(&targetKcal); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	result, err := db.Exec("INSERT INTO targetKcal (targetKcal) VALLUES (?)", targetKcal.TargetKcal)
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
	targetKcal.ID = id
	c.IndentedJSON(http.StatusCreated, targetKcal)
}

func postFoods(c *gin.Context) {
	var newFood Food
	if err := c.BindJSON(&newFood); err != nil {
		log.Panicf("Error binding JSON: %v", err)
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

	var targetKcal TargetKcal
	row := db.QueryRow("SELECT id, targetKcal FROM targetKcal ORDER BY id DESC LIMIT 1")
	if err := row.Scan(&targetKcal.ID, &targetKcal.TargetKcal); err != nil {
		if err == sql.ErrNoRows {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No target calories found"})
		} else {
			log.Printf("Error scanning row: %v", err)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
		return
	}

	newTargetKcal := targetKcal.TargetKcal - newFood.Calories
	_, err = db.Exec("UPDATE targetKcal SET targetKcal = ? WHERE id = ?", newTargetKcal, targetKcal.ID)
	if err != nil {
		log.Printf("Error updating target calories: %v", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, newFood)
}

func restTargetKcal(c *gin.Context) {
	// Rest to target of 1000
	_, err := db.Exec("UPDATE targetKcal SET targetKcal = 1000")
	if err != nil {
		log.Printf("Error updating target calories: %v", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"target": err.Error()})
		return

	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": 1000})
}
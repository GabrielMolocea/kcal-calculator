package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

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


func getTargetKcal(c *gin.Context) {
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

	c.IndentedJSON(http.StatusOK, targetKcal)
}

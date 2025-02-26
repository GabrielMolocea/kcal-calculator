package postHandlers

import (
	database "kcalcalculator/databse"
	"kcalcalculator/types"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostFoods(c *gin.Context) {
    var newFood types.Food

    if err := c.BindJSON(&newFood); err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
        return
    }

    _, err := database.DB.Exec("INSERT INTO foods (name, calories) VALUES (?, ?)", newFood.Name, newFood.Calories)
    if err != nil {
        log.Printf("Error inserting new food: %v", err)
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
        return
    }

    c.IndentedJSON(http.StatusCreated, newFood)
}

func PostTargetKcal(c *gin.Context) {
    var targetKcal types.TargetKcal

    if err := c.BindJSON(&targetKcal); err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
        return
    }

    _, err := database.DB.Exec("INSERT INTO targetKcal (targetKcal) VALUES (?)", targetKcal.TargetKcal)
    if err != nil {
        log.Printf("Error inserting target calories: %v", err)
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
        return
    }

    c.IndentedJSON(http.StatusCreated, targetKcal)
}

func RestTargetKcal(c *gin.Context) {
    _, err := database.DB.Exec("UPDATE targetKcal SET targetKcal = 0 WHERE id = (SELECT id FROM targetKcal ORDER BY id DESC LIMIT 1)")
    if err != nil {
        log.Printf("Error resetting target calories: %v", err)
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
        return
    }

    c.IndentedJSON(http.StatusOK, gin.H{"message": "Target calories reset"})
}
package getHandlers

import (
	"database/sql"
	database "kcalcalculator/databse"
	"kcalcalculator/types"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetFoods(c *gin.Context) {
    var foods []types.Food
    rows, err := database.DB.Query("SELECT id, name, calories FROM foods")
    if err != nil {
        log.Printf("Error querying database: %v", err)
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
        return
    }
    defer rows.Close()

    for rows.Next() {
        var food types.Food
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

func GetTargetKcal(c *gin.Context) {
    var targetKcal types.TargetKcal

    row := database.DB.QueryRow("SELECT id, targetKcal FROM targetKcal ORDER BY id DESC LIMIT 1")
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
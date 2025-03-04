package main

import (
	database "kcalcalculator/databse"
	"kcalcalculator/getHandlers"
	"kcalcalculator/postHandlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
    // Initialize the database
    database.InitDB()

    router := gin.Default()

    // Enable CORS
    router.Use(cors.Default())

    router.GET("/foods", getHandlers.GetFoods)
    router.GET("/targetkcal", getHandlers.GetTargetKcal)
    router.POST("/newfood", postHandlers.PostFoods)
    router.POST("/targetkcal", postHandlers.PostTargetKcal)
    router.POST("/resttargetkcal", postHandlers.RestTargetKcal)
    router.POST("/deductkcal", postHandlers.DeductKcal)

    router.Run("localhost:8080") // Change the port number here
}
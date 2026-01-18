package main

import (
	"school-manage/config"
	"school-manage/models"
	"school-manage/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Connect to database
	config.ConnectDatabase()

	// Auto Migrate
	config.DB.AutoMigrate(&models.User{}, &models.Classroom{}, &models.Schedule{}, &models.Booking{}, &models.StudentWallet{})

	// Setup Routes
	routes.SetupRoutes(r)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(":8080")
}

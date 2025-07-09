package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"stringy-api/config"
	"stringy-api/routes"
)

func main() {

	if err := godotenv.Load("cmd/server/.env"); err != nil {
		log.Println("No .env file found")
	}

	config.Connect()

	router := gin.Default()
	routes.RegisterRoutes(router)
	routes.RoomRoutes(router)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "healthy",
			"message": "Task API is running",
			"version": "1.0.0",
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server starting on port %s\n", port)
	log.Fatal(router.Run(":" + port))

}

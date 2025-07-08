package main

import (
	"stringy-api/config"
	"stringy-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Connect()

	r := gin.Default()
	routes.RegisterRoutes(r)
	routes.RoomRoutes(r)

	r.Run(":8080") // Listen on port 8080
}

package routes

import (
	"stringy-api/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/users", handlers.GetUsers)
	}
}

func RoomRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/rooms", handlers.GetRooms)
	}
}

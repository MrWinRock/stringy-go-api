package routes

import (
	"stringy-api/handlers"
	"stringy-api/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	userApi := r.Group("/api/users")
	{
		userApi.GET("/", handlers.GetUsers)
		userApi.POST("/register", handlers.CreateUser)
		userApi.POST("/login", handlers.LoginUser)
	}

	protectedUserApi := r.Group("/api/users")
	protectedUserApi.Use(middleware.AuthMiddleware())
	{
		protectedUserApi.GET("/profile", handlers.GetMyProfile)
	}
}

func RoomRoutes(r *gin.Engine) {
	roomApi := r.Group("/api/rooms")
	{
		roomApi.GET("/", handlers.GetRooms)
	}
}

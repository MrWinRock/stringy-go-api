package handlers

import (
	"net/http"

	"stringy-api/config"
	"stringy-api/models"

	"github.com/gin-gonic/gin"
)

func GetRooms(c *gin.Context) {
	var rooms []models.Room
	query := `
	SELECT room_id, title, description FROM rooms
	`
	err := config.DB.Select(&rooms, query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, rooms)
}

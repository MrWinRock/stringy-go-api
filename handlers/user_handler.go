package handlers

import (
	"net/http"

	"stringy-api/config"
	"stringy-api/models"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	query := `
	SELECT user_id, username, email, role FROM users
	`
	err := config.DB.Select(&users, query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

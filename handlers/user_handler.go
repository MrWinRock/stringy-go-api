package handlers

import (
	"net/http"
	"time"

	"stringy-api/config"
	"stringy-api/models"
	"stringy-api/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingCount int
	checkQuery := `SELECT COUNT(*) FROM users WHERE email = $1`
	err := config.DB.Get(&existingCount, checkQuery, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check existing user"})
		return
	}

	if existingCount > 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "A user with this email already exists"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = string(hashedPassword)

	if user.Role == "" {
		user.Role = "user"
	}
	if user.CreatedAt == "" {
		user.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	}

	query := `
		INSERT INTO users (username, email, password, role, created_at, profile_picture_url)
		VALUES (:username, :email, :password, :role, :created_at, :profile_picture_url)
		RETURNING user_id
	`

	stmt, err := config.DB.PrepareNamed(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var insertedID int
	if err := stmt.Get(&insertedID, user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user.UserID = insertedID
	user.Password = ""

	c.JSON(http.StatusCreated, user)
}

func GetMyProfile(c *gin.Context) {
	userID := c.GetUint("user_id")

	var user models.User
	query := `SELECT user_id, username, email, role, created_at, profile_picture_url FROM users WHERE user_id = $1`

	err := config.DB.Get(&user, query, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func LoginUser(c *gin.Context) {
	var loginData struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
		return
	}

	var user models.User
	query := `SELECT user_id, username, email, password, role, created_at, profile_picture_url FROM users WHERE email = $1`

	err := config.DB.Get(&user, query, loginData.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	token, err := utils.GenerateJWT(uint(user.UserID), user.Username, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	user.Password = ""

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user":    user,
		"token":   token,
	})
}

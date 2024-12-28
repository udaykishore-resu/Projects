package handlers

import (
	"net/http"
	"paybackapp/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *gin.Context) {
	var req models.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user := models.User{
		ID:       uuid.New().String(),
		Email:    req.Email,
		Name:     req.Name,
		Password: string(hashedPassword),
		Points:   0,
	}

	// In a real application, save user to database here

	token := generateToken(user.ID) // Implement JWT token generation

	c.JSON(http.StatusOK, models.AuthResponse{
		Token: token,
		User:  user,
	})
}

func LoginUser(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// In a real application, fetch user from database and verify password
	// This is a mock implementation
	c.JSON(http.StatusOK, models.AuthResponse{
		Token: "mock-token",
		User: models.User{
			ID:     "1",
			Email:  req.Email,
			Name:   "Test User",
			Points: 100,
		},
	})
}

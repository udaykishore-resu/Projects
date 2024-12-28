package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Mock data - in a real application, this would be in a database
var restaurants = []models.Restaurant{
	{
		ID:          "1",
		Name:        "The Fancy Place",
		Description: "Fine dining at its best",
		Address:     "123 Main St",
		Rating:      4.5,
	},
}

func GetRestaurants(c *gin.Context) {
	c.JSON(http.StatusOK, restaurants)
}

func GetRestaurant(c *gin.Context) {
	id := c.Param("id")
	for _, r := range restaurants {
		if r.ID == id {
			c.JSON(http.StatusOK, r)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Restaurant not found"})
}

func CreateRestaurant(c *gin.Context) {
	var req models.CreateRestaurantRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	restaurant := models.Restaurant{
		ID:          uuid.New().String(),
		Name:        req.Name,
		Description: req.Description,
		Address:     req.Address,
		Rating:      req.Rating,
	}

	restaurants = append(restaurants, restaurant)
	c.JSON(http.StatusCreated, restaurant)
}

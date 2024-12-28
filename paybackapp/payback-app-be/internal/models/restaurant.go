package models

type Restaurant struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Address     string  `json:"address"`
	Rating      float64 `json:"rating"`
}

type CreateRestaurantRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Address     string  `json:"address" binding:"required"`
	Rating      float64 `json:"rating" binding:"required"`
}

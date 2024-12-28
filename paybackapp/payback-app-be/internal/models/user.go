package models

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"-"` // "-" means this field won't be included in JSON
	Points   int    `json:"points"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
}

type AuthResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}
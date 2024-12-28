package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {
	r := gin.Default()

	// Configure CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"} // Vite's default port
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	r.Use(cors.New(config))

	// Routes
	r.POST("/api/register", handlers.RegisterUser)
	r.POST("/api/login", handlers.LoginUser)
	r.GET("/api/restaurants", handlers.GetRestaurants)
	r.POST("/api/restaurants", handlers.CreateRestaurant)
	r.GET("/api/restaurants/:id", handlers.GetRestaurant)
	r.GET("/api/points/:userId", handlers.GetUserPoints)
	r.POST("/api/points", handlers.AddPoints)

	log.Fatal(r.Run(":8080"))
}
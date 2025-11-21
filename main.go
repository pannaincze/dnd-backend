package main

import (
	"dnd-backend/config"
	"dnd-backend/db"
	"dnd-backend/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	// Connect to PostgreSQL
	db.Connect(cfg.DatabaseURL)

	r := gin.Default()

	// Routes
	r.GET("/characters", handlers.GetCharacters)

	// Start server
	r.Run(":" + cfg.Port)
}

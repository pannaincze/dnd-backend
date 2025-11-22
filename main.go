package main

import (
	"dnd-backend/handlers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/pannaincze/dnd-backend/db"
	"github.com/pannaincze/dnd-backend/handlers"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found, relying on environment variables.")
    }

    databaseURL := os.Getenv("DATABASE_URL")
    if databaseURL == "" {
        log.Fatal("DATABASE_URL not set in environment or .env file.")
    }

    db.Connect(databaseURL)
    defer db.DB.Close() 

    router := gin.Default()

    router.POST("/characters", handlers.CreateCharacter) 
    router.GET("/characters", handlers.ListCharacters) 
    router.GET("/characters/:id", handlers.GetCharacter) 
    router.PUT("/characters/:id", handlers.UpdateCharacter)
		router.DELETE("/characters/:id", handlers.DeleteCharacter)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    log.Printf("Server listening on port %s...", port)
    router.Run(":" + port)
}
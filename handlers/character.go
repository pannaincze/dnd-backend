package handlers

import (
	"context"
	"net/http"

	"dnd-backend/db"
	"dnd-backend/models"
	"github.com/gin-gonic/gin"
)

func GetCharacters(c *gin.Context) {
	rows, err := db.DB.Query(context.Background(),
		"SELECT id, name, class, level, owner_id FROM characters")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var chars []models.Character

	for rows.Next() {
		var ch models.Character
		err := rows.Scan(&ch.ID, &ch.Name, &ch.Class, &ch.Level, &ch.OwnerID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		chars = append(chars, ch)
	}

	c.JSON(http.StatusOK, chars)
}

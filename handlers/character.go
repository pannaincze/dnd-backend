package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/pannaincze/dnd-backend/db"
    "github.com/pannaincze/dnd-backend/models"
)

func ListCharacters(c *gin.Context) {
	var characters []models.Character 

	err := db.DB.Select(&characters, "SELECT * FROM characters ORDER BY id")
	
	if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve characters: " + err.Error()})
			return
	}
	
	c.JSON(http.StatusOK, characters)
}

func GetCharacter(c *gin.Context) {
	id := c.Param("id") 
	
	var character models.Character
	err := db.DB.Get(&character, "SELECT * FROM characters WHERE id=$1", id) 
	
	if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Character not found"})
			return
	}

	c.JSON(http.StatusOK, character)
}

func UpdateCharacter(c *gin.Context) {
	id := c.Param("id")

	var char models.Character
	if err := c.ShouldBindJSON(&char); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}

	result, err := db.DB.Exec(
			`UPDATE characters 
			 SET name=$1,race=$2,class=$3,level=$4,str=$5,dex=$6,con=$7,int=$8,wis=$9,cha=$10,hp=$11,ac=$12,notes=$13,updated_at=NOW() 
			 WHERE id=$14`,
			char.Name, char.Race, char.Class, char.Level,
			char.Str, char.Dex, char.Con, char.Int, char.Wis, char.Cha,
			char.HP, char.AC, char.Notes,
			id,
	)
	if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update character: " + err.Error()})
			return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check update status."})
			return
	}

	if rowsAffected == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Character not found"})
			return
	}

	c.JSON(http.StatusOK, char)
}

func CreateCharacter(c *gin.Context) {
	var char models.Character
	if err := c.ShouldBindJSON(&char); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}

	err := db.DB.QueryRowx(
			`INSERT INTO characters 
			 (name, race, class, level, str, dex, con, int, wis, cha, hp, ac, notes) 
			 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) 
			 RETURNING id, created_at, updated_at`,
			char.Name, char.Race, char.Class, char.Level,
			char.Str, char.Dex, char.Con, char.Int, char.Wis, char.Cha,
			char.HP, char.AC, char.Notes,
	).Scan(
			&char.ID, 
			&char.CreatedAt, 
			&char.UpdatedAt,
	) 

	if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create character: " + err.Error()})
			return
	}

	c.JSON(http.StatusCreated, char)
}

func DeleteCharacter(c *gin.Context) {
	id := c.Param("id") 

	result, err := db.DB.Exec("DELETE FROM characters WHERE id=$1", id)
	
	if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete character: " + err.Error()})
			return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check deletion status."})
			return
	}

	if rowsAffected == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Character not found"})
			return
	}

	c.JSON(http.StatusNoContent, nil) 
}
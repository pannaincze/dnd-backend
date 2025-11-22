package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type CharacterNotes map[string]interface{}

func (n CharacterNotes) Value() (driver.Value, error) {
    if n == nil {
        return nil, nil
    }
	return json.Marshal(n)
}

func (n *CharacterNotes) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

    if b == nil {
        *n = make(CharacterNotes)
        return nil
    }
    
	return json.Unmarshal(b, n)
}


// --- The Main Character Struct ---

type Character struct {
    // Basic Attributes
    ID    int    `db:"id" json:"id"`
    Name  string `db:"name" json:"name"`
    Race  string `db:"race" json:"race"`
    Class string `db:"class" json:"class"`

    // Stats and Combat
    Level int `db:"level" json:"level"`
    HP    int `db:"hp" json:"hp"`
    AC    int `db:"ac" json:"ac"`

    Str int `db:"str" json:"str"`
    Dex int `db:"dex" json:"dex"`
    Con int `db:"con" json:"con"`
    Int int `db:"int" json:"int"`
    Wis int `db:"wis" json:"wis"`
    Cha int `db:"cha" json:"cha"`

    Notes CharacterNotes `db:"notes" json:"notes"` 

    CreatedAt time.Time `db:"created_at" json:"created_at"`
    UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
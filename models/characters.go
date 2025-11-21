package models

type Character struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Class    string `json:"class"`
	Level    int    `json:"level"`
	OwnerID  int    `json:"owner_id"`
}

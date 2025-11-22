package models

type Character struct {
    ID    int                    `db:"id" json:"id"`
    Name  string                 `db:"name" json:"name"`
    Race  string                 `db:"race" json:"race"`
    Class string                 `db:"class" json:"class"`
    Level int                    `db:"level" json:"level"`
    Str   int                    `db:"str" json:"str"`
    Dex   int                    `db:"dex" json:"dex"`
    Con   int                    `db:"con" json:"con"`
    Int   int                    `db:"int" json:"int"`
    Wis   int                    `db:"wis" json:"wis"`
    Cha   int                    `db:"cha" json:"cha"`
    HP    int                    `db:"hp" json:"hp"`
    AC    int                    `db:"ac" json:"ac"`
    Notes map[string]interface{} `db:"notes" json:"notes"` // JSONB
}

package db

import (
    "log"

    "github.com/jmoiron/sqlx"
    _ "github.com/lib/pq"
)

var DB *sqlx.DB

func Connect(databaseURL string) {
    var err error
    DB, err = sqlx.Connect("postgres", databaseURL)
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    log.Println("Connected to PostgreSQL!")
}

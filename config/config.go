package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	Port string
}

func LoadConfig() Config {
	_ = godotenv.Load()

	cfg := Config {
		DatabaseURL: os.Getenv("DATABASE_URL"),
		Port: os.Getenv("PORT"),
	}

	if cfg.DatabaseURL == "" {
		log.Fatal("DATABASE_URL not set")
 	}

	if cfg.Port == "" {
		cfg.Port = "8080"
	}

	return cfg
}
package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Port string
}

func Load() *AppConfig {
	err := godotenv.Load()

	if err != nil {
		log.Println(".env file doenst exist")
	}

	return &AppConfig{
		Port: os.Getenv("PORT"),
	}
}

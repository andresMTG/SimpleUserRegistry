package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Addrs string
}

func Load() *AppConfig {
	err := godotenv.Load()

	if err != nil {
		log.Println(".env file doenst exist")
	}

	return &AppConfig{
		Addrs: os.Getenv("ADDRESS"),
	}
}

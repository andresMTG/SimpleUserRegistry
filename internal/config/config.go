package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"


    "github.com/go-sql-driver/mysql"
)

type AppConfig struct {
	Port string 
	DBConfig *mysql.Config
}

func Load() *AppConfig {
	err := godotenv.Load()

	if err != nil {
		log.Println(".env file doenst exist")
	}

	return &AppConfig{
		Port: os.Getenv("PORT"),
		DBConfig: loadBD(),
	}
}

func loadBD() *mysql.Config {
	err := godotenv.Load()

	if err != nil {
		log.Println(".env file doenst exist")
	}

	return &mysql.Config{
		User: os.Getenv("DB_USER"),
		Passwd: os.Getenv("DB_PASSWORD"),
		Net: os.Getenv("DB_NET"),
		Addr: os.Getenv("DB_HOST"),
		DBName: os.Getenv("DB_NAME"),	
	}
}

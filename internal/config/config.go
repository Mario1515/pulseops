package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	AppPort    string
}

func Load() *Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("no .env file found")
	}
	log.Println("DB_NAME:", os.Getenv("DB_NAME"))

	return &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBPort:     os.Getenv("DB_PORT"),
		AppPort:    os.Getenv("APP_PORT"),
	}
}

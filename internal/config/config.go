package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     string
}

type AppConfig struct {
	Port   string
	APIKey string
}

type Config struct {
	DB  DBConfig
	App AppConfig
}

func Load() *Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("no .env file found")
	}
	return &Config{
		DB: DBConfig{
			Host:     os.Getenv("DB_HOST"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
			Port:     os.Getenv("DB_PORT"),
		},
		App: AppConfig{
			Port:   os.Getenv("APP_PORT"),
			APIKey: os.Getenv("API_KEY"),
		},
	}
}

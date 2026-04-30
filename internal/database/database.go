package database

import (
	"fmt"
	"log"
	"pulseops/internal/config"
	"pulseops/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(cfg *config.DBConfig) *gorm.DB {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("db connect failed:", err)
	}

	db.AutoMigrate(
		&model.User{},
		&model.Incident{},
	)

	log.Println("db connected")
	return db
}

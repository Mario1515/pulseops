package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id"         gorm:"primaryKey"`
	Name      string         `json:"name"       gorm:"not null"                validate:"required,min=2,max=100"`
	Email     string         `json:"email"      gorm:"uniqueIndex;not null"    validate:"required,email"`
	Password  string         `json:"-"          gorm:"not null"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `json:"-"          gorm:"index"`
}

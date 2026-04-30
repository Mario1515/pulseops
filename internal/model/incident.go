package model

import (
	"time"

	"gorm.io/gorm"
)

type Incident struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Title       string         `json:"title" gorm:"not null" validate:"required,min=3,max=150"`
	Description string         `json:"description" gorm:"not null" validate:"required,min=10,max=1000"`
	StartedAt   time.Time      `json:"started_at" gorm:"not null" validate:"required"`
	EndedAt     *time.Time     `json:"ended_at" gorm:"default:null" validate:"omitempty,gtfield=StartedAt"`
	CreatedAt   time.Time      `json:"created_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

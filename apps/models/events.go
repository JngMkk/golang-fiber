package models

import (
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model

	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"size:30;not null"`
	Description string `gorm:"size:100"`
	Location    string `gorm:"size:50"`
	Image       string `gorm:"size:255"`
	UserID      uint   `gorm:"not null"`
}

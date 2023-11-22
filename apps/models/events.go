package models

import (
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model

	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"size:30;not null"`
	Description string `gorm:"size:100;not null"`
	Location    string `gorm:"size:50;not null"`
	Image       string `gorm:"size:255;not null"`
	UserID      uint   `gorm:"not null"`
}

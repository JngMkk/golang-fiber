package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	ID       uint    `gorm:"primaryKey"`
	Email    string  `gorm:"uniqueIndex;size:255;not null"`
	Password string  `gorm:"not null"`
	IsActive bool    `gorm:"default:true;not null"`
	Events   []Event `gorm:"foreignKey:UserID"`
}

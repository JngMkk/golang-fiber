package database

import (
	"fmt"
	"os"
	"time"

	"github.com/JngMkk/golang-fiber/apps/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := os.Getenv("MYSQL_DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("DB connection Error: ", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("DB connection Error: ", err)
	}

	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetConnMaxLifetime(10 * time.Minute)

	return db
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{}, &models.Event{})
}

package queries

import (
	"github.com/JngMkk/golang-fiber/apps/models"
	"github.com/JngMkk/golang-fiber/apps/schemas/requests"
	"gorm.io/gorm"
)

func CreateUserQuery(db *gorm.DB, body *requests.CreateUserBody) (*models.User, error) {
	user := models.User{Email: body.Email, Password: body.Password}
	if err := db.Table("users").Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

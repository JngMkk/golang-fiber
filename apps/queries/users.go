package queries

import (
	"github.com/JngMkk/golang-fiber/apps/models"
	"github.com/JngMkk/golang-fiber/apps/schemas/requests"
	"gorm.io/gorm"
)

func SignUpQuery(db *gorm.DB, body *requests.SignUpBody) (*models.User, error) {
	user := models.User{Email: body.Email, Password: body.Password}
	if err := db.Table("users").Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func SignInQuery(db *gorm.DB, email string) (*models.User, error) {
	user := new(models.User)
	if err := db.Table("users").Where(&models.User{Email: email}).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func DetailUserQuery(db *gorm.DB, id uint) (*models.User, error) {
	user := new(models.User)

	// Unscoped: delete_at where 절에서 제거
	if err := db.Table("users").Unscoped().First(&user, id).Error; err != nil {
		return nil, err
	}

	return user, nil
}

package queries

import (
	"github.com/JngMkk/golang-fiber/apps/models"
	"github.com/JngMkk/golang-fiber/apps/schemas/requests"
	"gorm.io/gorm"
)

func CreateEventQuery(db *gorm.DB, body *requests.CreateEventBody, userID uint) (*models.Event, error) {
	event := models.Event{
		Title:       body.Title,
		Description: body.Description,
		Location:    body.Location,
		Image:       body.Image,
		UserID:      userID,
	}
	if err := db.Table("events").Create(&event).Error; err != nil {
		return nil, err
	}

	return &event, nil
}

func ListEventsQuery(db *gorm.DB, userID uint) (*[]models.Event, error) {
	events := new([]models.Event)
	if err := db.Table("events").Where(&models.Event{UserID: userID}).Order("ID desc").Find(&events).Error; err != nil {
		return nil, err
	}

	return events, nil
}

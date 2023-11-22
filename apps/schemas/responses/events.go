package responses

import "github.com/JngMkk/golang-fiber/apps/models"

type EventResp struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Location    string `json:"location"`
	Image       string `json:"image"`
	UserID      uint   `json:"userID"`
}

func NewEventResp(event *models.Event) *EventResp {
	res := new(EventResp)

	res.ID = event.ID
	res.Title = event.Title
	res.Description = event.Description
	res.Location = event.Location
	res.Image = event.Image
	res.UserID = event.UserID
	return res
}

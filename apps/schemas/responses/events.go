package responses

import (
	"time"

	"github.com/JngMkk/golang-fiber/apps/models"
)

type EventResp struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	Image       string    `json:"image"`
	CreatedAt   time.Time `json:"createdAt"`
	UserID      uint      `json:"userID"`
}

func NewEventResp(event *models.Event) *EventResp {
	res := new(EventResp)

	res.ID = event.ID
	res.Title = event.Title
	res.Description = event.Description
	res.Location = event.Location
	res.Image = event.Image
	res.CreatedAt = event.CreatedAt
	res.UserID = event.UserID
	return res
}

type EventsResp struct {
	Events []*EventResp `json:"events"`
	Counts int          `json:"counts"`
}

func NewEventsResp(events *[]models.Event) *EventsResp {
	res := new(EventsResp)
	res.Events = make([]*EventResp, 0)
	for _, event := range *events {
		eventRes := NewEventResp(&event)
		res.Events = append(res.Events, eventRes)
	}

	res.Counts = len(res.Events)
	return res
}

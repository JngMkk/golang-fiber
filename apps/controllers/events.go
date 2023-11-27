package controllers

import (
	"net/http"

	"github.com/JngMkk/golang-fiber/apps/queries"
	"github.com/JngMkk/golang-fiber/apps/schemas/requests"
	"github.com/JngMkk/golang-fiber/apps/schemas/responses"
	"github.com/JngMkk/golang-fiber/core/auth"
	"github.com/JngMkk/golang-fiber/core/database"
	"github.com/JngMkk/golang-fiber/core/handlers"
	"github.com/gofiber/fiber/v2"
)

// @Summary Create Event
// @Tags events
// @Accept json
// @Produce json
// @Param user body requests.CreateEventBody true "event info for creation"
// @Success 201 {object} responses.EventResp
// @Failure 400 {object} handlers.HTTPError
// @Failure 401 {object} handlers.HTTPError
// @Failure 409 {object} handlers.HTTPError
// @Failure 422 {object} handlers.HTTPError
// @Failure 503 {object} handlers.HTTPError
// @Security ApiKeyAuth
// @Router /events [post]
func CreateEvent(c *fiber.Ctx) error {
	body := &requests.CreateEventBody{}
	if err := c.BodyParser(body); err != nil {
		return handlers.NewHTTPResp(c, http.StatusBadRequest, err)
	}

	v := handlers.NewValidator()
	validBody, err := body.Validate(c, v)
	if err != nil {
		return handlers.NewHTTPResp(c, http.StatusUnprocessableEntity, err)
	}

	tokenData, err := auth.GetAccessTokenData(c)
	if err != nil {
		return handlers.NewHTTPResp(c, http.StatusServiceUnavailable, err)
	}

	db := database.Connect()
	event, err := queries.CreateEventQuery(db, validBody, tokenData.UserID)
	if err != nil {
		return handlers.NewHTTPResp(c, http.StatusConflict, err)
	}

	return handlers.NewHTTPResp(c, http.StatusCreated, responses.NewEventResp(event))
}

// @Summary List Events
// @Tags events
// @Accept json
// @Produce json
// @Success 200 {object} responses.EventsResp
// @Failure 401 {object} handlers.HTTPError
// @Failure 503 {object} handlers.HTTPError
// @Security ApiKeyAuth
// @Router /events [get]
func ListEvents(c *fiber.Ctx) error {
	tokenData, err := auth.GetAccessTokenData(c)
	if err != nil {
		return handlers.NewHTTPResp(c, http.StatusServiceUnavailable, err)
	}

	db := database.Connect()
	events, err := queries.ListEventsQuery(db, tokenData.UserID)
	if err != nil {
		return handlers.NewHTTPResp(c, http.StatusServiceUnavailable, err)
	}

	return handlers.NewHTTPResp(c, http.StatusOK, responses.NewEventsResp(events))
}

package controllers

import (
	"net/http"
	"strconv"

	"github.com/JngMkk/golang-fiber/apps/queries"
	"github.com/JngMkk/golang-fiber/apps/schemas/requests"
	"github.com/JngMkk/golang-fiber/apps/schemas/responses"
	"github.com/JngMkk/golang-fiber/core/database"
	"github.com/JngMkk/golang-fiber/core/handlers"
	"github.com/gofiber/fiber/v2"
)

// @Summary Sign Up user
// @Tags users
// @Accept json
// @Produce json
// @Param user body requests.SignUpBody true "User info for registration"
// @Success 201 {object} responses.UserResp
// @Failure 409 {object} handlers.HTTPError
// @Failure 422 {object} handlers.HTTPError
// @Failure 503 {object} handlers.HTTPError
// @Router /users/signup [post]
func SignUpUser(c *fiber.Ctx) error {
	body := &requests.SignUpBody{}
	if err := c.BodyParser(body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(handlers.NewError(err))
	}

	v := handlers.NewValidator()
	validBody, err := body.Validate(c, v)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(handlers.NewError(err))
	}

	db := database.Connect()
	user, err := queries.SignUpQuery(db, validBody)
	if err != nil {
		return c.Status(http.StatusConflict).JSON(handlers.NewError(err))
	}

	return c.Status(http.StatusCreated).JSON(responses.NewUserResponse(user))
}

// @Summary Get user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} responses.UserResp
// @Failure 404 {object} handlers.HTTPError
// @Failure 503 {object} handlers.HTTPError
// @Router /users/{id} [get]
func DetailUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(handlers.NewError(err))
	}
	db := database.Connect()
	user, err := queries.DetailUserQuery(db, id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(handlers.NewNotFound())
	}

	return c.Status(http.StatusOK).JSON(responses.NewUserResponse(user))
}

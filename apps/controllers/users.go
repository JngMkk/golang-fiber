package controllers

import (
	"net/http"

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
// @Param user body requests.CreateUserBody true "User info for registration"
// @Success 201 {object} responses.UserResp
// @Failure 409 {object} handlers.HTTPError
// @Failure 422 {object} handlers.HTTPError
// @Failure 503 {object} handlers.HTTPError
// @Router /users/signup [post]
func CreateUser(c *fiber.Ctx) error {
	v := handlers.NewValidator()
	body := new(requests.CreateUserBody)
	validBody, err := body.Validate(c, v)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(handlers.NewError(err))
	}

	db := database.Connect()
	user, err := queries.CreateUserQuery(db, validBody)
	if err != nil {
		return c.Status(http.StatusConflict).JSON(handlers.NewError(err))
	}

	return c.Status(http.StatusCreated).JSON(responses.NewUserResponse(user))
}

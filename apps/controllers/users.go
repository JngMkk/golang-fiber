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
	validBody, err := body.ValidateSignUp(c, v)
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

// @Summary Sign In user
// @Tags users
// @Accept json
// @Produce json
// @Param user body requests.SignUpBody true "User info for login"
// @Success 200 {object} auth.Tokens
// @Failure 400 {object} handlers.HTTPError
// @Failure 403 {object} handlers.HTTPError
// @Failure 404 {object} handlers.HTTPError
// @Failure 422 {object} handlers.HTTPError
// @Failure 503 {object} handlers.HTTPError
// @Router /users/signin [post]
func SignInUser(c *fiber.Ctx) error {
	body := &requests.SignUpBody{}
	if err := c.BodyParser(body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(handlers.NewError(err))
	}

	v := handlers.NewValidator()
	validBody, err := body.ValidateSignIn(c, v)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(handlers.NewError(err))
	}

	db := database.Connect()
	user, err := queries.SignInQuery(db, validBody.Email)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(handlers.NewNotFound())
	}

	if !auth.CheckPassword(validBody.Password, user.Password) {
		return c.Status(http.StatusForbidden).JSON(handlers.NewAccessForbidden())
	}

	tokens, err := auth.GenereateTokens(user.ID)
	if err != nil {
		return c.Status(http.StatusServiceUnavailable).JSON(handlers.NewError(err))
	}
	return c.Status(http.StatusOK).JSON(tokens)
}

// @Summary Get user info by token
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} responses.UserResp
// @Failure 404 {object} handlers.HTTPError
// @Failure 503 {object} handlers.HTTPError
// @Router /users/my [get]
func DetailUser(c *fiber.Ctx) error {
	data, err := auth.GetTokenData(c)
	if err != nil {
		return c.Status(http.StatusServiceUnavailable).JSON(handlers.NewError(err))
	}
	db := database.Connect()
	user, err := queries.DetailUserQuery(db, data.UserID)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(handlers.NewNotFound())
	}

	return c.Status(http.StatusOK).JSON(responses.NewUserResponse(user))
}

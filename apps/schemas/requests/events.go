package requests

import (
	"github.com/JngMkk/golang-fiber/core/handlers"
	"github.com/gofiber/fiber/v2"
)

type CreateEventBody struct {
	Title       string `json:"title" validate:"required,max=30"`
	Description string `json:"description" validate:"max=100"`
	Location    string `json:"location" validate:"max=50"`
	Image       string `json:"image" validate:"max=255"`
}

func (body *CreateEventBody) Validate(c *fiber.Ctx, v *handlers.Validator) (*CreateEventBody, error) {
	if err := v.Validate(body); err != nil {
		return nil, err
	}

	return body, nil
}

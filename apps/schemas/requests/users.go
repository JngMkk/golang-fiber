package requests

import (
	"github.com/JngMkk/golang-fiber/core/auth"
	"github.com/JngMkk/golang-fiber/core/handlers"
	"github.com/gofiber/fiber/v2"
)

type SignUpBody struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password"`
}

func (body *SignUpBody) Validate(c *fiber.Ctx, v *handlers.Validator) (*SignUpBody, error) {
	var err error
	if err = v.Validate(body); err != nil {
		return nil, err
	}
	if err = auth.ValidatePassword(body.Password); err != nil {
		return nil, err
	}

	body.Password, err = auth.HashPassword(body.Password)
	if err != nil {
		return nil, err
	}
	return body, nil
}

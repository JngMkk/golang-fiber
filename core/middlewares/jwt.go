package middlewares

import (
	"net/http"

	"github.com/JngMkk/golang-fiber/core/config"
	"github.com/JngMkk/golang-fiber/core/handlers"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func JWTProtected() func(*fiber.Ctx) error {
	config := jwtware.Config{
		SigningKey:   jwtware.SigningKey{JWTAlg: "HS256", Key: []byte(config.JWTSecret)},
		ErrorHandler: jwtError,
	}

	return jwtware.New(config)
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(http.StatusBadRequest).JSON(handlers.NewError(err))
	}

	return c.Status(http.StatusUnauthorized).JSON(handlers.NewError(err))
}

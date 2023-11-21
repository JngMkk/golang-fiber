package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func GetSwaggerRoutes(a *fiber.App) {
	route := a.Group("/swagger")

	route.Get("*", swagger.HandlerDefault)
}

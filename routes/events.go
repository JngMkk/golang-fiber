package routes

import (
	"github.com/JngMkk/golang-fiber/apps/controllers"
	"github.com/JngMkk/golang-fiber/core/middlewares"
	"github.com/gofiber/fiber/v2"
)

func GetEventRoutes(a *fiber.App) {
	route := a.Group("/api")

	route.Post("/events", middlewares.JWTProtected(), controllers.CreateEvent)
}

package routes

import (
	"github.com/JngMkk/golang-fiber/apps/controllers"
	"github.com/gofiber/fiber/v2"
)

func GetUserRoutes(a *fiber.App) {
	route := a.Group("/api")

	route.Post("/users/signup", controllers.CreateUser)
	route.Get("/users/:id", controllers.DetailUser)
}

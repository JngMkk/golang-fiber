package routes

import (
	"github.com/JngMkk/golang-fiber/apps/controllers"
	"github.com/JngMkk/golang-fiber/core/middlewares"
	"github.com/gofiber/fiber/v2"
)

func GetUserRoutes(a *fiber.App) {
	route := a.Group("/api")

	route.Post("/users/signup", controllers.SignUpUser)
	route.Post("/users/signin", controllers.SignInUser)
	route.Post("/users/reissue", controllers.ReIssueUserToken)
	route.Get("/users/:id", middlewares.JWTProtected(), controllers.DetailUser)
}

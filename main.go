package main

import (
	"github.com/JngMkk/golang-fiber/core/database"
	"github.com/JngMkk/golang-fiber/core/handlers"
	"github.com/JngMkk/golang-fiber/core/middlewares"
	_ "github.com/JngMkk/golang-fiber/docs"
	"github.com/JngMkk/golang-fiber/routes"
	"github.com/gofiber/fiber/v2"
)

// @title Fiber Example API
// @version 1.0
// @BasePath /api
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	db := database.Connect()
	database.AutoMigrate(db)

	app := fiber.New()

	routes.GetUserRoutes(app)
	routes.GetEventRoutes(app)
	routes.GetSwaggerRoutes(app)

	middlewares.FiberMiddleWare(app)
	handlers.StartServer(app)
}

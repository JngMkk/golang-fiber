package main

import (
	"github.com/JngMkk/golang-fiber/core/database"
	"github.com/JngMkk/golang-fiber/core/handlers"
	"github.com/JngMkk/golang-fiber/core/middlewares"
	_ "github.com/JngMkk/golang-fiber/docs"
	"github.com/JngMkk/golang-fiber/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db := database.Connect()
	database.AutoMigrate(db)

	app := fiber.New()

	routes.GetUserRoutes(app)
	routes.GetSwaggerRoutes(app)

	middlewares.FiberMiddleWare(app)
	handlers.StartServer(app)
}

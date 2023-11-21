package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func FiberMiddleWare(a *fiber.App) {
	a.Use(
		cors.New(cors.Config{
			AllowOrigins: "*",
			AllowHeaders: "Origin, Content-Type, Accept, Authorization",
			AllowMethods: "GET, HEAD, PUT, PATCH, POST, DELETE",
		}),
		logger.New(),
	)
}

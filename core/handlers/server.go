package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func StartServer(a *fiber.App) {
	if err := a.Listen(":3001"); err != nil {
		log.Printf("Server is not Running, Reason: %v", err)
	}
}

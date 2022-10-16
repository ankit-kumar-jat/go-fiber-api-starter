package routes

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes() *fiber.App {
	app := fiber.New()

	registerUserRoutes(app)

	return app
}
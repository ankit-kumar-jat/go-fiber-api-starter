package routes

import (
	"example.com/go-api-starter/api/v2/controllers"

	"github.com/gofiber/fiber/v2"
)

func registerUserRoutes(app *fiber.App)  {
	users := app.Group("users")

	users.Post("/", controllers.AddUserController)
}
package routes

import (
	"example.com/go-api-starter/api/v1/controllers"

	"github.com/gofiber/fiber/v2"
)

func registerUserRoutes(app *fiber.App)  {
	users := app.Group("users")

	users.Get("/", controllers.GetUsersController)
	users.Post("/", controllers.AddUserController)
	users.Get("/:id", controllers.GetUserController)
}

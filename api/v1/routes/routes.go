package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	_ "example.com/go-api-starter/api/v1/docs"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /
func RegisterRoutes() *fiber.App {
	app := fiber.New()

	app.Get("/", apiHandler)
	app.Get("/docs/*" , swagger.HandlerDefault)

	registerUserRoutes(app)

	return app
}

func apiHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"name": "Go Starter API",
		"version": "v1",
	})
}
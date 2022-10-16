package main

import (
	"log"
	"strings"

	v1Router "example.com/go-api-starter/api/v1/routes"
	v2Router "example.com/go-api-starter/api/v2/routes"
	"example.com/go-api-starter/internal/config"
	"example.com/go-api-starter/internal/db"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	c, err := config.LoadConfig()

    if err != nil {
        log.Fatalln("Failed at config", err)
    }
	
	db.Init(&c)
    app := fiber.New()

	// Recover: for more infi visit https://docs.gofiber.io/api/middleware/recover
	app.Use(recover.New())

	// CORS: for more info visit https://docs.gofiber.io/api/middleware/cors
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	// Logger: for more info visit https://docs.gofiber.io/api/middleware/logger
	app.Use(logger.New())

	// Cache: for more info visit https://docs.gofiber.io/api/middleware/cache
	// app.Use(cache.New())

	// Rate Limiter: for more info visit https://docs.gofiber.io/api/middleware/limiter
	// app.Use(limiter.New())

	// Compress: for more info visit https://docs.gofiber.io/api/middleware/compress
	app.Use(compress.New(compress.Config{
		Next:  func(c *fiber.Ctx) bool {
			// do not compress ststic files 
			return strings.Contains( c.Path(), "/api/static") 
		  },
		Level: compress.LevelBestSpeed,
	}))

	// for more middlewares visit https://docs.gofiber.io/api/middleware

	app.Static("/api/static", "./public") 

    v1Routes := v1Router.RegisterRoutes()
    v2Routes := v2Router.RegisterRoutes()

	app.Mount("/api/v1", v1Routes)
	app.Mount("/api/v2", v2Routes)

    app.Listen(c.Port)
}
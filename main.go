package main

import (
	"time"

	"github.com/DevCorvus/gondor/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Init
	app := fiber.New()

	// Middlewares
	app.Use(limiter.New(limiter.Config{
		Max:        1,
		Expiration: time.Second,
	}))
	app.Use(logger.New())
	app.Use(etag.New())
	app.Use(encryptcookie.New(encryptcookie.Config{
		Key: encryptcookie.GenerateKey(), // Use a fixed one in production
	}))

	// Routes
	api := app.Group("/api")
	routes.SetupV1(api)

	// Run
	app.Listen(":8080")
}

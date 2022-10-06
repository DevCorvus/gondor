package main

import (
	"github.com/DevCorvus/gondor/features/healthcheck"
	"github.com/gofiber/fiber/v2"
)

var version = "0.1.0"

func main() {
	app := fiber.New()

	healthcheck.RegisterHandlers(app, version)

	app.Listen(":8080")
}

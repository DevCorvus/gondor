package main

import (
	"github.com/DevCorvus/gondor/api/healthcheck"
	"github.com/DevCorvus/gondor/api/users"
	"github.com/gofiber/fiber/v2"
)

var version = "0.1.0"

func main() {
	app := fiber.New()

	users.RegisterHandlers(app)
	healthcheck.RegisterHandlers(app, version)

	app.Listen(":8080")
}

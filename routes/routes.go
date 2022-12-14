package routes

import (
	"github.com/DevCorvus/gondor/api/auth"
	"github.com/DevCorvus/gondor/api/gophers"
	"github.com/DevCorvus/gondor/api/healthcheck"
	"github.com/DevCorvus/gondor/api/users"
	"github.com/gofiber/fiber/v2"
)

func SetupV1(api fiber.Router) {
	v1 := api.Group("/v1")

	auth.RegisterHandlers(v1)
	users.RegisterHandlers(v1)
	gophers.RegisterHandlers(v1)
	healthcheck.RegisterHandlers(v1, "0.1.0")
}

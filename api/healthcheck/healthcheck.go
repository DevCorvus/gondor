package healthcheck

import "github.com/gofiber/fiber/v2"

func RegisterHandlers(app *fiber.App, version string) {
	app.Get("/healthcheck", healthcheck(version))
}

func healthcheck(version string) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString(version)
	}
}

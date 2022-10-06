package healthcheck

import "github.com/gofiber/fiber/v2"

func RegisterHandlers(app *fiber.App, version string) {
	app.Get("/healthcheck", healthcheck(version))
}

func healthcheck(version string) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		c.Status(fiber.StatusOK)
		return c.SendString(version)
	}
}

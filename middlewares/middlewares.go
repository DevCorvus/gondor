package middlewares

import (
	"github.com/DevCorvus/gondor/database"
	"github.com/DevCorvus/gondor/database/models"
	"github.com/gofiber/fiber/v2"
)

var db = database.Conn

func UserIsAuthenticated(c *fiber.Ctx) error {
	userId := c.Cookies("session")

	if userId == "" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	if result := db.First(&models.User{}, userId); result.RowsAffected == 0 {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	return c.Next()
}

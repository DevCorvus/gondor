package users

import (
	"github.com/DevCorvus/gondor/database"
	"github.com/DevCorvus/gondor/database/models"
	"github.com/DevCorvus/gondor/utils"

	"github.com/gofiber/fiber/v2"
)

var db = database.Conn

func RegisterHandlers(app *fiber.App) {
	router := app.Group("/users")

	router.Post("/", create)
}

func create(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	errors := utils.ValidateStruct(user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	result := db.First(&models.User{}, "email = ?", user.Email)
	if result.RowsAffected > 0 {
		return c.SendStatus(fiber.StatusConflict)
	}

	db.Create(&user)

	return c.SendStatus(fiber.StatusCreated)
}

package users

import (
	"github.com/DevCorvus/gondor/database"
	"github.com/DevCorvus/gondor/database/models"
	"github.com/DevCorvus/gondor/utils"

	"github.com/gofiber/fiber/v2"
)

var db = database.Conn

func RegisterHandlers(r fiber.Router) {
	router := r.Group("/users")

	router.Post("/", create)
}

func create(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if errors := utils.ValidateStruct(user); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	result := db.First(&models.User{}, "email = ?", user.Email)
	if result.RowsAffected > 0 {
		return c.SendStatus(fiber.StatusConflict)
	}

	db.Create(&user)

	return c.SendStatus(fiber.StatusCreated)
}

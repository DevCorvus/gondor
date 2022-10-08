package auth

import (
	"strconv"
	"time"

	"github.com/DevCorvus/gondor/database"
	"github.com/DevCorvus/gondor/database/models"
	"github.com/DevCorvus/gondor/middlewares"
	"github.com/DevCorvus/gondor/utils"
	"github.com/gofiber/fiber/v2"
)

var db = database.Conn

func RegisterHandlers(r fiber.Router) {
	router := r.Group("/auth")

	router.Post("/login", login)
	router.Post("/logout", middlewares.UserIsAuthenticated, logout)
}

type loginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func login(c *fiber.Ctx) error {
	var body loginRequest

	if err := c.BodyParser(&body); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if errors := utils.ValidateStruct(body); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	var user models.User
	result := db.First(&user, "email = ?", body.Email)

	if result.Error != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	if !user.ComparePassword(body.Password) {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	session := fiber.Cookie{
		Name:     "session",
		Value:    strconv.FormatUint(uint64(user.ID), 10),
		Secure:   false,
		HTTPOnly: true,
		MaxAge:   int(time.Hour * 24),
	}

	c.Cookie(&session)

	return c.SendStatus(fiber.StatusOK)
}

func logout(c *fiber.Ctx) error {
	c.ClearCookie("session")
	return c.SendStatus(fiber.StatusNoContent)
}

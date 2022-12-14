package users

import (
	"strconv"

	"github.com/DevCorvus/gondor/database"
	"github.com/DevCorvus/gondor/database/models"
	"github.com/DevCorvus/gondor/middlewares"
	"github.com/DevCorvus/gondor/utils"

	"github.com/gofiber/fiber/v2"
)

var db = database.Conn

func RegisterHandlers(r fiber.Router) {
	router := r.Group("/users")

	router.Post("/", addUser)
	router.Get("/", middlewares.UserIsAuthenticated, getUser)
	router.Put("/", middlewares.UserIsAuthenticated, updateUser)
	router.Delete("/", middlewares.UserIsAuthenticated, deleteUser)
}

func addUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if errors := utils.ValidateStruct(user); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	if result := db.First(&models.User{}, "email = ?", user.Email); result.RowsAffected > 0 {
		return c.SendStatus(fiber.StatusConflict)
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	user.Password = hashedPassword

	if result := db.Omit("Gophers").Create(&user); result.Error != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusCreated)
}

func getUser(c *fiber.Ctx) error {
	userId := c.Cookies("session")

	var user models.User
	if result := db.First(&user, userId); result.Error != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	res := map[string]string{
		"id":        strconv.FormatUint(uint64(user.ID), 10),
		"name":      user.Name,
		"email":     user.Email,
		"createdAt": user.CreatedAt.String(),
		"updatedAt": user.UpdatedAt.String(),
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

type updateUserRequest struct {
	Name string `json:"name" validate:"required,min=2,max=100"`
}

func updateUser(c *fiber.Ctx) error {
	userId := c.Cookies("session")

	var body updateUserRequest

	if err := c.BodyParser(&body); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if errors := utils.ValidateStruct(body); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	var user models.User

	if result := db.Model(&user).Where(userId).Update("name", body.Name); result.Error != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusOK)
}

func deleteUser(c *fiber.Ctx) error {
	userId := c.Cookies("session")

	// Deleting by OnDelete:CASCADE constraint or Select(clauses.Associations) doesn't seem to work so...
	// This is not the best solution but it works:
	if result := db.Delete(&models.Gopher{}, "user_id = ?", userId); result.Error != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	if result := db.Delete(&models.User{}, userId); result.Error != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	c.ClearCookie("session")

	return c.SendStatus(fiber.StatusNoContent)
}

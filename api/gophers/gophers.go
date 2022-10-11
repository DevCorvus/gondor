package gophers

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
	router := r.Group("/gophers")

	router.Use(middlewares.UserIsAuthenticated)

	router.Post("/", addGopher)
	router.Get("/", getGophers)
	router.Get("/:gopherId", getGopher)
	router.Put("/:gopherId", updateGopher)
	router.Delete("/:gopherId", deleteGopher)
}

func addGopher(c *fiber.Ctx) error {
	userIdUint, err := strconv.ParseUint(c.Cookies("session"), 10, 32)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	var gopher models.Gopher

	if err := c.BodyParser(&gopher); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if errors := utils.ValidateStruct(gopher); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	if err := db.Model(&models.User{ID: uint(userIdUint)}).Association("Gophers").Append(&gopher); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusCreated)
}

func getGophers(c *fiber.Ctx) error {
	userIdUint, err := strconv.ParseUint(c.Cookies("session"), 10, 32)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	var user = models.User{ID: uint(userIdUint)}
	var gophers []models.Gopher

	if err := db.Model(&user).Association("Gophers").Find(&gophers); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(gophers)
}

func getGopher(c *fiber.Ctx) error {
	userIdUint, err := strconv.ParseUint(c.Cookies("session"), 10, 32)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	var user = models.User{ID: uint(userIdUint)}
	var gopher models.Gopher

	gopherId := c.Params("gopherId")
	if err := db.Model(&user).Association("Gophers").Find(&gopher, gopherId); err != nil || gopher.ID == 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.JSON(gopher)
}

type updateGopherRequest struct {
	Name string `json:"name" validate:"required,min=2,max=100"`
}

func updateGopher(c *fiber.Ctx) error {
	userIdUint, err := strconv.ParseUint(c.Cookies("session"), 10, 32)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	var body updateGopherRequest

	if err := c.BodyParser(&body); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if errors := utils.ValidateStruct(body); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	var user = models.User{ID: uint(userIdUint)}
	var gopher models.Gopher

	gopherId := c.Params("gopherId")
	if err := db.Model(&user).Association("Gophers").Find(&gopher, gopherId); err != nil || gopher.ID == 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}

	gopher.Name = body.Name

	if result := db.Save(&gopher); result.Error != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusOK)
}

func deleteGopher(c *fiber.Ctx) error {
	userIdUint, err := strconv.ParseUint(c.Cookies("session"), 10, 32)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	var user = models.User{ID: uint(userIdUint)}
	var gopher models.Gopher

	gopherId := c.Params("gopherId")
	if err := db.Model(&user).Association("Gophers").Find(&gopher, gopherId); err != nil || gopher.ID == 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}

	if result := db.Delete(&gopher); result.Error != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

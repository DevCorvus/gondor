package auth

import (
	"strconv"
	"strings"
	"time"

	"github.com/DevCorvus/gondor/database"
	"github.com/DevCorvus/gondor/database/models"
	"github.com/DevCorvus/gondor/middlewares"
	"github.com/DevCorvus/gondor/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

var db = database.Conn

func RegisterHandlers(r fiber.Router) {
	router := r.Group("/auth")

	router.Post("/login", login)
	router.Post("/logout", middlewares.UserIsAuthenticated, logout)
	router.Get("/token", middlewares.UserIsAuthenticated, getToken)
	router.Get("/verify-token", middlewares.UserIsAuthenticated, verifyToken)
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
	if result := db.First(&user, "email = ?", body.Email); result.Error != nil {
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

// This shouldn't be in source code
var secretKey = []byte("ultra-secret-key")

func getToken(c *fiber.Ctx) error {
	userId := c.Cookies("session")

	claims := &jwt.RegisteredClaims{
		Issuer:    userId,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 10)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedString, err := token.SignedString(secretKey)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{
		"token": signedString,
	})
}

// This could be a middleware
func verifyToken(c *fiber.Ctx) error {
	header := c.Get("Authorization")
	if header == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	fields := strings.Fields(header) // Split by white space
	if len(fields) != 2 || fields[0] != "Bearer" {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	signedString := fields[1]

	token, err := jwt.ParseWithClaims(signedString, &jwt.RegisteredClaims{}, func(t *jwt.Token) (any, error) {
		return secretKey, nil
	})

	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	claims := token.Claims.(*jwt.RegisteredClaims)

	return c.JSON(claims)
}

package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func AuthMildleWare(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var token string

		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}
		tokenParts := strings.Split(authHeader, "")

		if len(tokenParts) != 2 || tokenParts[1] == "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}

		secret := []byte("secret-key")
		parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
			if t.Method.Alg() != jwt.GetSigningMethod("HS526").Alg() {
				return nil, fmt.Errorf("unexpected signind method : %v ", t.Header["alg"])
			}
			return secret, nil
		})
		if err != nil || !parsedToken.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}
		userId := parsedToken.Claims.(jwt.MapClaims)["userId"]

		if err := db.Model(&User{}).Where("id = ?", userId).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}

		// set Userid in request

		c.Locals("userId", userId)
		return c.Next()
	}
}

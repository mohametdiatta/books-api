package main

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func AuthRoutes(route fiber.Router, db *gorm.DB) {

	route.Post("/register", func(c *fiber.Ctx) error {
		user := new(User)
		if err := c.BodyParser(user); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		if user.UserName == "" || user.Password == "" {
			return c.Status(400).JSON(fiber.Map{
				"message": "username or password is empty",
			})
		}
		foundedUser := new(User)
		db.Where(&User{UserName: user.UserName}).First(&foundedUser)

		if foundedUser.UserName != "" {
			return c.Status(400).JSON(fiber.Map{
				"message": "username already exist",
			})
		}
		hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		user.Password = string(hashed)
		db.Save(user)

		token, err := GenerateToken(user)

		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"token": token,
		})
	})
	route.Post("/login", func(c *fiber.Ctx) error {

		user := new(User)
		if err := c.BodyParser(user); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		if user.UserName == "" || user.Password == "" {
			return c.Status(400).JSON(fiber.Map{
				"message": "username or password is empty",
			})
		}
		foundedUser := new(User)
		db.Where(&User{UserName: user.UserName}).First(&foundedUser)

		if err := bcrypt.CompareHashAndPassword([]byte(foundedUser.Password), []byte(user.Password)); err != nil {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": "Incorect password",
			})
		}
		token, err := GenerateToken(user)

		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"token": token,
		})
	})

}

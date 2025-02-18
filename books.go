package main

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func BookRoutes(route fiber.Router, db *gorm.DB) {
	route.Get("/", func(c *fiber.Ctx) error {
		books := []Book{}
		db.Find(&books)
		return c.JSON(books)
	})
	route.Post("/", func(c *fiber.Ctx) error {
		book := new(Book)
		book.UserId = int(c.Locals("userId").(float64))
		if err := c.BodyParser(book); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		db.Save(book)
		return c.JSON(book)

	})
}

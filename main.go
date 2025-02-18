package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	db := InitDb()
	app := fiber.New(fiber.Config{
		AppName: "Book Api",
	})

	AuthRoutes(app.Group("/auth"), db)

	// midleware ...
	protected := app.Use(AuthMildleWare(db))

	BookRoutes(protected.Group("/books"), db)
	app.Listen(":3000")

}

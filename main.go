package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Start a new fiber app
	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("Todo Server is up and running!")
	})

	app.Get("/todos", func(c *fiber.Ctx) error {
		return c.SendString("GET todo list!")
	})

	// Listen on PORT 300
	app.Listen(":3000")
}

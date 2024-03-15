package main

import (
	"github.com/chips03/go-todo-lite/models"
	"github.com/gofiber/fiber/v2"
)

func main() {
	//Initialize database
	models.ConnectDatabase()

	// Start a new fiber app
	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "ok",
			"message": "TODO API Server is up and running",
		})
	})

	app.Get("/todos", func(c *fiber.Ctx) error {
		return c.SendString("GET todo list!")
	})

	// Listen on PORT 300
	app.Listen(":3000")
}

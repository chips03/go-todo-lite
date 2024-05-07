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

	// Get all todos
	app.Get("/todos", func(c *fiber.Ctx) error {
		var todos []models.Todo
		models.DB.Find(&todos)

		return c.JSON(todos)
	})

	// Get a single todo
	app.Get("/todos/:id", func(c *fiber.Ctx) error {
		var todo models.Todo
		result := models.DB.First(&todo, c.Params("id"))

		if result.Error != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Todo not found",
			})
		}

		return c.JSON(todo)
	})

	// Create a new todo
	app.Post("/todos", func(c *fiber.Ctx) error {
		todo := new(models.Todo)
		if err := c.BodyParser(todo); err != nil {
			return c.Status(503).SendString(err.Error())
		}

		if err := models.DB.Create(&todo).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(todo)
	})

	// Update a todo
	app.Put("/todos/:id", func(c *fiber.Ctx) error {
		todo := new(models.Todo)
		if err := c.BodyParser(todo); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		id := c.Params("id")
		models.DB.Model(&todo).Where("id = ?", id).Updates(&todo)
		return c.JSON(todo)
	})

	// Listen on PORT 300
	app.Listen(":3000")
}

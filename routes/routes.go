package routes

import (
	// "golab/controllers"
	// "golab/home"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	app.Post("/api/register", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "success",
		})
	})
	app.Post("/api/login", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "success",
		})
	})
	app.Get("/api/user", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "success",
		})
	})
	app.Post("/api/logout", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "success",
		})
	})
	app.Get("/api/home", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "success",
		})
	})
}

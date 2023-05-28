package home

import (
	"github.com/gofiber/fiber"
)

func Home(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "success",
	})
}

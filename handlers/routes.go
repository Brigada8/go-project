package handlers

import (
	"golab/internal/weather/services"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	app.Post("/api/register", services.Register)
	app.Post("/api/login", services.Login)
	app.Get("/api/logout", services.Logout)
	app.Post("/api/weather", services.Weather)
	app.Get("/api/user", services.User)
}

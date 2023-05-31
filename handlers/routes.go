package handlers

import (
	"golab/internal/weather"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	app.Post("/api/register", Register)
	app.Post("/api/login", Login)
	app.Get("/api/logout", Logout)
	app.Post("/api/weather", weather.Weather)
	app.Get("/api/user", User)
}

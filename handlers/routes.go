package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	app.Post("/api/register")
	app.Post("/api/login", handlers.Login)
	app.Get("/api/logout", handlers.Logout)
	app.Post("/api/weather", handlers.Weather)
	app.Get("/api/user", handlers.User)
}

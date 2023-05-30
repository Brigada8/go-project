package routes

import (
	"golab/home"

	"golab/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/logout", controllers.Logout)
	app.Get("/api/weather", home.Home)
	app.Get("/api/user", controllers.User)
}

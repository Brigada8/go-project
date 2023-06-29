package handlers

import (
	"golab/internal/weather/repositories"
	"golab/internal/weather/services/AuthServices"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	Repo := repositories.NewUserRepository(repositories.DB)
	Services := AuthServices.NewAuthService(Repo)
	Handler := NewHttpHandler(Services)

	app.Post("/api/register", Handler.Register)
	app.Post("/api/login", Handler.Login)
	app.Get("/api/logout", Handler.Logout)
	app.Post("/api/weather", Handler.Weather)
	app.Get("/api/user", Handler.User)
}

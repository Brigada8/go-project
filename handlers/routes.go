package handlers

import (
	"golab/internal/weather/repositories"
	"golab/internal/weather/services/AuthServices"
	"golab/internal/weather/services/WeatherServices"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	UserRepo := repositories.NewUserRepository(repositories.DB)
	WeatherRepo := repositories.NewWeatherRepository(repositories.DB)
	WeatherServices := WeatherServices.NewWeatherService(WeatherRepo)
	UserServices := AuthServices.NewAuthService(UserRepo)
	WeatherHandler := NewWeatherHandler(WeatherServices)
	Handler := NewHttpHandler(UserServices)

	app.Post("/api/register", Handler.Register)
	app.Post("/api/login", Handler.Login)
	app.Get("/api/logout", Handler.Logout)
	app.Post("/api/weather", WeatherHandler.Weather)
	app.Get("/api/user", Handler.User)
	app.Post("/api/forecast", WeatherHandler.Forecast)
	app.Get("/api/history", WeatherHandler.History)
	app.Post("/api/astronomy", WeatherHandler.Astronomy)
}

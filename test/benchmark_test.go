package handlers_test

import (
	"golab/handlers"
	"golab/internal/weather/repositories"
	"golab/internal/weather/services/AuthServices"
	"golab/internal/weather/services/WeatherServices"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestHttpHandler_Register(t *testing.T) {
	repositories.Connect()
	Repo := repositories.NewUserRepository(repositories.DB)
	Services := AuthServices.NewAuthService(Repo)
	handler := handlers.NewHttpHandler(Services)

	app := fiber.New()
	app.Post("/api/register", handler.Register)

	name := "Johny Doe"
	email := "johyn@example.com"
	password := "password123"

	payload := `{"name":"` + name + `","email":"` + email + `", "password":"` + password + `"}`

	req := httptest.NewRequest(http.MethodPost, "/api/register", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestHttpHandler_Login(t *testing.T) {
	repositories.Connect()
	Repo := repositories.NewUserRepository(repositories.DB)
	Services := AuthServices.NewAuthService(Repo)
	handler := handlers.NewHttpHandler(Services)

	app := fiber.New()
	app.Post("/api/login", handler.Login)

	email := "john@example.com"
	password := "password123"

	payload := `{"email":"` + email + `", "password":"` + password + `"}`

	req := httptest.NewRequest(http.MethodPost, "/api/login", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

}

func TestHttpHandler_UserUnauthorized(t *testing.T) {
	repositories.Connect()
	Repo := repositories.NewUserRepository(repositories.DB)
	Services := AuthServices.NewAuthService(Repo)
	handler := handlers.NewHttpHandler(Services)

	app := fiber.New()
	app.Post("/api/user", handler.User)

	req := httptest.NewRequest(http.MethodPost, "/api/user", nil)
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)

}

func TestHttpHandler_Weather(t *testing.T) {
	repositories.Connect()
	WeatherRepo := repositories.NewWeatherRepository(repositories.DB)
	WeatherServices := WeatherServices.NewWeatherService(WeatherRepo)
	WeatherHandler := handlers.NewWeatherHandler(WeatherServices)

	app := fiber.New()
	app.Post("/api/weather", WeatherHandler.Weather)

	loc := "Kiev"

	payload := `{"loc":"` + loc + `"}`

	req := httptest.NewRequest(http.MethodPost, "/api/weather", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}

func TestHttpHandler_History(t *testing.T) {
	repositories.Connect()
	WeatherRepo := repositories.NewWeatherRepository(repositories.DB)
	WeatherServices := WeatherServices.NewWeatherService(WeatherRepo)
	WeatherHandler := handlers.NewWeatherHandler(WeatherServices)
	Repo := repositories.NewUserRepository(repositories.DB)
	Services := AuthServices.NewAuthService(Repo)
	handler := handlers.NewHttpHandler(Services)

	app := fiber.New()
	app.Post("/api/login", handler.Login)
	app.Get("/api/history", WeatherHandler.History)

	email := "john@example.com"
	password := "password123"

	payload := `{"email":"` + email + `", "password":"` + password + `"}`

	req := httptest.NewRequest(http.MethodPost, "/api/login", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)
	jwt := resp.Header.Get("Set-Cookie")

	req_hst := httptest.NewRequest(http.MethodGet, "/api/history", strings.NewReader(payload))
	req_hst.Header.Set("Content-Type", "application/json")
	req_hst.Header.Set("Cookie", jwt)


	hst_resp, _ := app.Test(req_hst)
	
	assert.Equal(t, http.StatusOK, hst_resp.StatusCode)
}


func TestHttpHandler_Forecast(t *testing.T) {
	repositories.Connect()
	WeatherRepo := repositories.NewWeatherRepository(repositories.DB)
	WeatherServices := WeatherServices.NewWeatherService(WeatherRepo)
	WeatherHandler := handlers.NewWeatherHandler(WeatherServices)

	app := fiber.New()
	app.Post("/api/forecast", WeatherHandler.Forecast)

	loc := "Lviv"

	payload := `{"loc":"` + loc + `"}`

	req := httptest.NewRequest(http.MethodPost, "/api/forecast", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
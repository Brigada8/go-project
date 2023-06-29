package handlers_test

import (
	"bytes"
	"encoding/json"
	"golab/internal/weather/repositories"
	"golab/internal/weather/services/AuthServices"
	"net/http"
	"net/http/httptest"
	"testing"

	"golab/handlers"
	"golab/internal/weather"

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

	payload := map[string]string{
		"name":     "John Doe",
		"email":    "john@example.com",
		"password": "password123",
	}
	jsonPayload, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPost, "/api/register", bytes.NewBuffer(jsonPayload))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var user weather.User
	json.NewDecoder(resp.Body).Decode(&user)

	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "john@example.com", user.Email)
}

func TestHttpHandler_Login(t *testing.T) {
	repositories.Connect()
	Repo := repositories.NewUserRepository(repositories.DB)
	Services := AuthServices.NewAuthService(Repo)
	handler := handlers.NewHttpHandler(Services)

	app := fiber.New()
	app.Post("/ap/login", handler.Login)

	// Login with registered user
	loginPayload := map[string]string{
		"email":    "john@example.com",
		"password": "password123",
	}
	loginPayloadBytes, _ := json.Marshal(loginPayload)

	loginReq := httptest.NewRequest(http.MethodPost, "/api/login", bytes.NewBuffer(loginPayloadBytes))
	loginReq.Header.Set("Content-Type", "application/json")

	loginResp, err := app.Test(loginReq)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, loginResp.StatusCode)

	var response map[string]string
	json.NewDecoder(loginResp.Body).Decode(&response)

	assert.Equal(t, "success", response["message"])
}

package handlers_test

import (
	"fmt"
	"golab/handlers"
	"golab/internal/weather/repositories"
	"golab/internal/weather/services/AuthServices"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

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

	// Создаем тестовый ответ
	resp, _ := app.Test(req)

	fmt.Println("")
	fmt.Println(resp)
	fmt.Println("")

	// Выполняем запрос к тестовому серверу

	// Проверяем код ответа
	assert.Equal(t, http.StatusOK, resp.StatusCode)

}

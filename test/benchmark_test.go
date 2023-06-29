package handlers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golab/handlers"
	"golab/internal/weather/repositories"
	"golab/internal/weather/services/AuthServices"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/steinfletcher/apitest"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
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
}

func TestHttpHandler_Login(t *testing.T) {
	repositories.Connect()
	Repo := repositories.NewUserRepository(repositories.DB)
	Services := AuthServices.NewAuthService(Repo)
	handler := handlers.NewHttpHandler(Services)

	app := fiber.New()
	app.Post("/api/login", handler.Login)

	testHandler := FiberToHandler(app)

	apitest.New(). // configuration
			HandlerFunc(testHandler).
			Post("/api/login"). // request
			Expect(t).          // expectations
			Assert(jsonpath.Equal(`$`, map[string]interface{}{"name": "John Doe", "email": "john@example.com", "password": "password123"})).
			End()
}
func FiberToHandler(app *fiber.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := app.Test(r)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Response Body: %s\n", string(body))

		if len(body) == 0 {
			panic("Empty response body")
		}

		var jsonBody map[string]interface{}
		if err := json.Unmarshal(body, &jsonBody); err != nil {
			panic(err)
		}

		for k, vv := range resp.Header {
			for _, v := range vv {
				w.Header().Add(k, v)
			}
		}
		w.WriteHeader(resp.StatusCode)

		if _, err := io.Copy(w, resp.Body); err != nil {
			panic(err)
		}
	}
}

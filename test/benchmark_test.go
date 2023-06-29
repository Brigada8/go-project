package login_test

import (
	"fmt"
	"golab/internal/weather/repositories"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/steinfletcher/apitest"
	"github.com/stretchr/testify/assert"
)

func BenchmarkLogin(b *testing.B) {
	repositories.Connect()

	err := os.Setenv("JWT_SECRET", "test")
	assert.NoError(b, err)
	app := fiber.New()

	for i := 0; i < b.N; i++ {
		apitest.New().
			HandlerFunc(FiberToHandlerFunc(app)).
			Post("/api/login").
			JSON(fmt.Sprintf(`{"email": "a%d@b.c", "password": "11111111"}`, i)).
			Expect(b).
			Status(http.StatusUnauthorized).
			End()
	}
}

func FiberToHandlerFunc(app *fiber.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := app.Test(r)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		// copy headers
		for k, vv := range resp.Header {
			for _, v := range vv {
				w.Header().Add(k, v)
			}
		}
		w.WriteHeader(resp.StatusCode)

		// copy body
		if _, err := io.Copy(w, resp.Body); err != nil {
			panic(err)
		}
	}
}

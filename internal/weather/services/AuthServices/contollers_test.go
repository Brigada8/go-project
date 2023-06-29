package AuthServices

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestFiberGet(t *testing.T) {
	server := fiber.New()

	server.Get("/api/user", User)

	req, _ := http.NewRequest(http.MethodGet, "/api/user", nil)
	resp, _ := server.Test(req, -1)

	t.Log(resp.StatusCode)
	t.Log(resp.Body)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestFiberPost(t *testing.T) {
	server := fiber.New()

	input := struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}{
		Name:     "Artem",
		Email:    "artem@gmail.com",
		Password: "123456",
	}

	bodyReq, _ := json.Marshal(input)

	server.Post("/api/register", Register)

	req, _ := http.NewRequest(http.MethodPost, "/api/register", bytes.NewBuffer(bodyReq))
	resp, _ := server.Test(req, -1)

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	t.Log(resp.StatusCode)
	t.Log(string(body))
	assert.Equal(t, 200, resp.StatusCode)
}

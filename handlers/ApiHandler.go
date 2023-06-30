package handlers

import (
	"encoding/json"
	"fmt"
	Models "golab/internal/weather"
	"net/http"
	"os"

	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type WeatherService interface {
	AddToHistory(c *fiber.Ctx, history Models.History) (string, error)
}

type WeatherHandler struct {
	weatherService WeatherService
}

func NewWeatherHandler(weatherService WeatherService) *WeatherHandler {
	return &WeatherHandler{
		weatherService: weatherService,
	}
}


func (h *WeatherHandler) Weather(c *fiber.Ctx) error {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	endpoint := "http://api.weatherapi.com/v1/current.json"
	apiKey := os.Getenv("API")

	type WeatherResponse struct {
		Location struct {
			Name    string `json:"name"`
			Country string `json:"country"`
		} `json:"location"`
		Current struct {
			TempC     float64 `json:"temp_c"`
			Condition struct {
				Text string `json:"text"`
				Icon string `json:"icon"`
			} `json:"condition"`
		} `json:"current"`
	}

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	client := &http.Client{}

	request, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		fmt.Println("Failed to create request:", err)
		return err
	}

	query := request.URL.Query()
	query.Add("key", apiKey)
	query.Add("q", c.Query("location", data["loc"]))
	request.URL.RawQuery = query.Encode()

	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Request failed:", err)
		return err
	}
	defer response.Body.Close()

	// Read the response body
	var weatherResp WeatherResponse
	err = json.NewDecoder(response.Body).Decode(&weatherResp)
	if err != nil {
		fmt.Println("Failed to read response body:", err)
		return err
	}

	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	user, _ := strconv.Atoi(claims.Issuer)

	weather := Models.History{
		UserID: uint(user),
		Location: weatherResp.Location.Country,
	}

	h.weatherService.AddToHistory(c, weather)

	// Set the response body as the JSON result
	return c.JSON(weatherResp)
}
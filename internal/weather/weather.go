package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func Weather(c *fiber.Ctx) error {
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

	// Set the response body as the JSON result
	return c.JSON(weatherResp)
}

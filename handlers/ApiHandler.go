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
	FindHistoryByID(c *fiber.Ctx, claims *jwt.StandardClaims) ([]Models.History, error)
	
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
		UserID:   uint(user),
		Location: weatherResp.Location.Country,
		City: weatherResp.Location.Name,
	}

	if _, err := h.weatherService.AddToHistory(c, weather); err != nil {
		fmt.Println(err)
	}

	// Set the response body as the JSON result
	return c.JSON(weatherResp)
}


func (h *WeatherHandler) History(c *fiber.Ctx) error {
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

	var history []Models.History

	if history, err = h.weatherService.FindHistoryByID(c, claims); err != nil {
		fmt.Println(err)
	}

	return c.JSON(history)
}

func (h *WeatherHandler) Forecast(c *fiber.Ctx) error {
	days := "3"
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	endpoint := "http://api.weatherapi.com/v1/forecast.json"
	apiKey := os.Getenv("API")

	type WeatherResponse struct {
		Location struct {
			Name    string `json:"name"`
			Country string `json:"country"`
		} `json:"location"`
		Forecast struct {
			ForecastDay []struct {
				Date string `json:"date"`
				Day struct {
					MaxTemp float64 `json:"maxtemp_c"`
					AvgTemp float64 `json:"avgtemp_c"`
				Condition struct {
					Text string `json:"text"`
					Icon string `json:"icon"`
				} `json:"condition"`
				} `json:"day"`
			} `json:"forecastday"`
		} `json:"forecast"`
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

	if data["loc"] == ""{
		data["loc"] = GetRealIP(request)
	}



	query := request.URL.Query()
	query.Add("key", apiKey)
	query.Add("q", c.Query("location", data["loc"]))
	query.Add("days", c.Query("days", days))
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

func (h *WeatherHandler) Astronomy(c *fiber.Ctx) error {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	endpoint := "http://api.weatherapi.com/v1/astronomy.json"
	apiKey := os.Getenv("API")

	type AstronomyResponse struct {
		Location struct {
			Name    string `json:"name"`
			Country string `json:"country"`
		} `json:"location"`
		Astronomy struct {
			Astro struct {
				Sunrise string `json:"sunrise"`
				Sunset string `json:"sunset"`
				Moonrise string `json:"moonrise"`
				Moonset string `json:"moonset"`
				Moon_phase string `json:"moon_phase"`
				} `json:"astro"`
				} `json:"astronomy"`
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

	if data["loc"] == ""{
		data["loc"] = GetRealIP(request)
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
	var astronomyResp AstronomyResponse
	err = json.NewDecoder(response.Body).Decode(&astronomyResp)
	if err != nil {
		fmt.Println("Failed to read response body:", err)
		return err
	}


	// Set the response body as the JSON result
	return c.JSON(astronomyResp)
}


func GetRealIP(r *http.Request) string {
    IPAddress := r.Header.Get("X-Real-IP")
    if IPAddress == "" {
        IPAddress = r.Header.Get("X-Forwarder-For")
    }
    if IPAddress == "" {
        IPAddress = r.RemoteAddr
    }
    return IPAddress
}
package WeatherServices

import (
	Models "golab/internal/weather"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

type WeatherRepository interface {
	AddToHistory(c *fiber.Ctx, history Models.History) (string, error)
	FindHistoryByID(c *fiber.Ctx, claims *jwt.StandardClaims) ([]Models.History, error)
}

type weatherService struct {
	weatherRepository WeatherRepository
}

func NewWeatherService(weatherRepository WeatherRepository) *weatherService {
	return &weatherService{
		weatherRepository: weatherRepository,
	}
}


func (s *weatherService) AddToHistory(c *fiber.Ctx, history Models.History) (string, error) {
	return s.weatherRepository.AddToHistory(c, history)
}

func (s *weatherService) FindHistoryByID(c *fiber.Ctx, claims *jwt.StandardClaims) ([]Models.History, error) {
	return s.weatherRepository.FindHistoryByID(c, claims)
}
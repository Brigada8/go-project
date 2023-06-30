package repositories

import (
	Models "golab/internal/weather"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type WeatherRepository struct {
	DB *gorm.DB
}

func NewWeatherRepository(DB *gorm.DB) *WeatherRepository {
	return &WeatherRepository{DB}
}

func (o *WeatherRepository) AddToHistory(c *fiber.Ctx, history Models.History) (string, error) {
	DB.Create(&history)
	return "", c.JSON(history)
}
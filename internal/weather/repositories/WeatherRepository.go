package repositories

import (
	Models "golab/internal/weather"

	"github.com/dgrijalva/jwt-go"
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

func (o *WeatherRepository) FindHistoryByID(c *fiber.Ctx, claims *jwt.StandardClaims) (Models.History, error) {
	var history Models.History
	DB.Where("user_id = ?", claims.Issuer).Find(&history)
	return history, nil
}
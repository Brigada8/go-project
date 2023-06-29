package AuthServices

import (
	Models "golab/internal/weather"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

type AuthRepository interface {
	CreateUser(c *fiber.Ctx, user Models.User) (string, error)
	FindUserByEmail(c *fiber.Ctx, data map[string]string) (Models.User, error)
	FindUserByID(c *fiber.Ctx, claims *jwt.StandardClaims) (Models.User, error)
}

type AuthService struct {
	authRepository AuthRepository
}

func NewAuthService(authRepository AuthRepository) *AuthService {
	return &AuthService{
		authRepository: authRepository,
	}
}

func (s *AuthService) AddUser(c *fiber.Ctx, user Models.User) (string, error) {
	return s.authRepository.CreateUser(c, user)
}

func (s *AuthService) GetUserByID(c *fiber.Ctx, claims *jwt.StandardClaims) (Models.User, error) {
	return s.authRepository.FindUserByID(c, claims)
}

func (s *AuthService) GetUserByEmail(c *fiber.Ctx, data map[string]string) (Models.User, error) {
	return s.authRepository.FindUserByEmail(c, data)
}

func (s *AuthService) Destroy(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

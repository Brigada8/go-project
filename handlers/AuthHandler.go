package handlers

import (
	"fmt"
	"golab/internal/weather"
	Models "golab/internal/weather"
	"golab/internal/weather/repositories"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

const SecretKey = "secret"

type HttpHandler struct {
	authService AuthService
}

type AuthService interface {
	AddUser(c *fiber.Ctx, user Models.User) (string, error)
	GetUserByID(c *fiber.Ctx, claims *jwt.StandardClaims) (Models.User, error)
	GetUserByEmail(c *fiber.Ctx, data map[string]string) (Models.User, error)
	Destroy(c *fiber.Ctx) error
}

func NewHttpHandler(authService AuthService) *HttpHandler {
	return &HttpHandler{
		authService: authService,
	}
}

func (h *HttpHandler) Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := Models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}

	if _, err := h.authService.AddUser(c, user); err != nil {
		fmt.Println(err)
	}

	return c.JSON(user)
}

func (h *HttpHandler) Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user weather.User

	repositories.DB.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //1 day
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "could not login",
		})
	}

	cookie := fiber.Cookie{
		Name:    "jwt",
		Value:   token,
		Expires: time.Now().Add(time.Hour * 24),
		Domain:  "gofront.onrender.com",
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func (h *HttpHandler) User(c *fiber.Ctx) error {
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

	var user weather.User

	if user, err = h.authService.GetUserByID(c, claims); err != nil {
		fmt.Println(err)
	}

	return c.JSON(user)
}
func (h *HttpHandler) Logout(c *fiber.Ctx) error {
	if err := h.authService.Destroy(c); err != nil {
		fmt.Println(err)
	}
	return nil
}

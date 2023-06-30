package repositories

import (
	"fmt"
	Models "golab/internal/weather"
	internal "golab/internal/weather"

	"github.com/dgrijalva/jwt-go"
	"github.com/glebarez/sqlite"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) *UserRepository {
	return &UserRepository{DB}
}

func (o *UserRepository) CreateUser(c *fiber.Ctx, user Models.User) (string, error) {
	DB.Create(&user)
	return "", c.JSON(user)
}

func (o *UserRepository) FindUserByID(c *fiber.Ctx, claims *jwt.StandardClaims) (Models.User, error) {
	var user Models.User
	DB.Where("id = ?", claims.Issuer).First(&user)
	return user, nil
}

func (o *UserRepository) FindUserByEmail(c *fiber.Ctx, data map[string]string) (Models.User, error) {
	var user Models.User
	DB.Where("email = ?", data["email"]).First(&user)
	return user, nil
}

func (o *UserRepository) Connect() {
	var err error
	connection, err := gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB = connection

	// Migrate the schema
	DB.AutoMigrate(
		&internal.User{},
	&internal.History{})

	fmt.Println("Successfully connected!", DB)
}

package main

import (
	"fmt"
	"golab/handlers"
	"golab/internal/weather/repositories"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	repositories.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "https://gofront.onrender.com, https://35.160.120.126, https://44.233.151.27, https://34.211.200.85",
	}))
	handlers.Setup(app)

	if err := app.Listen(":8000"); err != nil {
		fmt.Println(err)
	}

}

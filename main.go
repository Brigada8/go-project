package main

import (
	"golab/handlers"
	"golab/internal"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	internal.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "https://gofront.onrender.com, https://35.160.120.126, https://44.233.151.27, https://34.211.200.85",
	}))
	handlers.Setup(app)

	app.Listen(":8000")
}

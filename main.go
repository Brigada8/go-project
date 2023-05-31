package main

import (
	"golab/routes"
	"golab/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	utils.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "*",
	}))
	routes.Setup(app)

	app.Listen(":8000")
}

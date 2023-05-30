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
		AllowOrigins:     "https://gofront.onrender.com, 35.160.120.126, 44.233.151.27, 34.211.200.85",
		AllowHeaders:     "Origin, Content-Type, Accept",
	}))
	routes.Setup(app)

	app.Listen(":8000")
}

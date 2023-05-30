package main

import (
	"fmt"
	"golab/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func main() {
	utils.Connect()

	app := fiber.New()

	app.Listen(":8000")
}

func HelloHandler(w http.ResponseWriter, _ *http.Request) {

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://gofront.onrender.com/")
	w.Header().Set("Access-Control-Max-Age", "15")
	fmt.Fprintf(w, "Hello, there!")
}

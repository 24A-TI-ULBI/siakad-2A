package main

import (
	"log"

	"backend/config"
	"backend/url"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Load .env
	config.InitConfig()

	// Init Fiber
	app := fiber.New(config.FiberConfig)

	// CORS
	app.Use(cors.New(config.CorsConfig()))

	// Routes API
	url.Web(app)

	log.Printf("Server running on %s", config.IPPort)
	log.Fatal(app.Listen(config.IPPort))
}

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

	// Init Fiber dengan config mengikuti pola boilerplate gocroot
	app := fiber.New(config.FiberConfig)

	// CORS
	app.Use(cors.New(config.CorsConfig()))

	// Routes
	url.Web(app)

	// Static files frontend
	app.Static("/", "./frontend")

	log.Printf("Server running on %s", config.IPPort)
	log.Fatal(app.Listen(config.IPPort))
}

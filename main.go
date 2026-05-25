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

	// Static files frontend
	app.Static("/frontend", "./frontend")
	app.Get("/berita", func(c *fiber.Ctx) error {
		return c.SendFile("./frontend/berita/index.html")
	})
	app.Get("/admin/berita", func(c *fiber.Ctx) error {
		return c.SendFile("./frontend/admin/berita/index.html")
	})

	// Routes
	url.Web(app)

	log.Printf("Server running on %s", config.IPPort)
	log.Fatal(app.Listen(config.IPPort))
}

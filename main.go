package main

import (
	"embed"
	"log"
	"net/http"

	"backend/config"
	"backend/url"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

//go:embed frontend
var frontendFS embed.FS

func main() {
	// Load .env
	config.InitConfig()

	// Init Fiber
	app := fiber.New(config.FiberConfig)

	// CORS
	app.Use(cors.New(config.CorsConfig()))

	// Routes API (didaftarkan sebelum static agar tidak tertimpa)
	url.Web(app)

	// Static files frontend — embed ke binary
	app.Use("/", filesystem.New(filesystem.Config{
		Root:       http.FS(frontendFS),
		PathPrefix: "frontend",
		Browse:     false,
		Index:      "index.html",
	}))

	log.Printf("Server running on %s", config.IPPort)
	log.Fatal(app.Listen(config.IPPort))
}

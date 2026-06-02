package url

import (
	"backend/controller"

	"github.com/gofiber/fiber/v2"
)

func BeasiswaRoute(app *fiber.App) {
	// Without /api prefix
	app.Get("/beasiswa", controller.GetBeasiswa)
	app.Get("/beasiswa/:id", controller.GetDetailBeasiswa)
	app.Post("/beasiswa", controller.AddBeasiswa)
	app.Put("/beasiswa/status/:npm", controller.UpdateBeasiswa)
	app.Delete("/beasiswa/:npm", controller.DeleteBeasiswa)

	// With /api prefix
	app.Get("/api/beasiswa", controller.GetBeasiswa)
	app.Get("/api/beasiswa/:id", controller.GetDetailBeasiswa)
	app.Post("/api/beasiswa", controller.AddBeasiswa)
	app.Put("/api/beasiswa/status/:npm", controller.UpdateBeasiswa)
	app.Delete("/api/beasiswa/:npm", controller.DeleteBeasiswa)
}

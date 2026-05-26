package url

import (
	"backend/controller" //

	"github.com/gofiber/fiber/v2"
)

func PKLRoute(app *fiber.App) {
	// Backend 1: PKL Routing
	app.Get("/pkl", controller.GetAllPKL)
	app.Get("/pkl/:npm", controller.GetPKLByNPM)
	app.Post("/pkl", controller.CreatePKL)
	app.Put("/pkl/:id", controller.UpdatePKL)
	app.Delete("/pkl/:id", controller.DeletePKL)

	// Backend 2: Laporan PKL Routing
	app.Post("/pkl/laporan", controller.SubmitLaporan)
	app.Get("/pkl/laporan/:npm", controller.GetLaporanByNPM)
	app.Put("/pkl/laporan/:id", controller.UpdateLaporan)
	app.Get("/pkl/laporan/:id/nilai", controller.GetNilaiPKL)
}
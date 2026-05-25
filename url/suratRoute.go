package url

import (
	"backend/controller"

	"github.com/gofiber/fiber/v2"
)

// SuratRoute registers all REST endpoints for Modul 16
func SuratRoute(app *fiber.App) {
	// Surat Templates
	app.Get("/surat", controller.GetTemplates)
	app.Get("/surat/:id", controller.GetTemplateByID)
	app.Post("/surat", controller.CreateTemplate)
	app.Put("/surat/:id", controller.UpdateTemplate)
	app.Delete("/surat/:id", controller.DeleteTemplate)

	// Pengajuan Surat
	app.Post("/surat/ajukan", controller.SubmitPengajuan)
	app.Get("/surat/pengajuan/:npm", controller.GetPengajuanByNPM)
	app.Put("/surat/pengajuan/:id", controller.UpdateStatusPengajuan)
	app.Get("/surat/pengajuan/status/:status", controller.GetPengajuanByStatus)
}

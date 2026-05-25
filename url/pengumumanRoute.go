package url

import (
	"backend/controller"

	"github.com/gofiber/fiber/v2"
)

// PengumumanRoute mendaftarkan semua endpoint modul 7 ke aplikasi Fiber
func PengumumanRoute(app *fiber.App) {
	// Endpoint Pengumuman
	app.Get("/pengumuman", controller.GetPengumuman)
	app.Get("/pengumuman-data", controller.GetPengumuman)
	app.Get("/pengumuman/:id", controller.GetPengumumanByID)
	app.Post("/pengumuman", controller.CreatePengumuman)
	app.Post("/pengumuman/upload", controller.UploadFile)
	app.Put("/pengumuman/:id", controller.UpdatePengumuman)
	app.Delete("/pengumuman/:id", controller.DeletePengumuman)

	// Endpoint Kategori
	app.Get("/kategori", controller.GetKategori)
	app.Post("/kategori", controller.CreateKategori)
	app.Delete("/kategori/:id", controller.DeleteKategori)
	app.Get("/pengumuman/kategori/:nama", controller.GetPengumumanByKategori)
}

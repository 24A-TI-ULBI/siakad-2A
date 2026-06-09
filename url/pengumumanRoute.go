package url

import (
	"backend/controller"

	"github.com/gofiber/fiber/v2"
)

// PengumumanRoute mendaftarkan semua endpoint modul 7 ke aplikasi Fiber
func PengumumanRoute(app *fiber.App) {
	registerPengumumanRoutes(app, "")
	registerPengumumanRoutes(app, "/api")
}

func registerPengumumanRoutes(app *fiber.App, prefix string) {
	// Endpoint Pengumuman
	app.Get(prefix+"/pengumuman", controller.GetPengumuman)
	app.Get(prefix+"/pengumuman-data", controller.GetPengumuman)
	app.Get(prefix+"/pengumuman/:id", controller.GetPengumumanByID)
	app.Post(prefix+"/pengumuman", controller.CreatePengumuman)
	app.Post(prefix+"/pengumuman/upload", controller.UploadFile)
	app.Put(prefix+"/pengumuman/:id", controller.UpdatePengumuman)
	app.Delete(prefix+"/pengumuman/:id", controller.DeletePengumuman)

	// Endpoint Kategori
	app.Get(prefix+"/kategori", controller.GetKategori)
	app.Post(prefix+"/kategori", controller.CreateKategori)
	app.Delete(prefix+"/kategori/:id", controller.DeleteKategori)
	app.Get(prefix+"/pengumuman/kategori/:nama", controller.GetPengumumanByKategori)
}

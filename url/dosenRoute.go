package url

import (
	"backend/controller"

	"github.com/gofiber/fiber/v2"
)

func DosenRoute(app *fiber.App) {
	app.Get("/dosen", controller.GetAllDosen)
	app.Get("/dosen/:nidn", controller.GetDosenByNIDN)
	app.Post("/dosen", controller.CreateDosen)
	app.Put("/dosen/:nidn", controller.UpdateDosen)
	app.Delete("/dosen/:nidn", controller.DeleteDosen)

	app.Get("/jabatan", controller.GetAllJabatan)
	app.Get("/jabatan/:id", controller.GetJabatanByID)
	app.Post("/jabatan", controller.CreateJabatan)
	app.Put("/jabatan/:id", controller.UpdateJabatan)
	app.Delete("/jabatan/:id", controller.DeleteJabatan)

	app.Get("/api/dosen", controller.GetAllDosen)
	app.Get("/api/dosen/:nidn", controller.GetDosenByNIDN)
	app.Post("/api/dosen", controller.CreateDosen)
	app.Put("/api/dosen/:nidn", controller.UpdateDosen)
	app.Delete("/api/dosen/:nidn", controller.DeleteDosen)

	app.Get("/api/jabatan", controller.GetAllJabatan)
	app.Get("/api/jabatan/:id", controller.GetJabatanByID)
	app.Post("/api/jabatan", controller.CreateJabatan)
	app.Put("/api/jabatan/:id", controller.UpdateJabatan)
	app.Delete("/api/jabatan/:id", controller.DeleteJabatan)
}

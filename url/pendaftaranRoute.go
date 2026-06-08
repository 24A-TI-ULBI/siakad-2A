package url

import (
	"backend/controller"

	"github.com/gofiber/fiber/v2"
)

func PendaftaranRoute(app *fiber.App) {
	app.Post("/beasiswa/daftar", controller.DaftarBeasiswa)
	app.Get("/beasiswa/pendaftar/:id", controller.GetPendaftarBeasiswa)
	app.Get("/beasiswa/status/:npm", controller.CekStatus)
	app.Put("/beasiswa/status/:npm", controller.UpdateStatus)
	app.Delete("/beasiswa/:npm", controller.DeletePendaftaran)
}

package url

import (
	"github.com/gofiber/fiber/v2"
	"backend/controller"
)

func JadwalRoute(app *fiber.App) {

	app.Get("/api/jadwal", controller.GetJadwal)

	app.Post("/api/jadwal", controller.AddJadwal)

	app.Delete("/api/jadwal/:kode", controller.DeleteJadwal)
}
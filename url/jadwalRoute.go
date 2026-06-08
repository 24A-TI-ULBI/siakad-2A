package url

import (
	"backend/controller"

	"github.com/gofiber/fiber/v2"
)

func JadwalRoute(app *fiber.App) {
	jadwal := app.Group("/jadwal")
	jadwal.Get("", controller.GetAllJadwal)
	jadwal.Post("", controller.CreateJadwal)
	jadwal.Get("/:id", controller.GetJadwalByID)
	jadwal.Put("/:id", controller.UpdateJadwal)
	jadwal.Delete("/:id", controller.DeleteJadwal)

	ruangan := app.Group("/ruangan")
	ruangan.Get("", controller.GetAllRuangan)
	ruangan.Post("", controller.CreateRuangan)
	ruangan.Get("/:kode", controller.GetRuanganByKode)
	ruangan.Put("/:kode", controller.UpdateRuangan)
	ruangan.Delete("/:kode", controller.DeleteRuangan)
}

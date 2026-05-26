package url

import (
	"backend/controller"

	"github.com/gofiber/fiber/v2"
)

func NotifikasiRoute(app *fiber.App) {
	registerNotifikasiRoutes(app, "/notifikasi")
	registerNotifikasiRoutes(app, "/api/notifikasi")
}

func registerNotifikasiRoutes(app *fiber.App, prefix string) {
	app.Get(prefix, controller.GetAllNotifikasi)
	app.Post(prefix, controller.CreateNotifikasi)

	app.Get(prefix+"/riwayat/:npm", controller.GetRiwayatNotifikasi)
	app.Get(prefix+"/belum-baca/:npm", controller.GetNotifikasiBelumBaca)
	app.Delete(prefix+"/riwayat/:npm", controller.DeleteRiwayatNotifikasi)

	app.Put(prefix+"/:id/baca", controller.MarkNotifikasiDibaca)
	app.Delete(prefix+"/:id", controller.DeleteNotifikasi)
	app.Get(prefix+"/:npm", controller.GetNotifikasiByNPM)
}

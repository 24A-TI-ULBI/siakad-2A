package url

import (
	"backend/controller"

	"github.com/gofiber/fiber/v2"
)

func OrmawaRoute(app *fiber.App) {
	app.Get("/ormawa", controller.GetAllOrmawa)
	app.Get("/ormawa/:id", controller.GetOrmawaByID)
	app.Post("/ormawa", controller.CreateOrmawa)
	app.Put("/ormawa/:id", controller.UpdateOrmawa)
	app.Delete("/ormawa/:id", controller.DeleteOrmawa)

	app.Get("/kegiatan", controller.GetAllKegiatan)
	app.Get("/kegiatan/ormawa/:id", controller.GetKegiatanByOrmawa)
	app.Get("/kegiatan/:id", controller.GetKegiatanByID)
	app.Post("/kegiatan", controller.CreateKegiatan)
	app.Put("/kegiatan/:id", controller.UpdateKegiatan)
}

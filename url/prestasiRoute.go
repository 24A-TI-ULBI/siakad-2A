package url

import (
	"backend/controller"

	"github.com/gofiber/fiber/v2"
)

func PrestasiRoute(app *fiber.App) {
	app.Get("/prestasi", controller.GetPrestasi)
	app.Get("/prestasi/:npm", controller.GetPrestasiByNPM)
	app.Post("/prestasi", controller.PostPrestasi)
	app.Put("/prestasi/:id", controller.PutPrestasi)
	app.Delete("/prestasi/:id", controller.DeletePrestasi)

	app.Get("/kategori-prestasi", controller.GetKategoriPrestasi)
	app.Post("/kategori-prestasi", controller.PostKategoriPrestasi)
	app.Get("/prestasi/kategori/:nama", controller.GetPrestasiByKategori)
}

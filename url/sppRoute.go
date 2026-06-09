package url

import (
	"backend/controller"

	"github.com/gofiber/fiber/v2"
)

func SppRoute(app *fiber.App) {
	// CRUD Pembayaran SPP
	app.Get("/spp", controller.GetAllSPP)
	app.Get("/spp/:npm", func(c *fiber.Ctx) error {
		if c.Params("npm") == "index.html" {
			return c.Next()
		}
		return controller.GetSPPByNPM(c)
	})
	app.Post("/spp", controller.CreateSPP)
	app.Put("/spp/:id", controller.UpdateSPP)
	app.Delete("/spp/:id", controller.DeleteSPP)

	// Riwayat Pembayaran
	app.Get("/spp/riwayat/:npm", controller.GetRiwayatByNPM)
	app.Get("/spp/lunas/:semester", controller.GetLunasBySemester)
	app.Get("/spp/belum-lunas/:semester", controller.GetBelumLunasBySemester)
}

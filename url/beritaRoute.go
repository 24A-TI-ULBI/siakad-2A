package url

import (
	"backend/controller"

	"github.com/gofiber/fiber/v2"
)

// BeritaRoute mendaftarkan endpoint untuk portal berita kampus.
func BeritaRoute(app *fiber.App) {
	berita := app.Group("/api/berita")
	berita.Get("/", controller.GetAllBerita)
	berita.Get("/:id", controller.GetBeritaByID)
	berita.Post("/", controller.CreateBerita)
	berita.Put("/:id", controller.UpdateBerita)
	berita.Delete("/:id", controller.DeleteBerita)
	berita.Get("/:id/komentar", controller.GetKomentarByBeritaID)
	berita.Post("/:id/komentar", controller.CreateKomentar)

	app.Delete("/api/komentar/:id", controller.DeleteKomentar)
}

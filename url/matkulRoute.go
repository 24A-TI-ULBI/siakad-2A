package url

import (
	"backend/controller"

	"github.com/gofiber/fiber/v2"
)

func MatkulRoute(app *fiber.App) {

	app.Get("/api/matkul", controller.GetMatkul)

	app.Post("/api/matkul", controller.AddMatkul)

	app.Delete("/api/matkul/:kode", controller.DeleteMatkul)
}

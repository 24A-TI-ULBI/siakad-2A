package url

import (
	"backend/controller"

	"github.com/gofiber/fiber/v2"
)

func KuesionerRoute(app *fiber.App) {
	app.Get("/kuesioner", controller.GetAllKuesioner)
	app.Get("/kuesioner/all", controller.GetAllKuesioner) // Added to bypass static interception of GET /kuesioner
	app.Get("/kuesioner/:id", controller.GetKuesionerByID)
	app.Post("/kuesioner", controller.CreateKuesioner)
	app.Put("/kuesioner/:id", controller.UpdateKuesioner)
	app.Delete("/kuesioner/:id", controller.DeleteKuesioner)

	app.Post("/kuesioner/jawab", controller.SubmitJawaban)
	app.Get("/kuesioner/jawaban/:id", controller.GetJawabanByKuesionerID)
	app.Get("/kuesioner/status/:npm", controller.GetStatusByNPM)
}

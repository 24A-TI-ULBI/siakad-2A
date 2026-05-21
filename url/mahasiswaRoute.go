package url

import (
	"backend/controller"

	"github.com/gofiber/fiber/v2"
)

func MahasiswaRoute(app *fiber.App) {
	// Mahasiswa
	app.Get("/mahasiswa", controller.GetAllMahasiswa)
	app.Get("/mahasiswa/:npm", controller.GetMahasiswaByNPM)
	app.Post("/mahasiswa", controller.CreateMahasiswa)
	app.Put("/mahasiswa/:npm", controller.UpdateMahasiswa)
	app.Delete("/mahasiswa/:npm", controller.DeleteMahasiswa)

	// Auth
	app.Post("/auth/login", controller.Login)
	app.Get("/auth/profile/:phone", controller.GetProfile)
}

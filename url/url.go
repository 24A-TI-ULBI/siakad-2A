package url

import (
	"backend/controller"

	"github.com/gofiber/fiber/v2"
)

// Web mendaftarkan semua route ke aplikasi Fiber.
// Setiap modul punya file route sendiri di folder url/.
// Untuk menambah modul baru, buat file [modul]Route.go lalu panggil fungsinya di sini.
func Web(app *fiber.App) {
	// Global
	app.Get("/", controller.Homepage)
	app.Get("/ip", controller.IPServer)

	// Route modul
	// Contoh: MahasiswaRoute(app)
	OrmawaRoute(app)
}

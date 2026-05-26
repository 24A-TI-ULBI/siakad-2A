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
	// GET / dihandle oleh static filesystem (frontend/index.html)
	app.Get("/ip", controller.IPServer)
	app.Get("/api", controller.Homepage) // info API tetap tersedia di /api

	// Route modul
	// Contoh: MahasiswaRoute(app)
	OrmawaRoute(app)
	PrestasiRoute(app)
	// Modul 1 - Mahasiswa & Auth (raditya)
	MahasiswaRoute(app)
}

package url

import (
	"backend/controller"
	"backend/controller/modul9"

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

	// Tambahkan route modul di bawah ini setelah PR di-merge
	PKLRoute(app)
	// Modul 8 - Beasiswa & Pendaftaran (yasmin)
	BeasiswaRoute(app)
	PendaftaranRoute(app)
	// Route modul
	// Contoh: MahasiswaRoute(app)
	SkripsiRoute(app)
	modul9.RegisterRoutes(app)
  DosenRoute(app)

	PengumumanRoute(app)
	NotifikasiRoute(app)
	AbsensiRoute(app)
	OrmawaRoute(app)
	PrestasiRoute(app)
	// Modul 1 - Mahasiswa & Auth (raditya)
	MahasiswaRoute(app)
}

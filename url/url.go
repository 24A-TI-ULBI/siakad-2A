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
	// Pindahkan homepage API ke `/api` supaya root `/` dapat melayani frontend statis
	app.Get("/api", controller.Homepage)
	app.Get("/ip", controller.IPServer)

	// Tambahkan route modul di bawah ini setelah PR di-merge
	PKLRoute(app)
	// Contoh: MahasiswaRoute(app)
}

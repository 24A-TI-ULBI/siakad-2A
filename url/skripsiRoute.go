package url

import (
	"backend/controller"

	"github.com/gofiber/fiber/v2"
)

// SkripsiRoute mendaftarkan semua route untuk modul Skripsi & Bimbingan
func SkripsiRoute(app *fiber.App) {
	// ===== Route Skripsi =====
	app.Get("/skripsi", controller.GetAllSkripsi)          // ambil semua data skripsi
	app.Get("/skripsi/:npm", func(c *fiber.Ctx) error {
		if c.Params("npm") == "index.html" {
			return c.Next()
		}
		return controller.GetSkripsiByNPM(c)
	})   // ambil data skripsi mahasiswa
	app.Post("/skripsi", controller.CreateSkripsi)          // daftarkan judul skripsi
	app.Put("/skripsi/:id", controller.UpdateSkripsi)       // update data skripsi
	app.Delete("/skripsi/:id", controller.DeleteSkripsi)    // hapus data skripsi

	// ===== Route Bimbingan =====
	app.Post("/bimbingan", controller.CreateBimbingan)              // catat sesi bimbingan baru
	app.Get("/bimbingan/:npm", func(c *fiber.Ctx) error {
		if c.Params("npm") == "index.html" {
			return c.Next()
		}
		return controller.GetBimbinganByNPM(c)
	})        // riwayat bimbingan mahasiswa
	app.Put("/bimbingan/:id", controller.UpdateBimbingan)            // update catatan bimbingan
	app.Get("/bimbingan/dosen/:nidn", controller.GetBimbinganByDosen) // daftar bimbingan per dosen
}

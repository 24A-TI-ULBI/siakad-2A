package modul9

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app *fiber.App) {

	app.Get("/modul9", func(c *fiber.Ctx) error {
		return c.Redirect("/Modul-9/perpustakaan.html")
	})

	app.Get("/modul9/admin", func(c *fiber.Ctx) error {
		return c.Redirect("/Modul-9/admin.html")
	})

	// Backend Modul 9 — Buku
	app.Get("/buku", GetAllBuku)
	app.Get("/buku/cari", CariBukuByJudul)
	app.Get("/buku/:id", GetBukuByID)
	app.Post("/buku", CreateBuku)
	app.Put("/buku/:id", UpdateBuku)
	app.Delete("/buku/:id", DeleteBuku)

	// Backend Modul 9 — Peminjaman
	app.Post("/peminjaman", PinjamBuku)
	app.Put("/peminjaman/:id/kembali", KembalikanBuku)
	app.Get("/peminjaman/aktif", GetPeminjamanAktif)
	app.Get("/peminjaman/:npm", GetRiwayatPeminjamanByNPM)
}
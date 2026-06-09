package url

import (
	"backend/controller"
	"github.com/gofiber/fiber/v2"
)

// AbsensiRoute mendaftarkan rute khusus Modul 6 ke router Fiber utama
func AbsensiRoute(r *fiber.App) {
	r.Get("/absensi/hari-ini", controller.GetAbsensiHariIni)
	r.Get("/absensi/all", controller.GetAllAbsensi)
	r.Get("/absensi/:npm", controller.GetAbsensiByNPM)
	r.Post("/absensi", controller.InsertAbsensi)
	r.Put("/absensi/:id", controller.UpdateAbsensi)
	r.Delete("/absensi/:id", controller.DeleteAbsensi)

	// Rekap absensi
	r.Get("/rekap-absensi/matkul/:kode", controller.GetRekapAbsensiByMatkul)
	r.Get("/rekap-absensi/:npm", controller.GetRekapAbsensiByNPM)
}
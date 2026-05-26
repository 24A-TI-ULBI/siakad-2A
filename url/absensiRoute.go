package url

import (
	"github.com/24a-ti-ulbi/siakad-2a/controller"
	"github.com/gofiber/fiber/v2"
)

// AbsensiRoute mendaftarkan rute khusus Modul 6 ke router Fiber utama
func AbsensiRoute(r *fiber.App) {
	r.Get("/absensi/hari-ini", controller.GetAbsensiHariIni)
	r.Get("/absensi/:npm", controller.GetAbsensiByNPM)
	r.Post("/absensi", controller.PostAbsensi)
	r.Put("/absensi/:id", controller.UpdateAbsensiStatus)
	r.Delete("/absensi/:id", controller.DeleteAttendance)
}
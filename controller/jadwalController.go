package controller

import (
	"github.com/gofiber/fiber/v2"
	"backend/model"
)

var JadwalData = []model.Jadwal{
	{
		Kode:   "IF101",
		Nama:   "Pemrograman Web",
		Dosen:    "Pak Budi",
		Hari:     "Senin",
		Jam:      "08:00 - 10:00",
		Ruangan:  "Lab Komputer 1",
		SKS:      3,
		Semester: 2,
	},
	{
		Kode:   "IF102",
		Nama:   "Basis Data",
		Dosen:    "Bu Sinta",
		Hari:     "Selasa",
		Jam:      "10:00 - 12:00",
		Ruangan:  "Lab Komputer 2",
		SKS:      3,
		Semester: 3,
	},
}

// GET ALL JADWAL
func GetJadwal(c *fiber.Ctx) error {
	return c.JSON(JadwalData)
}

// POST TAMBAH JADWAL
func AddJadwal(c *fiber.Ctx) error {

	var data model.Jadwal

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	JadwalData = append(JadwalData, data)

	return c.JSON(fiber.Map{
		"message": "Jadwal berhasil ditambahkan",
	})
}

// DELETE JADWAL
func DeleteJadwal(c *fiber.Ctx) error {

	kode := c.Params("kode")

	for i, jadwal := range JadwalData {

		if jadwal.Kode == kode {

			JadwalData = append(JadwalData[:i], JadwalData[i+1:]...)

			return c.JSON(fiber.Map{
				"message": "Jadwal berhasil dihapus",
			})
		}
	}

	return c.JSON(fiber.Map{
		"message": "Data tidak ditemukan",
	})
}
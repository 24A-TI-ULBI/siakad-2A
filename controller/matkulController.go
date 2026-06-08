package controller

import (
	"backend/model"

	"github.com/gofiber/fiber/v2"
)

var MatkulData = []model.Matkul{
	{
		Kode:     "IF101",
		Nama:     "Pemrograman Web",
		Dosen:    "Pak Budi",
		Hari:     "Senin",
		Jam:      "08:00 - 10:00",
		Ruangan:  "Lab Komputer 1",
		SKS:      3,
		Semester: 2,
	},
	{
		Kode:     "IF102",
		Nama:     "Basis Data",
		Dosen:    "Bu Sinta",
		Hari:     "Selasa",
		Jam:      "10:00 - 12:00",
		Ruangan:  "Lab Komputer 2",
		SKS:      3,
		Semester: 3,
	},
}

// GET ALL MATKUL
func GetMatkul(c *fiber.Ctx) error {
	return c.JSON(MatkulData)
}

// POST TAMBAH MATKUL
func AddMatkul(c *fiber.Ctx) error {

	var data model.Matkul

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	MatkulData = append(MatkulData, data)

	return c.JSON(fiber.Map{
		"message": "Matkul berhasil ditambahkan",
	})
}

// DELETE MATKUL
func DeleteMatkul(c *fiber.Ctx) error {

	kode := c.Params("kode")

	for i, matkul := range MatkulData {

		if matkul.Kode == kode {

			MatkulData = append(MatkulData[:i], MatkulData[i+1:]...)

			return c.JSON(fiber.Map{
				"message": "Matkul berhasil dihapus",
			})
		}
	}

	return c.JSON(fiber.Map{
		"message": "Data tidak ditemukan",
	})
}

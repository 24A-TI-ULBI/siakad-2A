package controller

import (
	"backend/helper"
	"backend/model"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

// GET ALL MATKUL
func GetMatkul(c *fiber.Ctx) error {
	db := helper.GetDB()
	matkuls, err := helper.GetAllDoc[model.Matkul](db, "matkul", bson.M{})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "gagal mengambil data matkul", "detail": err.Error()})
	}
	if matkuls == nil {
		matkuls = make([]model.Matkul, 0)
	}
	return c.JSON(matkuls)
}

// POST TAMBAH MATKUL
func AddMatkul(c *fiber.Ctx) error {
	var data model.Matkul
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	db := helper.GetDB()
	_, err := helper.InsertOneDoc(db, "matkul", data)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "gagal menyimpan matkul", "detail": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Matkul berhasil ditambahkan"})
}

// DELETE MATKUL
func DeleteMatkul(c *fiber.Ctx) error {
	kode := c.Params("kode")
	db := helper.GetDB()
	_, err := helper.DeleteDoc(db, "matkul", bson.M{"kode": kode})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "gagal menghapus matkul", "detail": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Matkul berhasil dihapus"})
}

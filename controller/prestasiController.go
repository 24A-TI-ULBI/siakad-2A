package controller

import (
	"backend/helper"
	"backend/model"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetPrestasi mengambil semua data prestasi
func GetPrestasi(c *fiber.Ctx) error {
	db := helper.GetDB()
	docs, err := helper.GetAllDoc[model.Prestasi](db, "prestasi", bson.M{})
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal mengambil data prestasi: "+err.Error())
	}
	return helper.SuccessResponse(c, docs)
}

// GetPrestasiByNPM mengambil prestasi mahasiswa tertentu berdasarkan NPM
func GetPrestasiByNPM(c *fiber.Ctx) error {
	npm := c.Params("npm")
	db := helper.GetDB()
	docs, err := helper.GetAllDoc[model.Prestasi](db, "prestasi", bson.M{"npm": npm})
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal mengambil data prestasi: "+err.Error())
	}
	return helper.SuccessResponse(c, docs)
}

// PostPrestasi input prestasi baru
func PostPrestasi(c *fiber.Ctx) error {
	var prestasi model.Prestasi
	if err := c.BodyParser(&prestasi); err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Request body tidak valid: "+err.Error())
	}

	if prestasi.NPM == "" || prestasi.NamaEvent == "" || prestasi.Tingkat == "" || prestasi.Juara == "" || prestasi.Tanggal == "" || prestasi.Kategori == "" {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Semua field harus diisi")
	}

	prestasi.ID = primitive.NewObjectID()
	db := helper.GetDB()
	_, err := helper.InsertOneDoc(db, "prestasi", prestasi)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal menyimpan data prestasi: "+err.Error())
	}
	return helper.SuccessResponse(c, prestasi)
}

// PutPrestasi update data prestasi
func PutPrestasi(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Format ID tidak valid: "+err.Error())
	}

	var prestasi model.Prestasi
	if err := c.BodyParser(&prestasi); err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Request body tidak valid: "+err.Error())
	}

	if prestasi.NPM == "" || prestasi.NamaEvent == "" || prestasi.Tingkat == "" || prestasi.Juara == "" || prestasi.Tanggal == "" || prestasi.Kategori == "" {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Semua field harus diisi")
	}

	db := helper.GetDB()
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"npm":        prestasi.NPM,
			"nama_event": prestasi.NamaEvent,
			"tingkat":    prestasi.Tingkat,
			"juara":      prestasi.Juara,
			"tanggal":    prestasi.Tanggal,
			"kategori":   prestasi.Kategori,
		},
	}

	_, err = helper.UpdateDoc(db, "prestasi", filter, update)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal memperbarui data prestasi: "+err.Error())
	}

	prestasi.ID = id
	return helper.SuccessResponse(c, prestasi)
}

// DeletePrestasi hapus data prestasi
func DeletePrestasi(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Format ID tidak valid: "+err.Error())
	}

	db := helper.GetDB()
	_, err = helper.DeleteDoc(db, "prestasi", bson.M{"_id": id})
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal menghapus data prestasi: "+err.Error())
	}

	return helper.SuccessResponse(c, fiber.Map{
		"message": "Data prestasi berhasil dihapus",
	})
}

// GetKategoriPrestasi ambil semua kategori
func GetKategoriPrestasi(c *fiber.Ctx) error {
	db := helper.GetDB()
	docs, err := helper.GetAllDoc[model.KategoriPrestasi](db, "kategoriprestasi", bson.M{})
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal mengambil data kategori: "+err.Error())
	}
	return helper.SuccessResponse(c, docs)
}

// PostKategoriPrestasi tambah kategori baru
func PostKategoriPrestasi(c *fiber.Ctx) error {
	var kategori model.KategoriPrestasi
	if err := c.BodyParser(&kategori); err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Request body tidak valid: "+err.Error())
	}

	if kategori.Nama == "" {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Nama kategori harus diisi")
	}

	kategori.ID = primitive.NewObjectID()
	db := helper.GetDB()
	_, err := helper.InsertOneDoc(db, "kategoriprestasi", kategori)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal menyimpan kategori: "+err.Error())
	}
	return helper.SuccessResponse(c, kategori)
}

// GetPrestasiByKategori filter prestasi berdasarkan kategori
func GetPrestasiByKategori(c *fiber.Ctx) error {
	namaKategori := c.Params("nama")
	db := helper.GetDB()
	docs, err := helper.GetAllDoc[model.Prestasi](db, "prestasi", bson.M{"kategori": namaKategori})
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal mengambil data prestasi: "+err.Error())
	}
	return helper.SuccessResponse(c, docs)
}

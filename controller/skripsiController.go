package controller

import (
	"backend/helper"
	"backend/model"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ===================== SKRIPSI CONTROLLER =====================

// GetAllSkripsi mengambil semua data skripsi
func GetAllSkripsi(c *fiber.Ctx) error {
	db := helper.GetDB()
	skripsi, err := helper.GetAllDoc[model.Skripsi](db, "skripsi", bson.M{})
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal mengambil data skripsi: "+err.Error())
	}
	return helper.SuccessResponse(c, skripsi)
}

// GetSkripsiByNPM mengambil data skripsi berdasarkan NPM mahasiswa
func GetSkripsiByNPM(c *fiber.Ctx) error {
	npm := c.Params("npm")
	db := helper.GetDB()
	skripsi, err := helper.GetOneDoc[model.Skripsi](db, "skripsi", bson.M{"npm": npm})
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusNotFound, "Data skripsi tidak ditemukan untuk NPM: "+npm)
	}
	return helper.SuccessResponse(c, skripsi)
}

// CreateSkripsi menambahkan data skripsi baru (daftarkan judul)
func CreateSkripsi(c *fiber.Ctx) error {
	var skripsi model.Skripsi
	if err := c.BodyParser(&skripsi); err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Format data tidak valid: "+err.Error())
	}

	// Validasi field wajib
	if skripsi.NPM == "" || skripsi.Judul == "" || skripsi.DosenPembimbing1 == "" {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "NPM, judul, dan dosen pembimbing 1 wajib diisi")
	}

	// Set status default jika kosong
	if skripsi.Status == "" {
		skripsi.Status = "pengajuan"
	}

	db := helper.GetDB()
	insertedID, err := helper.InsertOneDoc(db, "skripsi", skripsi)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal menyimpan data skripsi: "+err.Error())
	}
	return helper.SuccessResponse(c, fiber.Map{
		"message":     "Data skripsi berhasil ditambahkan",
		"inserted_id": insertedID,
	})
}

// UpdateSkripsi mengupdate data skripsi berdasarkan ID
func UpdateSkripsi(c *fiber.Ctx) error {
	idParam := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Format ID tidak valid")
	}

	var skripsi model.Skripsi
	if err := c.BodyParser(&skripsi); err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Format data tidak valid: "+err.Error())
	}

	update := bson.M{}
	if skripsi.Judul != "" {
		update["judul"] = skripsi.Judul
	}
	if skripsi.DosenPembimbing1 != "" {
		update["dosen_pembimbing_1"] = skripsi.DosenPembimbing1
	}
	if skripsi.NIDNPembimbing1 != "" {
		update["nidn_pembimbing_1"] = skripsi.NIDNPembimbing1
	}
	if skripsi.DosenPembimbing2 != "" {
		update["dosen_pembimbing_2"] = skripsi.DosenPembimbing2
	}
	if skripsi.NIDNPembimbing2 != "" {
		update["nidn_pembimbing_2"] = skripsi.NIDNPembimbing2
	}
	if skripsi.Status != "" {
		update["status"] = skripsi.Status
	}
	if skripsi.TanggalSelesai != "" {
		update["tanggal_selesai"] = skripsi.TanggalSelesai
	}

	if len(update) == 0 {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Tidak ada field yang diupdate")
	}

	db := helper.GetDB()
	result, err := helper.UpdateDoc(db, "skripsi", bson.M{"_id": objID}, bson.M{"$set": update})
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal mengupdate data skripsi: "+err.Error())
	}
	if result.MatchedCount == 0 {
		return helper.ErrorResponse(c, fiber.StatusNotFound, "Data skripsi tidak ditemukan")
	}
	return helper.SuccessResponse(c, fiber.Map{
		"message": "Data skripsi berhasil diupdate",
	})
}

// DeleteSkripsi menghapus data skripsi berdasarkan ID
func DeleteSkripsi(c *fiber.Ctx) error {
	idParam := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Format ID tidak valid")
	}

	db := helper.GetDB()
	result, err := helper.DeleteDoc(db, "skripsi", bson.M{"_id": objID})
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal menghapus data skripsi: "+err.Error())
	}
	if result.DeletedCount == 0 {
		return helper.ErrorResponse(c, fiber.StatusNotFound, "Data skripsi tidak ditemukan")
	}
	return helper.SuccessResponse(c, fiber.Map{
		"message": "Data skripsi berhasil dihapus",
	})
}

// ===================== BIMBINGAN CONTROLLER =====================

// CreateBimbingan mencatat sesi bimbingan baru
func CreateBimbingan(c *fiber.Ctx) error {
	var bimbingan model.Bimbingan
	if err := c.BodyParser(&bimbingan); err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Format data tidak valid: "+err.Error())
	}

	// Validasi field wajib
	if bimbingan.NPM == "" || bimbingan.NIDN == "" || bimbingan.Tanggal == "" || bimbingan.Catatan == "" {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "NPM, NIDN, tanggal, dan catatan wajib diisi")
	}

	// Set status default jika kosong
	if bimbingan.Status == "" {
		bimbingan.Status = "menunggu"
	}

	db := helper.GetDB()
	insertedID, err := helper.InsertOneDoc(db, "bimbingan", bimbingan)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal menyimpan data bimbingan: "+err.Error())
	}
	return helper.SuccessResponse(c, fiber.Map{
		"message":     "Sesi bimbingan berhasil dicatat",
		"inserted_id": insertedID,
	})
}

// GetBimbinganByNPM mengambil riwayat bimbingan mahasiswa berdasarkan NPM
func GetBimbinganByNPM(c *fiber.Ctx) error {
	npm := c.Params("npm")
	db := helper.GetDB()
	bimbingan, err := helper.GetAllDoc[model.Bimbingan](db, "bimbingan", bson.M{"npm": npm})
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal mengambil data bimbingan: "+err.Error())
	}
	return helper.SuccessResponse(c, bimbingan)
}

// UpdateBimbingan mengupdate catatan bimbingan berdasarkan ID
func UpdateBimbingan(c *fiber.Ctx) error {
	idParam := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Format ID tidak valid")
	}

	var bimbingan model.Bimbingan
	if err := c.BodyParser(&bimbingan); err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Format data tidak valid: "+err.Error())
	}

	update := bson.M{}
	if bimbingan.Catatan != "" {
		update["catatan"] = bimbingan.Catatan
	}
	if bimbingan.Status != "" {
		update["status"] = bimbingan.Status
	}
	if bimbingan.ProgressBab != "" {
		update["progress_bab"] = bimbingan.ProgressBab
	}

	if len(update) == 0 {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Tidak ada field yang diupdate")
	}

	db := helper.GetDB()
	result, err := helper.UpdateDoc(db, "bimbingan", bson.M{"_id": objID}, bson.M{"$set": update})
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal mengupdate data bimbingan: "+err.Error())
	}
	if result.MatchedCount == 0 {
		return helper.ErrorResponse(c, fiber.StatusNotFound, "Data bimbingan tidak ditemukan")
	}
	return helper.SuccessResponse(c, fiber.Map{
		"message": "Data bimbingan berhasil diupdate",
	})
}

// GetBimbinganByDosen mengambil daftar bimbingan berdasarkan NIDN dosen
func GetBimbinganByDosen(c *fiber.Ctx) error {
	nidn := c.Params("nidn")
	db := helper.GetDB()
	bimbingan, err := helper.GetAllDoc[model.Bimbingan](db, "bimbingan", bson.M{"nidn": nidn})
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal mengambil data bimbingan: "+err.Error())
	}
	return helper.SuccessResponse(c, bimbingan)
}

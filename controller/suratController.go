package controller

import (
	"time"

	"backend/helper"
	"backend/model"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ==========================================
// 1. SURAT KETERANGAN (TEMPLATE) HANDLERS
// ==========================================

// GetTemplates retrieves all letter templates
func GetTemplates(c *fiber.Ctx) error {
	db := helper.GetDB()
	templates, err := helper.GetAllDoc[model.SuratKeterangan](db, "surat_template", bson.M{})
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal mengambil data template: "+err.Error())
	}
	return helper.SuccessResponse(c, templates)
}

// GetTemplateByID retrieves a letter template by ID
func GetTemplateByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Format ID tidak valid")
	}

	db := helper.GetDB()
	template, err := helper.GetOneDoc[model.SuratKeterangan](db, "surat_template", bson.M{"_id": id})
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusNotFound, "Template tidak ditemukan")
	}
	return helper.SuccessResponse(c, template)
}

// CreateTemplate creates a new letter template
func CreateTemplate(c *fiber.Ctx) error {
	var template model.SuratKeterangan
	if err := c.BodyParser(&template); err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Request body tidak valid")
	}

	if template.Kode == "" || template.Nama == "" {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Kode dan Nama template wajib diisi")
	}

	template.ID = primitive.NewObjectID()
	template.CreatedAt = time.Now()
	template.UpdatedAt = time.Now()

	db := helper.GetDB()
	_, err := helper.InsertOneDoc(db, "surat_template", template)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal menambahkan template: "+err.Error())
	}

	return helper.SuccessResponse(c, template)
}

// UpdateTemplate updates an existing letter template
func UpdateTemplate(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Format ID tidak valid")
	}

	var templateUpdate model.SuratKeterangan
	if err := c.BodyParser(&templateUpdate); err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Request body tidak valid")
	}

	if templateUpdate.Kode == "" || templateUpdate.Nama == "" {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Kode dan Nama template wajib diisi")
	}

	db := helper.GetDB()
	update := bson.M{
		"$set": bson.M{
			"kode":        templateUpdate.Kode,
			"nama":        templateUpdate.Nama,
			"deskripsi":   templateUpdate.Deskripsi,
			"persyaratan": templateUpdate.Persyaratan,
			"updated_at":  time.Now(),
		},
	}

	_, err = helper.UpdateDoc(db, "surat_template", bson.M{"_id": id}, update)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal memperbarui template: "+err.Error())
	}

	templateUpdate.ID = id
	return helper.SuccessResponse(c, templateUpdate)
}

// DeleteTemplate deletes a letter template by ID
func DeleteTemplate(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Format ID tidak valid")
	}

	db := helper.GetDB()
	_, err = helper.DeleteDoc(db, "surat_template", bson.M{"_id": id})
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal menghapus template: "+err.Error())
	}

	return helper.SuccessResponse(c, fiber.Map{
		"message": "Template berhasil dihapus",
	})
}

// ==========================================
// 2. PENGAJUAN SURAT HANDLERS
// ==========================================

// SubmitPengajuan creates a request for a letter
func SubmitPengajuan(c *fiber.Ctx) error {
	var pengajuan model.PengajuanSurat
	if err := c.BodyParser(&pengajuan); err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Request body tidak valid")
	}

	if pengajuan.SuratID == "" || pengajuan.NPM == "" || pengajuan.NamaMahasiswa == "" || pengajuan.Prodi == "" {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "SuratID, NPM, Nama, dan Prodi wajib diisi")
	}

	pengajuan.ID = primitive.NewObjectID()
	pengajuan.Status = "proses" // Default status
	pengajuan.TanggalPengajuan = time.Now()
	pengajuan.TanggalUpdate = time.Now()

	db := helper.GetDB()
	_, err := helper.InsertOneDoc(db, "pengajuan_surat", pengajuan)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal mengajukan surat: "+err.Error())
	}

	return helper.SuccessResponse(c, pengajuan)
}

// GetPengajuanByNPM retrieves requests made by a specific student using NPM
func GetPengajuanByNPM(c *fiber.Ctx) error {
	npm := c.Params("npm")
	if npm == "" {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "NPM wajib disertakan")
	}

	db := helper.GetDB()
	pengajuans, err := helper.GetAllDoc[model.PengajuanSurat](db, "pengajuan_surat", bson.M{"npm": npm})
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal mengambil data pengajuan: "+err.Error())
	}

	return helper.SuccessResponse(c, pengajuans)
}

// UpdateStatusPengajuan updates the status of a letter request (proses/selesai/ditolak)
func UpdateStatusPengajuan(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Format ID tidak valid")
	}

	var input struct {
		Status  string `json:"status"`  // "proses", "selesai", "ditolak"
		Catatan string `json:"catatan"` // Admin note/reason
	}

	if err := c.BodyParser(&input); err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Request body tidak valid")
	}

	if input.Status != "proses" && input.Status != "selesai" && input.Status != "ditolak" {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Status tidak valid. Harus 'proses', 'selesai', atau 'ditolak'")
	}

	db := helper.GetDB()
	update := bson.M{
		"$set": bson.M{
			"status":         input.Status,
			"catatan":        input.Catatan,
			"tanggal_update": time.Now(),
		},
	}

	_, err = helper.UpdateDoc(db, "pengajuan_surat", bson.M{"_id": id}, update)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal memperbarui status pengajuan: "+err.Error())
	}

	return helper.SuccessResponse(c, fiber.Map{
		"message": "Status pengajuan berhasil diperbarui",
		"id":      idStr,
		"status":  input.Status,
		"catatan": input.Catatan,
	})
}

// GetPengajuanByStatus filters all submissions by their status
func GetPengajuanByStatus(c *fiber.Ctx) error {
	status := c.Params("status")
	if status != "proses" && status != "selesai" && status != "ditolak" {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Status tidak valid. Harus 'proses', 'selesai', atau 'ditolak'")
	}

	db := helper.GetDB()
	pengajuans, err := helper.GetAllDoc[model.PengajuanSurat](db, "pengajuan_surat", bson.M{"status": status})
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal mengambil data pengajuan: "+err.Error())
	}

	return helper.SuccessResponse(c, pengajuans)
}

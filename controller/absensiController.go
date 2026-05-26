package controller

import (
	"backend/helper"
	"backend/model"
	"math"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const AbsensiCollection = "absensi"

// GetAbsensiByNPM menangani GET /absensi/:npm
func GetAbsensiByNPM(c *fiber.Ctx) error {
	npm := c.Params("npm")
	if npm == "" {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "NPM tidak boleh kosong")
	}

	db := helper.GetDB()
	filter := bson.M{"npm": npm}

	docs, err := helper.GetAllDoc[model.Absensi](db, AbsensiCollection, filter)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal mengambil data absensi: "+err.Error())
	}

	// Mengembalikan array objek absensi langsung sesuai ekspektasi frontend
	return c.JSON(docs)
}

// GetAbsensiHariIni menangani GET /absensi/hari-ini
func GetAbsensiHariIni(c *fiber.Ctx) error {
	loc := time.FixedZone("WIB", 7*3600)
	todayStr := time.Now().In(loc).Format("2006-01-02")

	db := helper.GetDB()
	filter := bson.M{"tanggal": todayStr}

	docs, err := helper.GetAllDoc[model.Absensi](db, AbsensiCollection, filter)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal mengambil data absensi hari ini: "+err.Error())
	}

	// Mengembalikan array objek absensi langsung
	return c.JSON(docs)
}

// InsertAbsensi menangani POST /absensi
func InsertAbsensi(c *fiber.Ctx) error {
	var req model.Absensi
	if err := c.BodyParser(&req); err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}

	// Waktu Indonesia Barat (WIB) UTC+7
	loc := time.FixedZone("WIB", 7*3600)
	now := time.Now().In(loc)

	// Auto-generasi field opsional jika kosong
	if req.ID == "" {
		req.ID = primitive.NewObjectID().Hex()
	}
	if req.Tanggal == "" {
		req.Tanggal = now.Format("2006-01-02")
	}
	if req.Timestamp == "" {
		req.Timestamp = now.Format("15:04")
	}

	// Validasi input wajib
	if req.NPM == "" || req.MatkulCode == "" || req.Status == "" {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Data absensi tidak lengkap")
	}

	db := helper.GetDB()
	_, err := helper.InsertOneDoc(db, AbsensiCollection, req)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal menyimpan data absensi: "+err.Error())
	}

	return helper.SuccessResponse(c, req)
}

// UpdateAbsensi menangani PUT /absensi/:id
func UpdateAbsensi(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "ID tidak boleh kosong")
	}

	var req struct {
		Status string `json:"status"`
	}
	if err := c.BodyParser(&req); err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}

	if req.Status == "" {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Status tidak boleh kosong")
	}

	db := helper.GetDB()
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": req.Status}}

	_, err := helper.UpdateDoc(db, AbsensiCollection, filter, update)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal memperbarui status absensi: "+err.Error())
	}

	return helper.SuccessResponse(c, fiber.Map{"message": "Status absensi berhasil diperbarui"})
}

// DeleteAbsensi menangani DELETE /absensi/:id
func DeleteAbsensi(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "ID tidak boleh kosong")
	}

	db := helper.GetDB()
	filter := bson.M{"_id": id}

	_, err := helper.DeleteDoc(db, AbsensiCollection, filter)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal menghapus data absensi: "+err.Error())
	}

	return helper.SuccessResponse(c, fiber.Map{"message": "Data absensi berhasil dihapus"})
}

// GetRekapAbsensiByNPM menangani GET /rekap-absensi/:npm
func GetRekapAbsensiByNPM(c *fiber.Ctx) error {
	npm := c.Params("npm")
	if npm == "" {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "NPM tidak boleh kosong")
	}

	db := helper.GetDB()
	filter := bson.M{"npm": npm}

	docs, err := helper.GetAllDoc[model.Absensi](db, AbsensiCollection, filter)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal mengambil data rekap: "+err.Error())
	}

	// Map untuk mengelompokkan absensi berdasarkan matkulCode
	rekapMap := make(map[string]*model.RekapMatkul)

	var totalHadir, totalIzin, totalAlpha, totalSesi int

	for _, doc := range docs {
		if doc.MatkulCode == "" {
			continue
		}

		if _, exists := rekapMap[doc.MatkulCode]; !exists {
			rekapMap[doc.MatkulCode] = &model.RekapMatkul{
				MatkulCode: doc.MatkulCode,
			}
		}

		mRekap := rekapMap[doc.MatkulCode]
		mRekap.TotalSesi++
		totalSesi++

		switch doc.Status {
		case "Hadir":
			mRekap.Hadir++
			totalHadir++
		case "Izin":
			mRekap.Izin++
			totalIzin++
		case "Alpha":
			mRekap.Alpha++
			totalAlpha++
		}
	}

	// Konversi map ke slice dan hitung persentase per matkul
	rekapList := make([]model.RekapMatkul, 0, len(rekapMap))
	for _, mRekap := range rekapMap {
		if mRekap.TotalSesi > 0 {
			mRekap.Persentase = math.Round((float64(mRekap.Hadir)/float64(mRekap.TotalSesi))*100*100) / 100
		}
		rekapList = append(rekapList, *mRekap)
	}

	overallPersentase := 0.0
	if totalSesi > 0 {
		overallPersentase = math.Round((float64(totalHadir)/float64(totalSesi))*100*100) / 100
	}

	res := model.StudentRekap{
		NPM:                 npm,
		TotalSesi:           totalSesi,
		Hadir:               totalHadir,
		Izin:                totalIzin,
		Alpha:               totalAlpha,
		PersentaseKehadiran: overallPersentase,
		RekapPerMatkul:      rekapList,
	}

	return helper.SuccessResponse(c, res)
}

// GetRekapAbsensiByMatkul menangani GET /rekap-absensi/matkul/:kode
func GetRekapAbsensiByMatkul(c *fiber.Ctx) error {
	kode := c.Params("kode")
	if kode == "" {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Kode matkul tidak boleh kosong")
	}

	db := helper.GetDB()
	filter := bson.M{"matkulCode": kode}

	docs, err := helper.GetAllDoc[model.Absensi](db, AbsensiCollection, filter)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal mengambil data rekap matkul: "+err.Error())
	}

	// Map untuk mengelompokkan absensi berdasarkan NPM mahasiswa
	studentMap := make(map[string]*model.StudentMatkulRekap)

	for _, doc := range docs {
		if doc.NPM == "" {
			continue
		}

		if _, exists := studentMap[doc.NPM]; !exists {
			studentMap[doc.NPM] = &model.StudentMatkulRekap{
				NPM: doc.NPM,
			}
		}

		sRekap := studentMap[doc.NPM]
		sRekap.TotalSesi++

		switch doc.Status {
		case "Hadir":
			sRekap.Hadir++
		case "Izin":
			sRekap.Izin++
		case "Alpha":
			sRekap.Alpha++
		}
	}

	// Konversi map ke slice dan hitung persentase per mahasiswa
	pendaftarList := make([]model.StudentMatkulRekap, 0, len(studentMap))
	for _, sRekap := range studentMap {
		if sRekap.TotalSesi > 0 {
			sRekap.Persentase = math.Round((float64(sRekap.Hadir)/float64(sRekap.TotalSesi))*100*100) / 100
		}
		pendaftarList = append(pendaftarList, *sRekap)
	}

	res := model.CourseRekap{
		MatkulCode: kode,
		Pendaftar:  pendaftarList,
	}

	return helper.SuccessResponse(c, res)
}

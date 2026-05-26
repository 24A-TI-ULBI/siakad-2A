package controller

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
	"time"

	"backend/helper"
	"backend/model"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ==========================================
// 1. PENGUMUMAN CONTROLLERS
// ==========================================

// GetPengumuman handles GET /pengumuman
// Ambil semua pengumuman (terbaru di atas)
func GetPengumuman(c *fiber.Ctx) error {
	db := helper.GetDB()
	list, err := helper.GetAllDoc[model.Pengumuman](db, "pengumuman", bson.M{})
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal mengambil data pengumuman: "+err.Error())
	}

	// Urutkan berdasarkan tanggal terbaru di atas
	sort.Slice(list, func(i, j int) bool {
		return list[i].Tanggal.After(list[j].Tanggal)
	})

	return helper.SuccessResponse(c, list)
}

// GetPengumumanByID handles GET /pengumuman/:id
// Detail pengumuman berdasarkan ID
func GetPengumumanByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Format ID tidak valid")
	}

	db := helper.GetDB()
	doc, err := helper.GetOneDoc[model.Pengumuman](db, "pengumuman", bson.M{"_id": objID})
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusNotFound, "Pengumuman tidak ditemukan")
	}

	return helper.SuccessResponse(c, doc)
}

// CreatePengumuman handles POST /pengumuman
// Tambah pengumuman baru
func CreatePengumuman(c *fiber.Ctx) error {
	var req model.Pengumuman
	if err := c.BodyParser(&req); err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Format body tidak valid: "+err.Error())
	}

	// Validasi input
	if req.Judul == "" || req.Isi == "" || req.Kategori == "" {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Judul, Isi, dan Kategori tidak boleh kosong")
	}

	req.ID = primitive.NewObjectID()
	req.Tanggal = time.Now()
	if req.Penulis == "" {
		req.Penulis = "Admin Akademik"
	}

	db := helper.GetDB()
	_, err := helper.InsertOneDoc(db, "pengumuman", req)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal menambahkan pengumuman: "+err.Error())
	}

	return helper.SuccessResponse(c, req)
}

// UpdatePengumuman handles PUT /pengumuman/:id
// Update data pengumuman
func UpdatePengumuman(c *fiber.Ctx) error {
	idStr := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Format ID tidak valid")
	}

	var req model.Pengumuman
	if err := c.BodyParser(&req); err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Format body tidak valid: "+err.Error())
	}

	// Validasi input
	if req.Judul == "" || req.Isi == "" || req.Kategori == "" {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Judul, Isi, dan Kategori tidak boleh kosong")
	}

	if req.Penulis == "" {
		req.Penulis = "Admin Akademik"
	}

	db := helper.GetDB()
	update := bson.M{
		"$set": bson.M{
			"judul":     req.Judul,
			"isi":       req.Isi,
			"kategori":  req.Kategori,
			"penulis":   req.Penulis,
			"file_url":  req.FileUrl,
			"file_type": req.FileType,
			"tanggal":   time.Now(), // Update timestamp saat diedit
		},
	}

	_, err = helper.UpdateDoc(db, "pengumuman", bson.M{"_id": objID}, update)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal memperbarui pengumuman: "+err.Error())
	}

	// Ambil data terbaru hasil update untuk di-return
	updatedDoc, err := helper.GetOneDoc[model.Pengumuman](db, "pengumuman", bson.M{"_id": objID})
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal mengambil data terupdate: "+err.Error())
	}

	return helper.SuccessResponse(c, updatedDoc)
}

// DeletePengumuman handles DELETE /pengumuman/:id
// Hapus pengumuman
func DeletePengumuman(c *fiber.Ctx) error {
	idStr := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Format ID tidak valid")
	}

	db := helper.GetDB()
	res, err := helper.DeleteDoc(db, "pengumuman", bson.M{"_id": objID})
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal menghapus pengumuman: "+err.Error())
	}

	if res.DeletedCount == 0 {
		return helper.ErrorResponse(c, fiber.StatusNotFound, "Pengumuman tidak ditemukan")
	}

	return helper.SuccessResponse(c, fiber.Map{
		"message": "Pengumuman berhasil dihapus",
	})
}

// ==========================================
// 2. KATEGORI CONTROLLERS
// ==========================================

// GetKategori handles GET /kategori
// Ambil semua kategori
func GetKategori(c *fiber.Ctx) error {
	db := helper.GetDB()
	list, err := helper.GetAllDoc[model.Kategori](db, "kategori", bson.M{})
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal mengambil data kategori: "+err.Error())
	}

	return helper.SuccessResponse(c, list)
}

// CreateKategori handles POST /kategori
// Tambah kategori baru
func CreateKategori(c *fiber.Ctx) error {
	var req model.Kategori
	if err := c.BodyParser(&req); err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Format body tidak valid: "+err.Error())
	}

	// Validasi input
	if req.Nama == "" {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Nama kategori tidak boleh kosong")
	}

	req.ID = primitive.NewObjectID()

	db := helper.GetDB()
	_, err := helper.InsertOneDoc(db, "kategori", req)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal menambahkan kategori: "+err.Error())
	}

	return helper.SuccessResponse(c, req)
}

// GetPengumumanByKategori handles GET /pengumuman/kategori/:nama
// Filter pengumuman berdasarkan nama kategori (case-insensitive)
func GetPengumumanByKategori(c *fiber.Ctx) error {
	katNama := c.Params("nama")

	db := helper.GetDB()

	// Gunakan regex case-insensitive untuk pencarian nama kategori agar lebih fleksibel
	filter := bson.M{
		"kategori": bson.M{
			"$regex": primitive.Regex{
				Pattern: "^" + regexp.QuoteMeta(katNama) + "$",
				Options: "i",
			},
		},
	}

	list, err := helper.GetAllDoc[model.Pengumuman](db, "pengumuman", filter)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal mengambil data pengumuman filter: "+err.Error())
	}

	// Urutkan berdasarkan tanggal terbaru di atas
	sort.Slice(list, func(i, j int) bool {
		return list[i].Tanggal.After(list[j].Tanggal)
	})

	return helper.SuccessResponse(c, list)
}

// DeleteKategori handles DELETE /kategori/:id
// Hapus kategori berdasarkan ID
func DeleteKategori(c *fiber.Ctx) error {
	idStr := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Format ID tidak valid")
	}

	db := helper.GetDB()
	res, err := helper.DeleteDoc(db, "kategori", bson.M{"_id": objID})
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal menghapus kategori: "+err.Error())
	}

	if res.DeletedCount == 0 {
		return helper.ErrorResponse(c, fiber.StatusNotFound, "Kategori tidak ditemukan")
	}

	return helper.SuccessResponse(c, fiber.Map{
		"message": "Kategori berhasil dihapus",
	})
}

// UploadFile handles file upload for announcements
func UploadFile(c *fiber.Ctx) error {
	// 1. Get file from form
	file, err := c.FormFile("file")
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "File tidak ditemukan dalam request: "+err.Error())
	}

	// 2. Security validation: Size limit (max size: 5MB) to prevent DOS/abuse
	const MaxFileSize = 5 * 1024 * 1024 // 5 Megabytes
	if file.Size > MaxFileSize {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Ukuran file terlalu besar. Maksimum: 5MB")
	}

	// 3. Security validation: Content-Type checks to prevent malicious uploads (XSS/shell execution)
	contentType := file.Header.Get("Content-Type")
	validTypes := map[string]string{
		"image/jpeg":      "image",
		"image/png":       "image",
		"image/gif":       "image",
		"image/webp":      "image",
		"application/pdf": "pdf",
		"video/mp4":       "video",
		"video/webm":      "video",
		"video/ogg":       "video",
	}

	fileType, isValid := validTypes[contentType]
	if !isValid {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Tipe file tidak diizinkan. Hanya menerima Gambar (JPG, PNG, GIF, WEBP), PDF, dan Video (MP4, WEBM).")
	}

	// 4. Generate unique filename to avoid path traversal / collision
	ext := ""
	if strings.Contains(file.Filename, ".") {
		parts := strings.Split(file.Filename, ".")
		ext = "." + parts[len(parts)-1]
	}
	
	// Sanitasi ekstensi agar tidak ada karakter aneh
	ext = regexp.MustCompile(`[^a-zA-Z0-9.]`).ReplaceAllString(ext, "")
	
	// Hindari upload script executable (.exe, .go, .php, .js, .html, etc.)
	badExts := map[string]bool{
		".exe":  true,
		".go":   true,
		".php":  true,
		".html": true,
		".htm":  true,
		".js":   true,
		".jsp":  true,
		".sh":   true,
		".bat":  true,
		".cmd":  true,
	}
	if badExts[strings.ToLower(ext)] {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Ekstensi file tidak diizinkan.")
	}

	uniqueName := fmt.Sprintf("%d_%s%s", time.Now().UnixNano(), primitive.NewObjectID().Hex(), ext)

	// 5. Ensure upload directory exists
	uploadDir := "./frontend/pengumuman/uploads"
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal membuat direktori upload: "+err.Error())
	}

	// 6. Save file to path
	filePath := fmt.Sprintf("%s/%s", uploadDir, uniqueName)
	if err := c.SaveFile(file, filePath); err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal menyimpan file: "+err.Error())
	}

	// Return URL that is publicly accessible via frontend static files
	publicUrl := fmt.Sprintf("/pengumuman/uploads/%s", uniqueName)

	return helper.SuccessResponse(c, fiber.Map{
		"file_url":  publicUrl,
		"file_type": fileType,
		"filename":  file.Filename,
	})
}



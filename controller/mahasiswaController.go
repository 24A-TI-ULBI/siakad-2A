package controller

import (
	"backend/helper"
	"backend/model"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var phoneRegex = regexp.MustCompile(`^\d{10,15}$`)

func validateMahasiswa(mhs model.Mahasiswa) string {
	if mhs.NPM == "" {
		return "NPM wajib diisi"
	}
	if mhs.Nama == "" {
		return "Nama wajib diisi"
	}
	if mhs.Phone == "" {
		return "Nomor HP wajib diisi"
	}
	cleaned := strings.ReplaceAll(mhs.Phone, "-", "")
	cleaned = strings.ReplaceAll(cleaned, " ", "")
	if !phoneRegex.MatchString(cleaned) {
		return "Nomor HP harus berupa angka, minimal 10 digit dan maksimal 15 digit"
	}
	if mhs.Email != "" {
		if !strings.Contains(mhs.Email, "@") || !strings.Contains(mhs.Email, ".") {
			return "Format email tidak valid"
		}
		if utf8.RuneCountInString(mhs.Email) > 100 {
			return "Email terlalu panjang"
		}
	}
	if mhs.Angkatan != 0 && (mhs.Angkatan < 2000 || mhs.Angkatan > 2099) {
		return "Angkatan harus antara 2000 dan 2099"
	}
	return ""
}

func GetAllMahasiswa(c *fiber.Ctx) error {
	col := helper.GetCollection("mahasiswa")
	ctx, cancel := helper.GetContext()
	defer cancel()

	filter := bson.M{}

	if nama := strings.TrimSpace(c.Query("nama")); nama != "" {
		filter["nama"] = bson.M{"$regex": nama, "$options": "i"}
	}
	if prodi := strings.TrimSpace(c.Query("prodi")); prodi != "" {
		filter["prodi"] = bson.M{"$regex": prodi, "$options": "i"}
	}
	if angkatanStr := c.Query("angkatan"); angkatanStr != "" {
		if angkatan, err := strconv.Atoi(angkatanStr); err == nil {
			filter["angkatan"] = angkatan
		}
	}

	cursor, err := col.Find(ctx, filter)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal mengambil data mahasiswa")
	}
	defer cursor.Close(ctx)

	var list []model.Mahasiswa
	if err := cursor.All(ctx, &list); err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal membaca data mahasiswa")
	}
	if list == nil {
		list = []model.Mahasiswa{}
	}
	return helper.SuccessResponse(c, list)
}

func GetMahasiswaByNPM(c *fiber.Ctx) error {
	npm := c.Params("npm")
	col := helper.GetCollection("mahasiswa")
	ctx, cancel := helper.GetContext()
	defer cancel()

	var mhs model.Mahasiswa
	if err := col.FindOne(ctx, bson.M{"npm": npm}).Decode(&mhs); err != nil {
		return helper.ErrorResponse(c, fiber.StatusNotFound, "Mahasiswa tidak ditemukan")
	}
	return helper.SuccessResponse(c, mhs)
}

func CreateMahasiswa(c *fiber.Ctx) error {
	var mhs model.Mahasiswa
	if err := c.BodyParser(&mhs); err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Format request tidak valid")
	}

	if errMsg := validateMahasiswa(mhs); errMsg != "" {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, errMsg)
	}

	col := helper.GetCollection("mahasiswa")
	ctx, cancel := helper.GetContext()
	defer cancel()

	var existing model.Mahasiswa
	if err := col.FindOne(ctx, bson.M{"npm": mhs.NPM}).Decode(&existing); err == nil {
		return helper.ErrorResponse(c, fiber.StatusConflict, "NPM sudah terdaftar")
	}

	mhs.ID = primitive.NewObjectID()
	result, err := col.InsertOne(ctx, mhs)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal menyimpan data mahasiswa")
	}
	return helper.SuccessResponse(c, fiber.Map{"inserted_id": result.InsertedID})
}

func UpdateMahasiswa(c *fiber.Ctx) error {
	npm := c.Params("npm")

	var body bson.M
	if err := c.BodyParser(&body); err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Format request tidak valid")
	}
	delete(body, "npm")
	delete(body, "_id")

	// Validasi email jika ada di body
	if email, ok := body["email"].(string); ok && email != "" {
		if !strings.Contains(email, "@") || !strings.Contains(email, ".") {
			return helper.ErrorResponse(c, fiber.StatusBadRequest, "Format email tidak valid")
		}
	}
	// Validasi phone jika ada di body
	if phone, ok := body["phone"].(string); ok && phone != "" {
		cleaned := strings.ReplaceAll(phone, "-", "")
		cleaned = strings.ReplaceAll(cleaned, " ", "")
		if !phoneRegex.MatchString(cleaned) {
			return helper.ErrorResponse(c, fiber.StatusBadRequest, "Nomor HP harus berupa angka, minimal 10 digit")
		}
	}

	col := helper.GetCollection("mahasiswa")
	ctx, cancel := helper.GetContext()
	defer cancel()

	result, err := col.UpdateOne(ctx, bson.M{"npm": npm}, bson.M{"$set": body})
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal mengupdate data mahasiswa")
	}
	if result.MatchedCount == 0 {
		return helper.ErrorResponse(c, fiber.StatusNotFound, "Mahasiswa tidak ditemukan")
	}
	return helper.SuccessResponse(c, fiber.Map{"updated": result.ModifiedCount})
}

func DeleteMahasiswa(c *fiber.Ctx) error {
	npm := c.Params("npm")
	col := helper.GetCollection("mahasiswa")
	ctx, cancel := helper.GetContext()
	defer cancel()

	result, err := col.DeleteOne(ctx, bson.M{"npm": npm})
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal menghapus data mahasiswa")
	}
	if result.DeletedCount == 0 {
		return helper.ErrorResponse(c, fiber.StatusNotFound, "Mahasiswa tidak ditemukan")
	}
	return helper.SuccessResponse(c, fiber.Map{"deleted": result.DeletedCount})
}

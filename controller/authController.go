package controller

import (
	"backend/config"
	"backend/helper"
	"backend/model"
	"strings"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func Login(c *fiber.Ctx) error {
	var body struct {
		Phone string `json:"phone"`
	}
	if err := c.BodyParser(&body); err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Format request tidak valid")
	}
	if body.Phone == "" {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Nomor telepon wajib diisi")
	}

	col := helper.GetCollection("mahasiswa")
	ctx, cancel := helper.GetContext()
	defer cancel()

	var mhs model.Mahasiswa
	if err := col.FindOne(ctx, bson.M{"phone": body.Phone}).Decode(&mhs); err != nil {
		return helper.ErrorResponse(c, fiber.StatusUnauthorized, "Nomor telepon tidak terdaftar")
	}

	token, err := config.GenerateToken(mhs.NPM, mhs.Phone, mhs.Nama)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal membuat token")
	}

	return helper.SuccessResponse(c, fiber.Map{
		"token": token,
		"user": fiber.Map{
			"npm":   mhs.NPM,
			"nama":  mhs.Nama,
			"phone": mhs.Phone,
			"prodi": mhs.Prodi,
		},
	})
}

func GetProfile(c *fiber.Ctx) error {
	phone := c.Params("phone")
	col := helper.GetCollection("mahasiswa")
	ctx, cancel := helper.GetContext()
	defer cancel()

	var mhs model.Mahasiswa
	if err := col.FindOne(ctx, bson.M{"phone": phone}).Decode(&mhs); err != nil {
		return helper.ErrorResponse(c, fiber.StatusNotFound, "Profil tidak ditemukan")
	}
	return helper.SuccessResponse(c, mhs)
}

// GetMe mengambil profil mahasiswa dari JWT token di header Authorization.
func GetMe(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return helper.ErrorResponse(c, fiber.StatusUnauthorized, "Token tidak ditemukan")
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return helper.ErrorResponse(c, fiber.StatusUnauthorized, "Format token tidak valid")
	}

	claims, err := config.ValidateToken(parts[1])
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusUnauthorized, "Token tidak valid atau sudah expired")
	}

	npm, ok := claims["npm"].(string)
	if !ok || npm == "" {
		return helper.ErrorResponse(c, fiber.StatusUnauthorized, "Token tidak valid")
	}

	col := helper.GetCollection("mahasiswa")
	ctx, cancel := helper.GetContext()
	defer cancel()

	var mhs model.Mahasiswa
	if err := col.FindOne(ctx, bson.M{"npm": npm}).Decode(&mhs); err != nil {
		return helper.ErrorResponse(c, fiber.StatusNotFound, "Profil tidak ditemukan")
	}
	return helper.SuccessResponse(c, mhs)
}

package controller

import (
	"backend/helper"
	"backend/model"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DaftarBeasiswa — POST /beasiswa/daftar
func DaftarBeasiswa(c *fiber.Ctx) error {
	var data model.Pendaftaran
	if err := c.BodyParser(&data); err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Data tidak valid")
	}
	if data.NPM == "" || data.Beasiswa == "" {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "NPM dan nama beasiswa wajib diisi")
	}

	data.ID = primitive.NewObjectID()
	data.Status = "Pending"

	col := helper.GetCollection("pendaftaran")
	ctx, cancel := helper.GetContext()
	defer cancel()

	result, err := col.InsertOne(ctx, data)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal menyimpan pendaftaran")
	}
	return helper.SuccessResponse(c, fiber.Map{"inserted_id": result.InsertedID})
}

// GetPendaftarBeasiswa — GET /beasiswa/pendaftar/:id
func GetPendaftarBeasiswa(c *fiber.Ctx) error {
	beasiswaID := c.Params("id")

	col := helper.GetCollection("pendaftaran")
	ctx, cancel := helper.GetContext()
	defer cancel()

	cursor, err := col.Find(ctx, bson.M{"beasiswa_id": beasiswaID})
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal mengambil data pendaftar")
	}
	defer cursor.Close(ctx)

	var list []model.Pendaftaran
	if err := cursor.All(ctx, &list); err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal membaca data pendaftar")
	}
	if list == nil {
		list = []model.Pendaftaran{}
	}
	return helper.SuccessResponse(c, list)
}

// CekStatus — GET /beasiswa/status/:npm
func CekStatus(c *fiber.Ctx) error {
	npm := c.Params("npm")

	col := helper.GetCollection("pendaftaran")
	ctx, cancel := helper.GetContext()
	defer cancel()

	var result model.Pendaftaran
	if err := col.FindOne(ctx, bson.M{"npm": npm}).Decode(&result); err != nil {
		return helper.ErrorResponse(c, fiber.StatusNotFound, "Data pendaftaran tidak ditemukan")
	}
	return helper.SuccessResponse(c, result)
}

// UpdateStatus — PUT /beasiswa/status/:npm
func UpdateStatus(c *fiber.Ctx) error {
	npm := c.Params("npm")

	var body struct {
		Status string `json:"status"`
	}
	if err := c.BodyParser(&body); err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Format request tidak valid")
	}
	if body.Status == "" {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Status wajib diisi")
	}

	col := helper.GetCollection("pendaftaran")
	ctx, cancel := helper.GetContext()
	defer cancel()

	result, err := col.UpdateOne(ctx, bson.M{"npm": npm}, bson.M{"$set": bson.M{"status": body.Status}})
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal mengupdate status")
	}
	if result.MatchedCount == 0 {
		return helper.ErrorResponse(c, fiber.StatusNotFound, "Data pendaftaran tidak ditemukan")
	}
	return helper.SuccessResponse(c, fiber.Map{"updated": result.ModifiedCount})
}

// DeletePendaftaran — DELETE /beasiswa/pendaftaran/:id
func DeletePendaftaran(c *fiber.Ctx) error {

	npm := c.Params("npm")

	col := helper.GetCollection("pendaftaran")

	ctx, cancel := helper.GetContext()
	defer cancel()

	result, err := col.DeleteOne(ctx, bson.M{
		"npm": npm,
	})

	if err != nil {
		return helper.ErrorResponse(
			c,
			fiber.StatusInternalServerError,
			"Gagal menghapus data pendaftaran",
		)
	}

	if result.DeletedCount == 0 {
		return helper.ErrorResponse(
			c,
			fiber.StatusNotFound,
			"Data pendaftaran tidak ditemukan",
		)
	}

	return helper.SuccessResponse(c, fiber.Map{
		"deleted": result.DeletedCount,
	})
}
package controller

import (
	"backend/helper"
	"backend/model"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetAllSPP - GET /spp
func GetAllSPP(c *fiber.Ctx) error {
	col := helper.GetCollection("pembayaran_spp")
	ctx, cancel := helper.GetContext()
	defer cancel()

	cursor, err := col.Find(ctx, bson.D{})
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal mengambil data SPP")
	}
	defer cursor.Close(ctx)

	var list []model.PembayaranSPP
	if err := cursor.All(ctx, &list); err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal membaca data SPP")
	}
	if list == nil {
		list = []model.PembayaranSPP{}
	}
	return helper.SuccessResponse(c, list)
}

// GetSPPByNPM - GET /spp/:npm
func GetSPPByNPM(c *fiber.Ctx) error {
	npm := c.Params("npm")
	col := helper.GetCollection("pembayaran_spp")
	ctx, cancel := helper.GetContext()
	defer cancel()

	var spp []model.PembayaranSPP
	cursor, err := col.Find(ctx, bson.M{"npm": npm})
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusNotFound, "Tagihan SPP tidak ditemukan")
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &spp); err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal membaca data")
	}
	if spp == nil {
		spp = []model.PembayaranSPP{}
	}
	return helper.SuccessResponse(c, spp)
}

// CreateSPP - POST /spp
func CreateSPP(c *fiber.Ctx) error {
	var spp model.PembayaranSPP
	if err := c.BodyParser(&spp); err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Format request tidak valid")
	}
	if spp.NPM == "" || spp.Semester == "" || spp.JumlahTagihan == 0 {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "NPM, semester, dan jumlah tagihan wajib diisi")
	}

	col := helper.GetCollection("pembayaran_spp")
	ctx, cancel := helper.GetContext()
	defer cancel()

	spp.ID = primitive.NewObjectID()
	if spp.StatusPembayaran == "" {
		spp.StatusPembayaran = "belum_lunas"
	}
	result, err := col.InsertOne(ctx, spp)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal menyimpan tagihan SPP")
	}
	return helper.SuccessResponse(c, fiber.Map{"inserted_id": result.InsertedID})
}

// UpdateSPP - PUT /spp/:id
func UpdateSPP(c *fiber.Ctx) error {
	id := c.Params("id")
	var body bson.M
	if err := c.BodyParser(&body); err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Format request tidak valid")
	}

	// Hapus field yang tidak boleh diupdate langsung
	delete(body, "_id")
	delete(body, "npm")
	delete(body, "semester")

	col := helper.GetCollection("pembayaran_spp")
	ctx, cancel := helper.GetContext()
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "ID tidak valid")
	}

	result, err := col.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": body})
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal mengupdate status pembayaran")
	}
	if result.MatchedCount == 0 {
		return helper.ErrorResponse(c, fiber.StatusNotFound, "Tagihan tidak ditemukan")
	}
	return helper.SuccessResponse(c, fiber.Map{"updated": result.ModifiedCount})
}

// DeleteSPP - DELETE /spp/:id
func DeleteSPP(c *fiber.Ctx) error {
	id := c.Params("id")
	col := helper.GetCollection("pembayaran_spp")
	ctx, cancel := helper.GetContext()
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "ID tidak valid")
	}

	result, err := col.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal menghapus tagihan")
	}
	if result.DeletedCount == 0 {
		return helper.ErrorResponse(c, fiber.StatusNotFound, "Tagihan tidak ditemukan")
	}
	return helper.SuccessResponse(c, fiber.Map{"deleted": result.DeletedCount})
}

// GetRiwayatByNPM - GET /spp/riwayat/:npm
func GetRiwayatByNPM(c *fiber.Ctx) error {
	npm := c.Params("npm")
	col := helper.GetCollection("pembayaran_spp")
	ctx, cancel := helper.GetContext()
	defer cancel()

	var riwayat []model.PembayaranSPP
	cursor, err := col.Find(ctx, bson.M{
		"npm":               npm,
		"status_pembayaran": "lunas",
	})
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal mengambil riwayat")
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &riwayat); err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal membaca riwayat")
	}
	if riwayat == nil {
		riwayat = []model.PembayaranSPP{}
	}
	return helper.SuccessResponse(c, riwayat)
}

// GetLunasBySemester - GET /spp/lunas/:semester
func GetLunasBySemester(c *fiber.Ctx) error {
	semester := c.Params("semester")
	col := helper.GetCollection("pembayaran_spp")
	ctx, cancel := helper.GetContext()
	defer cancel()

	var list []model.PembayaranSPP
	cursor, err := col.Find(ctx, bson.M{
		"semester":          semester,
		"status_pembayaran": "lunas",
	})
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal mengambil data")
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &list); err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal membaca data")
	}
	if list == nil {
		list = []model.PembayaranSPP{}
	}
	return helper.SuccessResponse(c, list)
}

// GetBelumLunasBySemester - GET /spp/belum-lunas/:semester
func GetBelumLunasBySemester(c *fiber.Ctx) error {
	semester := c.Params("semester")
	col := helper.GetCollection("pembayaran_spp")
	ctx, cancel := helper.GetContext()
	defer cancel()

	var list []model.PembayaranSPP
	cursor, err := col.Find(ctx, bson.M{
		"semester":          semester,
		"status_pembayaran": "belum_lunas",
	})
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal mengambil data")
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &list); err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal membaca data")
	}
	if list == nil {
		list = []model.PembayaranSPP{}
	}
	return helper.SuccessResponse(c, list)
}

package controller

import (
	"strings"
	"time"

	"backend/helper"
	"backend/model"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const notifikasiCollection = "notifikasi"

func GetAllNotifikasi(c *fiber.Ctx) error {
	return findNotifikasi(c, bson.M{})
}

func GetNotifikasiByNPM(c *fiber.Ctx) error {
	npm := strings.TrimSpace(c.Params("npm"))
	if npm == "" {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "NPM wajib diisi")
	}

	return findNotifikasi(c, bson.M{"npm": npm})
}

func CreateNotifikasi(c *fiber.Ctx) error {
	var payload model.Notifikasi
	if err := c.BodyParser(&payload); err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Body request tidak valid")
	}

	payload.NPM = strings.TrimSpace(payload.NPM)
	payload.Judul = strings.TrimSpace(payload.Judul)
	payload.Pesan = strings.TrimSpace(payload.Pesan)
	payload.Tipe = strings.TrimSpace(payload.Tipe)
	payload.Prioritas = strings.TrimSpace(payload.Prioritas)
	payload.DikirimOleh = strings.TrimSpace(payload.DikirimOleh)

	if payload.NPM == "" || payload.Judul == "" || payload.Pesan == "" {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "NPM, judul, dan pesan wajib diisi")
	}
	if payload.Tipe == "" {
		payload.Tipe = "Umum"
	}
	if payload.Prioritas == "" {
		payload.Prioritas = "Normal"
	}
	if payload.DikirimOleh == "" {
		payload.DikirimOleh = "Admin Akademik"
	}

	now := time.Now()
	payload.ID = primitive.NewObjectID()
	payload.Dibaca = false
	payload.DibuatPada = now
	payload.DibacaPada = nil

	ctx, cancel := helper.GetContext()
	defer cancel()

	_, err := helper.GetCollection(notifikasiCollection).InsertOne(ctx, payload)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal menyimpan notifikasi")
	}

	return helper.SuccessResponse(c, payload)
}

func MarkNotifikasiDibaca(c *fiber.Ctx) error {
	id, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "ID notifikasi tidak valid")
	}

	now := time.Now()
	ctx, cancel := helper.GetContext()
	defer cancel()

	result, err := helper.GetCollection(notifikasiCollection).UpdateOne(ctx, bson.M{"_id": id}, bson.M{
		"$set": bson.M{
			"dibaca":      true,
			"dibaca_pada": now,
		},
	})
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal memperbarui notifikasi")
	}
	if result.MatchedCount == 0 {
		return helper.ErrorResponse(c, fiber.StatusNotFound, "Notifikasi tidak ditemukan")
	}

	return helper.SuccessResponse(c, fiber.Map{
		"id":       id.Hex(),
		"dibaca":   true,
		"modified": result.ModifiedCount,
	})
}

func DeleteNotifikasi(c *fiber.Ctx) error {
	id, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "ID notifikasi tidak valid")
	}

	ctx, cancel := helper.GetContext()
	defer cancel()

	result, err := helper.GetCollection(notifikasiCollection).DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal menghapus notifikasi")
	}
	if result.DeletedCount == 0 {
		return helper.ErrorResponse(c, fiber.StatusNotFound, "Notifikasi tidak ditemukan")
	}

	return helper.SuccessResponse(c, fiber.Map{
		"id":      id.Hex(),
		"deleted": result.DeletedCount,
	})
}

func GetRiwayatNotifikasi(c *fiber.Ctx) error {
	npm := strings.TrimSpace(c.Params("npm"))
	if npm == "" {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "NPM wajib diisi")
	}

	return findNotifikasi(c, bson.M{"npm": npm})
}

func GetNotifikasiBelumBaca(c *fiber.Ctx) error {
	npm := strings.TrimSpace(c.Params("npm"))
	if npm == "" {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "NPM wajib diisi")
	}

	return findNotifikasi(c, bson.M{"npm": npm, "dibaca": false})
}

func DeleteRiwayatNotifikasi(c *fiber.Ctx) error {
	npm := strings.TrimSpace(c.Params("npm"))
	if npm == "" {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "NPM wajib diisi")
	}

	ctx, cancel := helper.GetContext()
	defer cancel()

	result, err := helper.GetCollection(notifikasiCollection).DeleteMany(ctx, bson.M{"npm": npm})
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal menghapus riwayat notifikasi")
	}

	return helper.SuccessResponse(c, fiber.Map{
		"npm":     npm,
		"deleted": result.DeletedCount,
	})
}

func findNotifikasi(c *fiber.Ctx, filter bson.M) error {
	ctx, cancel := helper.GetContext()
	defer cancel()

	cursor, err := helper.GetCollection(notifikasiCollection).Find(
		ctx,
		filter,
		options.Find().SetSort(bson.D{{Key: "dibuat_pada", Value: -1}}),
	)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal mengambil data notifikasi")
	}
	defer cursor.Close(ctx)

	var data []model.Notifikasi
	if err := cursor.All(ctx, &data); err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal membaca data notifikasi")
	}
	if data == nil {
		data = []model.Notifikasi{}
	}

	return helper.SuccessResponse(c, data)
}

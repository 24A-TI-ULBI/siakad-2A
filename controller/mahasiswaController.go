package controller

import (
	"backend/helper"
	"backend/model"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllMahasiswa(c *fiber.Ctx) error {
	col := helper.GetCollection("mahasiswa")
	ctx, cancel := helper.GetContext()
	defer cancel()

	cursor, err := col.Find(ctx, bson.D{})
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
	if mhs.NPM == "" || mhs.Nama == "" || mhs.Phone == "" {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "NPM, nama, dan phone wajib diisi")
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

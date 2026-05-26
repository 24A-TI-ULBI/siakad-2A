package modul9

import (
"backend/helper"
	"backend/model/modul9"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllBuku(c *fiber.Ctx) error {
	ctx, cancel := helper.GetContext()
	defer cancel()

	collection := helper.GetCollection("buku")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}
	defer cursor.Close(ctx)

	var buku []model.Buku
	if err := cursor.All(ctx, &buku); err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return helper.SuccessResponse(c, buku)
}

func GetBukuByID(c *fiber.Ctx) error {
	id := c.Params("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "ID buku tidak valid")
	}

	ctx, cancel := helper.GetContext()
	defer cancel()

	collection := helper.GetCollection("buku")

	var buku model.Buku
	err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&buku)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusNotFound, "Buku tidak ditemukan")
	}

	return helper.SuccessResponse(c, buku)
}

func CariBukuByJudul(c *fiber.Ctx) error {
	judul := c.Query("judul")

	if judul == "" {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Query judul wajib diisi")
	}

	ctx, cancel := helper.GetContext()
	defer cancel()

	collection := helper.GetCollection("buku")

	filter := bson.M{
		"judul": bson.M{
			"$regex":   judul,
			"$options": "i",
		},
	}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}
	defer cursor.Close(ctx)

	var buku []model.Buku
	if err := cursor.All(ctx, &buku); err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return helper.SuccessResponse(c, buku)
}

func CreateBuku(c *fiber.Ctx) error {
	var buku model.Buku

	if err := c.BodyParser(&buku); err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Request body tidak valid")
	}

	if buku.Judul == "" || buku.Penulis == "" {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Judul dan penulis wajib diisi")
	}

	buku.ID = primitive.NewObjectID()
	buku.Tersedia = buku.Stok > 0
	buku.CreatedAt = time.Now()
	buku.UpdatedAt = time.Now()

	ctx, cancel := helper.GetContext()
	defer cancel()

	collection := helper.GetCollection("buku")

	_, err := collection.InsertOne(ctx, buku)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return helper.SuccessResponse(c, buku)
}

func UpdateBuku(c *fiber.Ctx) error {
	id := c.Params("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "ID buku tidak valid")
	}

	var buku model.Buku
	if err := c.BodyParser(&buku); err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Request body tidak valid")
	}

	buku.Tersedia = buku.Stok > 0
	buku.UpdatedAt = time.Now()

	update := bson.M{
		"$set": bson.M{
			"judul":      buku.Judul,
			"penulis":    buku.Penulis,
			"penerbit":   buku.Penerbit,
			"tahun":      buku.Tahun,
			"stok":       buku.Stok,
			"tersedia":   buku.Tersedia,
			"updated_at": buku.UpdatedAt,
		},
	}

	ctx, cancel := helper.GetContext()
	defer cancel()

	collection := helper.GetCollection("buku")

	result, err := collection.UpdateOne(ctx, bson.M{"_id": objectID}, update)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	if result.MatchedCount == 0 {
		return helper.ErrorResponse(c, fiber.StatusNotFound, "Buku tidak ditemukan")
	}

	return helper.SuccessResponse(c, fiber.Map{
		"message": "Data buku berhasil diupdate",
	})
}

func DeleteBuku(c *fiber.Ctx) error {
	id := c.Params("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "ID buku tidak valid")
	}

	ctx, cancel := helper.GetContext()
	defer cancel()

	collection := helper.GetCollection("buku")

	result, err := collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	if result.DeletedCount == 0 {
		return helper.ErrorResponse(c, fiber.StatusNotFound, "Buku tidak ditemukan")
	}

	return helper.SuccessResponse(c, fiber.Map{
		"message": "Buku berhasil dihapus",
	})
}
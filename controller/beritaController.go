package controller

import (
	"backend/helper"
	"backend/model"
	"context"
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetAllBerita returns all news articles ordered terbaru dulu.
func GetAllBerita(c *fiber.Ctx) error {
	db := helper.GetDB()
	ctx, cancel := helper.GetContext()
	defer cancel()

	cursor, err := db.Collection("berita").Find(ctx, bson.M{}, options.Find().SetSort(bson.D{{"tanggal", -1}}))
	if err != nil {
		return helper.ErrorResponse(c, 500, err.Error())
	}

	var berita []model.Berita
	if err := cursor.All(ctx, &berita); err != nil {
		return helper.ErrorResponse(c, 500, err.Error())
	}

	return helper.SuccessResponse(c, berita)
}

// GetBeritaByID returns a single news article by ID.
func GetBeritaByID(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return helper.ErrorResponse(c, 400, "id berita tidak valid")
	}

	db := helper.GetDB()

	var berita model.Berita
	if err := db.Collection("berita").FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&berita); err != nil {
		if err == mongo.ErrNoDocuments {
			return helper.ErrorResponse(c, 404, "berita tidak ditemukan")
		}
		return helper.ErrorResponse(c, 500, err.Error())
	}

	return helper.SuccessResponse(c, berita)
}

// CreateBerita creates a new news article.
func CreateBerita(c *fiber.Ctx) error {
	var input struct {
		Judul   string `json:"judul"`
		Isi     string `json:"isi"`
		Penulis string `json:"penulis"`
	}
	if err := json.Unmarshal(c.Body(), &input); err != nil {
		return helper.ErrorResponse(c, 400, "data tidak valid")
	}

	if input.Judul == "" || input.Isi == "" || input.Penulis == "" {
		return helper.ErrorResponse(c, 400, "judul, isi, dan penulis wajib diisi")
	}

	berita := model.Berita{
		ID:      primitive.NewObjectID(),
		Judul:   input.Judul,
		Isi:     input.Isi,
		Penulis: input.Penulis,
		Tanggal: time.Now(),
	}

	db := helper.GetDB()
	if _, err := helper.InsertOneDoc(db, "berita", berita); err != nil {
		return helper.ErrorResponse(c, 500, err.Error())
	}

	return helper.SuccessResponse(c, berita)
}

// UpdateBerita updates an existing news article.
func UpdateBerita(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return helper.ErrorResponse(c, 400, "id berita tidak valid")
	}

	var input struct {
		Judul   string `json:"judul"`
		Isi     string `json:"isi"`
		Penulis string `json:"penulis"`
	}
	if err := c.BodyParser(&input); err != nil {
		return helper.ErrorResponse(c, 400, "data tidak valid")
	}

	update := bson.M{"$set": bson.M{}}
	fields := update["$set"].(bson.M)
	if input.Judul != "" {
		fields["judul"] = input.Judul
	}
	if input.Isi != "" {
		fields["isi"] = input.Isi
	}
	if input.Penulis != "" {
		fields["penulis"] = input.Penulis
	}

	if len(fields) == 0 {
		return helper.ErrorResponse(c, 400, "tidak ada data yang diupdate")
	}

	db := helper.GetDB()
	result, err := helper.UpdateDoc(db, "berita", bson.M{"_id": objID}, update)
	if err != nil {
		return helper.ErrorResponse(c, 500, err.Error())
	}
	if result.MatchedCount == 0 {
		return helper.ErrorResponse(c, 404, "berita tidak ditemukan")
	}

	return helper.SuccessResponse(c, fiber.Map{"updated": result.ModifiedCount})
}

// DeleteBerita deletes a news article.
func DeleteBerita(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return helper.ErrorResponse(c, 400, "id berita tidak valid")
	}

	db := helper.GetDB()
	result, err := helper.DeleteDoc(db, "berita", bson.M{"_id": objID})
	if err != nil {
		return helper.ErrorResponse(c, 500, err.Error())
	}
	if result.DeletedCount == 0 {
		return helper.ErrorResponse(c, 404, "berita tidak ditemukan")
	}

	return helper.SuccessResponse(c, fiber.Map{"deleted": result.DeletedCount})
}

// GetKomentarByBeritaID returns comments for a specific news article.
func GetKomentarByBeritaID(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return helper.ErrorResponse(c, 400, "id berita tidak valid")
	}

	db := helper.GetDB()
	ctx, cancel := helper.GetContext()
	defer cancel()

	cursor, err := db.Collection("komentar").Find(ctx, bson.M{"berita_id": objID}, options.Find().SetSort(bson.D{{"tanggal", 1}}))
	if err != nil {
		return helper.ErrorResponse(c, 500, err.Error())
	}

	var komentar []model.Komentar
	if err := cursor.All(ctx, &komentar); err != nil {
		return helper.ErrorResponse(c, 500, err.Error())
	}

	return helper.SuccessResponse(c, komentar)
}

// CreateKomentar creates a new comment on a news article.
func CreateKomentar(c *fiber.Ctx) error {
	id := c.Params("id")
	beritaID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return helper.ErrorResponse(c, 400, "id berita tidak valid")
	}

	var input struct {
		Nama  string `json:"nama"`
		Pesan string `json:"pesan"`
	}
	if err := c.BodyParser(&input); err != nil {
		return helper.ErrorResponse(c, 400, "data tidak valid")
	}

	if input.Nama == "" || input.Pesan == "" {
		return helper.ErrorResponse(c, 400, "nama dan pesan wajib diisi")
	}

	komentar := model.Komentar{
		ID:       primitive.NewObjectID(),
		BeritaID: beritaID,
		Nama:     input.Nama,
		Pesan:    input.Pesan,
		Tanggal:  time.Now(),
	}

	db := helper.GetDB()
	if _, err := helper.InsertOneDoc(db, "komentar", komentar); err != nil {
		return helper.ErrorResponse(c, 500, err.Error())
	}

	return helper.SuccessResponse(c, komentar)
}

// DeleteKomentar removes a comment.
func DeleteKomentar(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return helper.ErrorResponse(c, 400, "id komentar tidak valid")
	}

	db := helper.GetDB()
	result, err := helper.DeleteDoc(db, "komentar", bson.M{"_id": objID})
	if err != nil {
		return helper.ErrorResponse(c, 500, err.Error())
	}
	if result.DeletedCount == 0 {
		return helper.ErrorResponse(c, 404, "komentar tidak ditemukan")
	}

	return helper.SuccessResponse(c, fiber.Map{"deleted": result.DeletedCount})
}

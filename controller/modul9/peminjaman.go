package modul9

import (
	"backend/helper"
	"backend/model/modul9"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RequestPeminjaman struct {
	BukuID        string `json:"buku_id"`
	NPM          string `json:"npm"`
	NamaMahasiswa string `json:"nama_mahasiswa"`
}

func PinjamBuku(c *fiber.Ctx) error {
	var request RequestPeminjaman

	if err := c.BodyParser(&request); err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Request body tidak valid")
	}

	if request.BukuID == "" || request.NPM == "" || request.NamaMahasiswa == "" {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "buku_id, npm, dan nama_mahasiswa wajib diisi")
	}

	bukuID, err := primitive.ObjectIDFromHex(request.BukuID)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "ID buku tidak valid")
	}

	ctx, cancel := helper.GetContext()
	defer cancel()

	bukuCollection := helper.GetCollection("buku")
	peminjamanCollection := helper.GetCollection("peminjaman")

	var buku model.Buku
	err = bukuCollection.FindOne(ctx, bson.M{"_id": bukuID}).Decode(&buku)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusNotFound, "Buku tidak ditemukan")
	}

	if buku.Stok <= 0 {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Stok buku habis")
	}

	now := time.Now()

	peminjaman := model.Peminjaman{
		ID:             primitive.NewObjectID(),
		BukuID:         bukuID,
		NPM:            request.NPM,
		NamaMahasiswa:  request.NamaMahasiswa,
		TanggalPinjam:  now,
		Status:         "aktif",
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	_, err = peminjamanCollection.InsertOne(ctx, peminjaman)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	stokBaru := buku.Stok - 1
	tersedia := stokBaru > 0

	_, err = bukuCollection.UpdateOne(
		ctx,
		bson.M{"_id": bukuID},
		bson.M{
			"$set": bson.M{
				"stok":       stokBaru,
				"tersedia":   tersedia,
				"updated_at": now,
			},
		},
	)

	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return helper.SuccessResponse(c, peminjaman)
}

func KembalikanBuku(c *fiber.Ctx) error {
	id := c.Params("id")

	peminjamanID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "ID peminjaman tidak valid")
	}

	ctx, cancel := helper.GetContext()
	defer cancel()

	peminjamanCollection := helper.GetCollection("peminjaman")
	bukuCollection := helper.GetCollection("buku")

	var peminjaman model.Peminjaman
	err = peminjamanCollection.FindOne(ctx, bson.M{"_id": peminjamanID}).Decode(&peminjaman)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusNotFound, "Data peminjaman tidak ditemukan")
	}

	if peminjaman.Status == "kembali" {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Buku sudah dikembalikan")
	}

	now := time.Now()

	_, err = peminjamanCollection.UpdateOne(
		ctx,
		bson.M{"_id": peminjamanID},
		bson.M{
			"$set": bson.M{
				"status":          "kembali",
				"tanggal_kembali": now,
				"updated_at":      now,
			},
		},
	)

	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	_, err = bukuCollection.UpdateOne(
		ctx,
		bson.M{"_id": peminjaman.BukuID},
		bson.M{
			"$inc": bson.M{"stok": 1},
			"$set": bson.M{
				"tersedia":   true,
				"updated_at": now,
			},
		},
	)

	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return helper.SuccessResponse(c, fiber.Map{
		"message": "Buku berhasil dikembalikan",
	})
}

func GetRiwayatPeminjamanByNPM(c *fiber.Ctx) error {
	npm := c.Params("npm")

	ctx, cancel := helper.GetContext()
	defer cancel()

	collection := helper.GetCollection("peminjaman")

	cursor, err := collection.Find(ctx, bson.M{"npm": npm})
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}
	defer cursor.Close(ctx)

	var peminjaman []model.Peminjaman
	if err := cursor.All(ctx, &peminjaman); err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return helper.SuccessResponse(c, peminjaman)
}

func GetPeminjamanAktif(c *fiber.Ctx) error {
	ctx, cancel := helper.GetContext()
	defer cancel()

	collection := helper.GetCollection("peminjaman")

	cursor, err := collection.Find(ctx, bson.M{"status": "aktif"})
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}
	defer cursor.Close(ctx)

	var peminjaman []model.Peminjaman
	if err := cursor.All(ctx, &peminjaman); err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return helper.SuccessResponse(c, peminjaman)
}
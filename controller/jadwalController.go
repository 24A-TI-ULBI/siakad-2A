package controller

import (
	"context"
	"net/http"

	"backend/helper"
	"backend/model"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetAllJadwal handles GET /jadwal
func GetAllJadwal(c *fiber.Ctx) error {
	db := helper.GetDB()
	filter := bson.M{}

	prodi := c.Query("prodi")
	if prodi != "" {
		filter["prodi"] = prodi
	}

	dosen := c.Query("dosen")
	if dosen != "" {
		filter["dosen"] = bson.M{"$regex": dosen, "$options": "i"} // Case-insensitive search
	}

	var jadwals []model.Jadwal
	cursor, err := db.Collection("jadwal").Find(context.TODO(), filter)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Gagal mengambil jadwal",
			"error":   err.Error(),
		})
	}

	if err = cursor.All(context.TODO(), &jadwals); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Gagal parsing jadwal",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Jadwal berhasil diambil",
		"data":    jadwals,
	})
}

// GetJadwalByID handles GET /jadwal/:id
func GetJadwalByID(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "ID tidak valid",
		})
	}

	db := helper.GetDB()
	var jadwal model.Jadwal

	err = db.Collection("jadwal").FindOne(context.TODO(), bson.M{"_id": id}).Decode(&jadwal)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"status":  "error",
				"message": "Jadwal tidak ditemukan",
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Gagal mengambil jadwal",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Jadwal ditemukan",
		"data":    jadwal,
	})
}

// CreateJadwal handles POST /jadwal
func CreateJadwal(c *fiber.Ctx) error {
	var jadwal model.Jadwal
	if err := c.BodyParser(&jadwal); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Input tidak valid",
			"error":   err.Error(),
		})
	}

	db := helper.GetDB()
	res, err := db.Collection("jadwal").InsertOne(context.TODO(), jadwal)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Gagal menyimpan jadwal",
			"error":   err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "Jadwal berhasil ditambahkan",
		"data":    res.InsertedID,
	})
}

// UpdateJadwal handles PUT /jadwal/:id
func UpdateJadwal(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "ID tidak valid",
		})
	}

	var data map[string]interface{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Input tidak valid",
			"error":   err.Error(),
		})
	}

	db := helper.GetDB()
	update := bson.M{"$set": data}

	res, err := db.Collection("jadwal").UpdateOne(context.TODO(), bson.M{"_id": id}, update)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Gagal update jadwal",
			"error":   err.Error(),
		})
	}

	if res.MatchedCount == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "Jadwal tidak ditemukan",
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Jadwal berhasil diupdate",
	})
}

// DeleteJadwal handles DELETE /jadwal/:id
func DeleteJadwal(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "ID tidak valid",
		})
	}

	db := helper.GetDB()
	res, err := db.Collection("jadwal").DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Gagal menghapus jadwal",
			"error":   err.Error(),
		})
	}

	if res.DeletedCount == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "Jadwal tidak ditemukan",
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Jadwal berhasil dihapus",
	})
}
// GetAllRuangan handles GET /ruangan
func GetAllRuangan(c *fiber.Ctx) error {
	db := helper.GetDB()
	filter := bson.M{}

	var ruangans []model.Ruangan
	cursor, err := db.Collection("ruangan").Find(context.TODO(), filter)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Gagal mengambil data ruangan",
			"error":   err.Error(),
		})
	}

	if err = cursor.All(context.TODO(), &ruangans); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Gagal parsing data ruangan",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Data ruangan berhasil diambil",
		"data":    ruangans,
	})
}

// GetRuanganByKode handles GET /ruangan/:kode
func GetRuanganByKode(c *fiber.Ctx) error {
	kode := c.Params("kode")
	db := helper.GetDB()

	var ruangan model.Ruangan
	err := db.Collection("ruangan").FindOne(context.TODO(), bson.M{"kode": kode}).Decode(&ruangan)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"status":  "error",
				"message": "Ruangan tidak ditemukan",
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Gagal mengambil data ruangan",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Data ruangan ditemukan",
		"data":    ruangan,
	})
}

// CreateRuangan handles POST /ruangan
func CreateRuangan(c *fiber.Ctx) error {
	var ruangan model.Ruangan
	if err := c.BodyParser(&ruangan); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Input tidak valid",
			"error":   err.Error(),
		})
	}

	db := helper.GetDB()
	
	// Cek apakah kode sudah ada
	err := db.Collection("ruangan").FindOne(context.TODO(), bson.M{"kode": ruangan.Kode}).Err()
	if err == nil {
		return c.Status(http.StatusConflict).JSON(fiber.Map{
			"status":  "error",
			"message": "Kode ruangan sudah digunakan",
		})
	}

	res, err := db.Collection("ruangan").InsertOne(context.TODO(), ruangan)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Gagal menyimpan ruangan",
			"error":   err.Error(),
		})
	}

	ruangan.ID = res.InsertedID.(bson.M)["_id"].(bson.M)["_id"].(primitive.ObjectID) // fallback atau skip, lebih baik reload id 
	// (Go primitive.ObjectID assertion):
	// res.InsertedID is primitive.ObjectID. 
	
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "Ruangan berhasil ditambahkan",
		"data":    res.InsertedID,
	})
}

// UpdateRuangan handles PUT /ruangan/:kode
func UpdateRuangan(c *fiber.Ctx) error {
	kode := c.Params("kode")
	var data map[string]interface{}

	if err := c.BodyParser(&data); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Input tidak valid",
			"error":   err.Error(),
		})
	}

	db := helper.GetDB()
	update := bson.M{"$set": data}

	res, err := db.Collection("ruangan").UpdateOne(context.TODO(), bson.M{"kode": kode}, update)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Gagal update ruangan",
			"error":   err.Error(),
		})
	}

	if res.MatchedCount == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "Ruangan tidak ditemukan",
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Ruangan berhasil diupdate",
	})
}

package controller

import (
	"context"
	"net/http"

	"backend/helper" // <--- PASTIKAN INI ADALAH helper, BUKAN config
	"backend/model"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// === BACKEND 1: MANAJEMEN PKL ===

func GetAllPKL(c *fiber.Ctx) error {
	collection := helper.GetDB().Collection("pkl")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"success": false, "message": err.Error()})
	}
	defer cursor.Close(context.Background())

	var results []model.PKL = []model.PKL{}
	if err = cursor.All(context.Background(), &results); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"success": false, "message": err.Error()})
	}
	return c.JSON(fiber.Map{"success": true, "data": results})
}

func GetPKLByNPM(c *fiber.Ctx) error {
	npm := c.Params("npm")
	collection := helper.GetDB().Collection("pkl")
	var result model.PKL

	err := collection.FindOne(context.Background(), bson.M{"npm": npm}).Decode(&result)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"success": false, "message": "Data tidak ditemukan"})
	}
	return c.JSON(fiber.Map{"success": true, "data": result})
}

func CreatePKL(c *fiber.Ctx) error {
	var pkl model.PKL
	if err := c.BodyParser(&pkl); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"success": false, "message": err.Error()})
	}

	pkl.ID = primitive.NewObjectID()
	if pkl.Status == "" {
		pkl.Status = "Pending"
	}

	collection := helper.GetDB().Collection("pkl")
	_, err := collection.InsertOne(context.Background(), pkl)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"success": false, "message": err.Error()})
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{"success": true, "data": pkl})
}

func UpdatePKL(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"success": false, "message": "ID tidak valid"})
	}

	var pkl model.PKL
	if err := c.BodyParser(&pkl); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"success": false, "message": err.Error()})
	}

	collection := helper.GetDB().Collection("pkl")
	update := bson.M{"$set": bson.M{
		"perusahaan": pkl.Perusahaan,
		"periode":    pkl.Periode,
		"pembimbing": pkl.Pembimbing,
		"status":     pkl.Status,
	}}

	_, err = collection.UpdateOne(context.Background(), bson.M{"_id": id}, update)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"success": false, "message": err.Error()})
	}
	return c.JSON(fiber.Map{"success": true, "message": "Data PKL berhasil diperbarui"})
}

func DeletePKL(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"success": false, "message": "ID tidak valid"})
	}

	collection := helper.GetDB().Collection("pkl")
	_, err = collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"success": false, "message": err.Error()})
	}
	return c.JSON(fiber.Map{"success": true, "message": "Data PKL berhasil dihapus"})
}

// === BACKEND 2: MANAJEMEN LAPORAN & NILAI ===

func SubmitLaporan(c *fiber.Ctx) error {
	var laporan model.LaporanPKL
	if err := c.BodyParser(&laporan); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"success": false, "message": err.Error()})
	}

	laporan.ID = primitive.NewObjectID()
	collection := helper.GetDB().Collection("pkl")
	_, err := collection.InsertOne(context.Background(), laporan)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"success": false, "message": err.Error()})
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{"success": true, "data": laporan})
}

func GetLaporanByNPM(c *fiber.Ctx) error {
	npm := c.Params("npm")
	collection := helper.GetDB().Collection("pkl")
	cursor, err := collection.Find(context.Background(), bson.M{"npm": npm})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"success": false, "message": err.Error()})
	}
	defer cursor.Close(context.Background())

	var results []model.LaporanPKL = []model.LaporanPKL{}
	_ = cursor.All(context.Background(), &results)
	return c.JSON(fiber.Map{"success": true, "data": results})
}

func UpdateLaporan(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"success": false, "message": "ID tidak valid"})
	}

	var laporan model.LaporanPKL
	if err := c.BodyParser(&laporan); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"success": false, "message": err.Error()})
	}

	collection := helper.GetDB().Collection("pkl")
	update := bson.M{"$set": bson.M{
		"url_laporan": laporan.URLLaporan,
		"nilai":       laporan.Nilai,
		"keterangan":  laporan.Keterangan,
	}}

	_, err = collection.UpdateOne(context.Background(), bson.M{"_id": id}, update)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"success": false, "message": err.Error()})
	}
	return c.JSON(fiber.Map{"success": true, "message": "Laporan/Nilai berhasil diperbarui"})
}

func GetNilaiPKL(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"success": false, "message": "ID tidak valid"})
	}

	collection := helper.GetDB().Collection("pkl")
	var result model.LaporanPKL
	err = collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&result)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"success": false, "message": "Data laporan tidak ditemukan"})
	}
	return c.JSON(fiber.Map{"success": true, "data": result})
}

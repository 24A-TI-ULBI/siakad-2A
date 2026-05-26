package controller

import (
	"errors"

	"backend/helper"
	"backend/model"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	ormawaCollection   = "ormawa"
	kegiatanCollection = "kegiatan"
)

func GetAllOrmawa(c *fiber.Ctx) error {
	db := helper.GetDB()
	ormawa, err := helper.GetAllDoc[model.Ormawa](db, ormawaCollection, bson.M{})
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "gagal mengambil data ormawa")
	}
	return helper.SuccessResponse(c, ormawa)
}

func GetOrmawaByID(c *fiber.Ctx) error {
	id, err := objectIDParam(c)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	db := helper.GetDB()
	ormawa, err := helper.GetOneDoc[model.Ormawa](db, ormawaCollection, bson.M{"_id": id})
	if err != nil {
		return handleMongoReadError(c, err, "ormawa tidak ditemukan", "gagal mengambil detail ormawa")
	}
	return helper.SuccessResponse(c, ormawa)
}

func CreateOrmawa(c *fiber.Ctx) error {
	var payload model.Ormawa
	if err := c.BodyParser(&payload); err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "body request tidak valid")
	}

	db := helper.GetDB()
	insertedID, err := helper.InsertOneDoc(db, ormawaCollection, payload)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "gagal menambahkan ormawa")
	}

	return helper.SuccessResponse(c, fiber.Map{"inserted_id": insertedID})
}

func UpdateOrmawa(c *fiber.Ctx) error {
	id, err := objectIDParam(c)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	var payload model.Ormawa
	if err := c.BodyParser(&payload); err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "body request tidak valid")
	}

	update := bson.M{
		"$set": bson.M{
			"nama":      payload.Nama,
			"deskripsi": payload.Deskripsi,
			"ketua":     payload.Ketua,
			"pengurus":  payload.Pengurus,
		},
	}

	db := helper.GetDB()
	result, err := helper.UpdateDoc(db, ormawaCollection, bson.M{"_id": id}, update)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "gagal mengupdate ormawa")
	}
	if result.MatchedCount == 0 {
		return helper.ErrorResponse(c, fiber.StatusNotFound, "ormawa tidak ditemukan")
	}

	return helper.SuccessResponse(c, fiber.Map{"modified_count": result.ModifiedCount})
}

func DeleteOrmawa(c *fiber.Ctx) error {
	id, err := objectIDParam(c)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	db := helper.GetDB()
	result, err := helper.DeleteDoc(db, ormawaCollection, bson.M{"_id": id})
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "gagal menghapus ormawa")
	}
	if result.DeletedCount == 0 {
		return helper.ErrorResponse(c, fiber.StatusNotFound, "ormawa tidak ditemukan")
	}

	return helper.SuccessResponse(c, fiber.Map{"deleted_count": result.DeletedCount})
}

func GetAllKegiatan(c *fiber.Ctx) error {
	db := helper.GetDB()
	kegiatan, err := helper.GetAllDoc[model.Kegiatan](db, kegiatanCollection, bson.M{})
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "gagal mengambil data kegiatan")
	}
	return helper.SuccessResponse(c, kegiatan)
}

func GetKegiatanByID(c *fiber.Ctx) error {
	id, err := objectIDParam(c)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	db := helper.GetDB()
	kegiatan, err := helper.GetOneDoc[model.Kegiatan](db, kegiatanCollection, bson.M{"_id": id})
	if err != nil {
		return handleMongoReadError(c, err, "kegiatan tidak ditemukan", "gagal mengambil detail kegiatan")
	}
	return helper.SuccessResponse(c, kegiatan)
}

func CreateKegiatan(c *fiber.Ctx) error {
	var payload model.Kegiatan
	if err := c.BodyParser(&payload); err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "body request tidak valid")
	}

	db := helper.GetDB()
	insertedID, err := helper.InsertOneDoc(db, kegiatanCollection, payload)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "gagal menambahkan kegiatan")
	}

	return helper.SuccessResponse(c, fiber.Map{"inserted_id": insertedID})
}

func UpdateKegiatan(c *fiber.Ctx) error {
	id, err := objectIDParam(c)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	var payload model.Kegiatan
	if err := c.BodyParser(&payload); err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "body request tidak valid")
	}

	update := bson.M{
		"$set": bson.M{
			"ormawa_id": payload.OrmawaID,
			"nama":      payload.Nama,
			"deskripsi": payload.Deskripsi,
			"tanggal":   payload.Tanggal,
			"lokasi":    payload.Lokasi,
		},
	}

	db := helper.GetDB()
	result, err := helper.UpdateDoc(db, kegiatanCollection, bson.M{"_id": id}, update)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "gagal mengupdate kegiatan")
	}
	if result.MatchedCount == 0 {
		return helper.ErrorResponse(c, fiber.StatusNotFound, "kegiatan tidak ditemukan")
	}

	return helper.SuccessResponse(c, fiber.Map{"modified_count": result.ModifiedCount})
}

func GetKegiatanByOrmawa(c *fiber.Ctx) error {
	id, err := objectIDParam(c)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	db := helper.GetDB()
	kegiatan, err := helper.GetAllDoc[model.Kegiatan](db, kegiatanCollection, bson.M{"ormawa_id": id})
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "gagal mengambil kegiatan ormawa")
	}
	return helper.SuccessResponse(c, kegiatan)
}

func objectIDParam(c *fiber.Ctx) (primitive.ObjectID, error) {
	idParam := c.Params("id")
	if idParam == "" {
		return primitive.NilObjectID, errors.New("id tidak boleh kosong")
	}

	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return primitive.NilObjectID, errors.New("id tidak valid")
	}
	return id, nil
}

func handleMongoReadError(c *fiber.Ctx, err error, notFoundMessage string, fallbackMessage string) error {
	if errors.Is(err, mongo.ErrNoDocuments) {
		return helper.ErrorResponse(c, fiber.StatusNotFound, notFoundMessage)
	}
	return helper.ErrorResponse(c, fiber.StatusInternalServerError, fallbackMessage)
}

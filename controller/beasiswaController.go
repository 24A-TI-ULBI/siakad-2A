package controller

import (
	"backend/helper"
	"backend/model"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetBeasiswa — GET /beasiswa
func GetBeasiswa(c *fiber.Ctx) error {
	col := helper.GetCollection("beasiswa")
	ctx, cancel := helper.GetContext()
	defer cancel()

	cursor, err := col.Find(ctx, bson.D{})
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal mengambil data beasiswa")
	}
	defer cursor.Close(ctx)

	var list []model.Beasiswa
	if err := cursor.All(ctx, &list); err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal membaca data beasiswa")
	}
	if list == nil {
		list = []model.Beasiswa{}
	}
	return helper.SuccessResponse(c, list)
}

// GetDetailBeasiswa — GET /beasiswa/:id
func GetDetailBeasiswa(c *fiber.Ctx) error {
	id, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "ID tidak valid")
	}

	col := helper.GetCollection("beasiswa")
	ctx, cancel := helper.GetContext()
	defer cancel()

	var item model.Beasiswa
	if err := col.FindOne(ctx, bson.M{"_id": id}).Decode(&item); err != nil {
		return helper.ErrorResponse(c, fiber.StatusNotFound, "Beasiswa tidak ditemukan")
	}
	return helper.SuccessResponse(c, item)
}

// AddBeasiswa — POST /beasiswa
func AddBeasiswa(c *fiber.Ctx) error {
	var item model.Beasiswa
	if err := c.BodyParser(&item); err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Format request tidak valid")
	}
	if item.Nama == "" {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Nama beasiswa wajib diisi")
	}

	item.ID = primitive.NewObjectID()
	col := helper.GetCollection("beasiswa")
	ctx, cancel := helper.GetContext()
	defer cancel()

	result, err := col.InsertOne(ctx, item)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal menyimpan data beasiswa")
	}
	return helper.SuccessResponse(c, fiber.Map{"inserted_id": result.InsertedID})
}

// UpdateBeasiswa — PUT /beasiswa/:id
func UpdateBeasiswa(c *fiber.Ctx) error {

	npm := c.Params("npm")

	var body bson.M

	if err := c.BodyParser(&body); err != nil {

		return helper.ErrorResponse(
			c,
			fiber.StatusBadRequest,
			"Format request tidak valid",
		)

	}

	col := helper.GetCollection("pendaftaran")

	ctx, cancel := helper.GetContext()
	defer cancel()

	result, err := col.UpdateOne(
		ctx,
		bson.M{"npm": npm},
		bson.M{"$set": body},
	)

	if err != nil {

		return helper.ErrorResponse(
			c,
			fiber.StatusInternalServerError,
			"Gagal update status",
		)

	}

	if result.MatchedCount == 0 {

		return helper.ErrorResponse(
			c,
			fiber.StatusNotFound,
			"Data tidak ditemukan",
		)

	}

	return helper.SuccessResponse(
		c,
		fiber.Map{
			"updated": result.ModifiedCount,
		},
	)

}

// DeleteBeasiswa — DELETE /beasiswa/:id
func DeleteBeasiswa(c *fiber.Ctx) error {

	npm := c.Params("npm")

	col := helper.GetCollection("pendaftaran")

	ctx, cancel := helper.GetContext()
	defer cancel()

	result, err := col.DeleteOne(
		ctx,
		bson.M{"npm": npm},
	)

	if err != nil {

		return helper.ErrorResponse(
			c,
			fiber.StatusInternalServerError,
			"Gagal menghapus data",
		)

	}

	if result.DeletedCount == 0 {

		return helper.ErrorResponse(
			c,
			fiber.StatusNotFound,
			"Data tidak ditemukan",
		)

	}

	return helper.SuccessResponse(
		c,
		fiber.Map{
			"deleted": result.DeletedCount,
		},
	)

}
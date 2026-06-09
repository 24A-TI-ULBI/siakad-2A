package controller

import (
	"errors"
	"strings"
	"time"

	"backend/helper"
	"backend/model"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dosenCollection   = "dosen"
	jabatanCollection = "jabatan"
)

func GetAllDosen(c *fiber.Ctx) error {
	db := helper.GetDB()
	ctx, cancel := helper.GetContext()
	defer cancel()

	filter := bson.M{}
	if jabatanID := strings.TrimSpace(c.Query("jabatan_id")); jabatanID != "" {
		filter["jabatan_id"] = jabatanID
	}

	cursor, err := db.Collection(dosenCollection).Find(ctx, filter, options.Find().SetSort(bson.D{{Key: "nama", Value: 1}}))
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal mengambil data dosen")
	}
	defer cursor.Close(ctx)

	dosen := make([]model.Dosen, 0)
	if err := cursor.All(ctx, &dosen); err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal membaca data dosen")
	}
	return helper.SuccessResponse(c, dosen)
}

func GetDosenByNIDN(c *fiber.Ctx) error {
	nidn := strings.TrimSpace(c.Params("nidn"))
	if nidn == "" {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "NIDN wajib diisi")
	}

	dosen, err := helper.GetOneDoc[model.Dosen](helper.GetDB(), dosenCollection, bson.M{"nidn": nidn})
	if errors.Is(err, mongo.ErrNoDocuments) {
		return helper.ErrorResponse(c, fiber.StatusNotFound, "Data dosen tidak ditemukan")
	}
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal mengambil detail dosen")
	}
	return helper.SuccessResponse(c, dosen)
}

func CreateDosen(c *fiber.Ctx) error {
	var dosen model.Dosen
	if err := c.BodyParser(&dosen); err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Format data dosen tidak valid")
	}

	if err := prepareDosenPayload(&dosen, true); err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	db := helper.GetDB()
	ctx, cancel := helper.GetContext()
	defer cancel()

	total, err := db.Collection(dosenCollection).CountDocuments(ctx, bson.M{"nidn": dosen.NIDN})
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal memeriksa NIDN dosen")
	}
	if total > 0 {
		return helper.ErrorResponse(c, fiber.StatusConflict, "NIDN dosen sudah terdaftar")
	}

	now := time.Now()
	dosen.ID = primitive.NilObjectID
	dosen.CreatedAt = now
	dosen.UpdatedAt = now

	insertedID, err := helper.InsertOneDoc(db, dosenCollection, dosen)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal menambah data dosen")
	}

	return helper.SuccessResponse(c, fiber.Map{
		"inserted_id": insertedID,
		"dosen":       dosen,
	})
}

func UpdateDosen(c *fiber.Ctx) error {
	nidn := strings.TrimSpace(c.Params("nidn"))
	if nidn == "" {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "NIDN wajib diisi")
	}

	var dosen model.Dosen
	if err := c.BodyParser(&dosen); err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Format data dosen tidak valid")
	}

	dosen.NIDN = nidn
	if err := prepareDosenPayload(&dosen, false); err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	update := bson.M{
		"$set": bson.M{
			"nama":         dosen.Nama,
			"email":        dosen.Email,
			"telepon":      dosen.Telepon,
			"prodi":        dosen.Prodi,
			"pendidikan":   dosen.Pendidikan,
			"status":       dosen.Status,
			"jabatan_id":   dosen.JabatanID,
			"jabatan_nama": dosen.JabatanNama,
			"updated_at":   time.Now(),
		},
	}

	result, err := helper.UpdateDoc(helper.GetDB(), dosenCollection, bson.M{"nidn": nidn}, update)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal memperbarui data dosen")
	}
	if result.MatchedCount == 0 {
		return helper.ErrorResponse(c, fiber.StatusNotFound, "Data dosen tidak ditemukan")
	}

	return helper.SuccessResponse(c, fiber.Map{
		"matched_count":  result.MatchedCount,
		"modified_count": result.ModifiedCount,
	})
}

func DeleteDosen(c *fiber.Ctx) error {
	nidn := strings.TrimSpace(c.Params("nidn"))
	if nidn == "" {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "NIDN wajib diisi")
	}

	result, err := helper.DeleteDoc(helper.GetDB(), dosenCollection, bson.M{"nidn": nidn})
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal menghapus data dosen")
	}
	if result.DeletedCount == 0 {
		return helper.ErrorResponse(c, fiber.StatusNotFound, "Data dosen tidak ditemukan")
	}

	return helper.SuccessResponse(c, fiber.Map{"deleted_count": result.DeletedCount})
}

func GetAllJabatan(c *fiber.Ctx) error {
	db := helper.GetDB()
	ctx, cancel := helper.GetContext()
	defer cancel()

	cursor, err := db.Collection(jabatanCollection).Find(ctx, bson.M{}, options.Find().SetSort(bson.D{{Key: "jenis", Value: 1}, {Key: "nama", Value: 1}}))
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal mengambil data jabatan")
	}
	defer cursor.Close(ctx)

	jabatan := make([]model.Jabatan, 0)
	if err := cursor.All(ctx, &jabatan); err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal membaca data jabatan")
	}
	return helper.SuccessResponse(c, jabatan)
}

func GetJabatanByID(c *fiber.Ctx) error {
	id, err := objectIDFromParam(c.Params("id"))
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "ID jabatan tidak valid")
	}

	jabatan, err := helper.GetOneDoc[model.Jabatan](helper.GetDB(), jabatanCollection, bson.M{"_id": id})
	if errors.Is(err, mongo.ErrNoDocuments) {
		return helper.ErrorResponse(c, fiber.StatusNotFound, "Data jabatan tidak ditemukan")
	}
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal mengambil detail jabatan")
	}
	return helper.SuccessResponse(c, jabatan)
}

func CreateJabatan(c *fiber.Ctx) error {
	var jabatan model.Jabatan
	if err := c.BodyParser(&jabatan); err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Format data jabatan tidak valid")
	}

	if err := prepareJabatanPayload(&jabatan); err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	now := time.Now()
	jabatan.ID = primitive.NilObjectID
	jabatan.CreatedAt = now
	jabatan.UpdatedAt = now

	insertedID, err := helper.InsertOneDoc(helper.GetDB(), jabatanCollection, jabatan)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal menambah data jabatan")
	}

	return helper.SuccessResponse(c, fiber.Map{
		"inserted_id": insertedID,
		"jabatan":     jabatan,
	})
}

func UpdateJabatan(c *fiber.Ctx) error {
	id, err := objectIDFromParam(c.Params("id"))
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "ID jabatan tidak valid")
	}

	var jabatan model.Jabatan
	if err := c.BodyParser(&jabatan); err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Format data jabatan tidak valid")
	}

	if err := prepareJabatanPayload(&jabatan); err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	update := bson.M{
		"$set": bson.M{
			"nama":       jabatan.Nama,
			"jenis":      jabatan.Jenis,
			"deskripsi":  jabatan.Deskripsi,
			"updated_at": time.Now(),
		},
	}

	result, err := helper.UpdateDoc(helper.GetDB(), jabatanCollection, bson.M{"_id": id}, update)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal memperbarui data jabatan")
	}
	if result.MatchedCount == 0 {
		return helper.ErrorResponse(c, fiber.StatusNotFound, "Data jabatan tidak ditemukan")
	}

	return helper.SuccessResponse(c, fiber.Map{
		"matched_count":  result.MatchedCount,
		"modified_count": result.ModifiedCount,
	})
}

func DeleteJabatan(c *fiber.Ctx) error {
	id, err := objectIDFromParam(c.Params("id"))
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "ID jabatan tidak valid")
	}

	result, err := helper.DeleteDoc(helper.GetDB(), jabatanCollection, bson.M{"_id": id})
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal menghapus data jabatan")
	}
	if result.DeletedCount == 0 {
		return helper.ErrorResponse(c, fiber.StatusNotFound, "Data jabatan tidak ditemukan")
	}

	return helper.SuccessResponse(c, fiber.Map{"deleted_count": result.DeletedCount})
}

func prepareDosenPayload(dosen *model.Dosen, requireNIDN bool) error {
	dosen.NIDN = strings.TrimSpace(dosen.NIDN)
	dosen.Nama = strings.TrimSpace(dosen.Nama)
	dosen.Email = strings.TrimSpace(dosen.Email)
	dosen.Telepon = strings.TrimSpace(dosen.Telepon)
	dosen.Prodi = strings.TrimSpace(dosen.Prodi)
	dosen.Pendidikan = strings.TrimSpace(dosen.Pendidikan)
	dosen.Status = strings.TrimSpace(dosen.Status)
	dosen.JabatanID = strings.TrimSpace(dosen.JabatanID)
	dosen.JabatanNama = strings.TrimSpace(dosen.JabatanNama)

	if requireNIDN && dosen.NIDN == "" {
		return errors.New("NIDN wajib diisi")
	}
	if dosen.Nama == "" {
		return errors.New("Nama dosen wajib diisi")
	}
	if dosen.Email == "" {
		return errors.New("Email dosen wajib diisi")
	}
	if dosen.Prodi == "" {
		return errors.New("Program studi wajib diisi")
	}
	if dosen.Pendidikan == "" {
		dosen.Pendidikan = "S2"
	}
	if dosen.Status == "" {
		dosen.Status = "Aktif"
	}
	if dosen.JabatanID == "" {
		return errors.New("Jabatan dosen wajib dipilih")
	}

	jabatan, err := findJabatanByHexID(dosen.JabatanID)
	if err != nil {
		return err
	}
	dosen.JabatanNama = jabatan.Nama
	return nil
}

func prepareJabatanPayload(jabatan *model.Jabatan) error {
	jabatan.Nama = strings.TrimSpace(jabatan.Nama)
	jabatan.Jenis = strings.TrimSpace(jabatan.Jenis)
	jabatan.Deskripsi = strings.TrimSpace(jabatan.Deskripsi)

	if jabatan.Nama == "" {
		return errors.New("Nama jabatan wajib diisi")
	}
	if jabatan.Jenis == "" {
		jabatan.Jenis = "Fungsional"
	}
	return nil
}

func findJabatanByHexID(hexID string) (model.Jabatan, error) {
	id, err := objectIDFromParam(hexID)
	if err != nil {
		return model.Jabatan{}, errors.New("ID jabatan tidak valid")
	}

	jabatan, err := helper.GetOneDoc[model.Jabatan](helper.GetDB(), jabatanCollection, bson.M{"_id": id})
	if errors.Is(err, mongo.ErrNoDocuments) {
		return model.Jabatan{}, errors.New("Jabatan dosen tidak ditemukan")
	}
	if err != nil {
		return model.Jabatan{}, errors.New("Gagal memeriksa jabatan dosen")
	}
	return jabatan, nil
}

func objectIDFromParam(param string) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(strings.TrimSpace(param))
}

package controller

import (
	"context"
	"time"

	"backend/helper"
	"backend/model"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetAllKuesioner handles GET /kuesioner
func GetAllKuesioner(c *fiber.Ctx) error {
	db := helper.GetDB()
	kuesioners, err := helper.GetAllDoc[model.Kuesioner](db, "kuesioner", bson.M{})
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal mengambil data kuesioner: "+err.Error())
	}
	return helper.SuccessResponse(c, kuesioners)
}

// GetKuesionerByID handles GET /kuesioner/:id
func GetKuesionerByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		// Jika ID bukan ObjectID valid (misal: "index.html"), lanjutkan ke handler berikutnya (static file)
		return c.Next()
	}

	db := helper.GetDB()
	kuesioner, err := helper.GetOneDoc[model.Kuesioner](db, "kuesioner", bson.M{"_id": objID})
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusNotFound, "Kuesioner tidak ditemukan: "+err.Error())
	}

	return helper.SuccessResponse(c, kuesioner)
}

// CreateKuesioner handles POST /kuesioner
func CreateKuesioner(c *fiber.Ctx) error {
	var kuesioner model.Kuesioner
	if err := c.BodyParser(&kuesioner); err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Body request tidak valid: "+err.Error())
	}

	if kuesioner.Title == "" {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Judul kuesioner wajib diisi")
	}

	kuesioner.ID = primitive.NewObjectID()
	kuesioner.CreatedAt = primitive.NewDateTimeFromTime(time.Now())

	if kuesioner.Questions == nil {
		kuesioner.Questions = []model.Question{}
	}

	db := helper.GetDB()
	_, err := helper.InsertOneDoc(db, "kuesioner", kuesioner)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal menyimpan kuesioner: "+err.Error())
	}

	return helper.SuccessResponse(c, kuesioner)
}

// UpdateKuesioner handles PUT /kuesioner/:id
func UpdateKuesioner(c *fiber.Ctx) error {
	idStr := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Format ID kuesioner tidak valid")
	}

	var updateData model.Kuesioner
	if err := c.BodyParser(&updateData); err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Body request tidak valid: "+err.Error())
	}

	if updateData.Title == "" {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Judul kuesioner wajib diisi")
	}

	db := helper.GetDB()
	_, err = helper.GetOneDoc[model.Kuesioner](db, "kuesioner", bson.M{"_id": objID})
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusNotFound, "Kuesioner tidak ditemukan")
	}

	update := bson.M{
		"$set": bson.M{
			"title":       updateData.Title,
			"description": updateData.Description,
			"questions":   updateData.Questions,
		},
	}

	_, err = helper.UpdateDoc(db, "kuesioner", bson.M{"_id": objID}, update)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal memperbarui kuesioner: "+err.Error())
	}

	updateData.ID = objID
	return helper.SuccessResponse(c, updateData)
}

// DeleteKuesioner handles DELETE /kuesioner/:id
func DeleteKuesioner(c *fiber.Ctx) error {
	idStr := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Format ID kuesioner tidak valid")
	}

	db := helper.GetDB()
	_, err = helper.GetOneDoc[model.Kuesioner](db, "kuesioner", bson.M{"_id": objID})
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusNotFound, "Kuesioner tidak ditemukan")
	}

	_, err = helper.DeleteDoc(db, "kuesioner", bson.M{"_id": objID})
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal menghapus kuesioner: "+err.Error())
	}

	// Hapus semua jawaban terkait jika kuesioner dihapus
	_, _ = db.Collection("jawaban_kuesioner").DeleteMany(context.Background(), bson.M{"kuesioner_id": objID})

	return helper.SuccessResponse(c, fiber.Map{
		"message": "Kuesioner dan jawaban terkait berhasil dihapus",
	})
}

// SubmitJawaban handles POST /kuesioner/jawab
func SubmitJawaban(c *fiber.Ctx) error {
	var jawabanInput struct {
		KuesionerID string             `json:"kuesioner_id"`
		NPM         string             `json:"npm"`
		Nama        string             `json:"nama"`
		Answers     []model.AnswerItem `json:"answers"`
	}

	if err := c.BodyParser(&jawabanInput); err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Body request tidak valid: "+err.Error())
	}

	if jawabanInput.KuesionerID == "" || jawabanInput.NPM == "" || jawabanInput.Nama == "" {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Kuesioner ID, NPM, dan Nama wajib diisi")
	}

	kObjID, err := primitive.ObjectIDFromHex(jawabanInput.KuesionerID)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Format ID kuesioner tidak valid")
	}

	db := helper.GetDB()

	// Pastikan kuesioner ada
	_, err = helper.GetOneDoc[model.Kuesioner](db, "kuesioner", bson.M{"_id": kObjID})
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusNotFound, "Kuesioner tidak ditemukan")
	}

	// Cek apakah sudah pernah dijawab oleh NPM ini
	var existing model.JawabanKuesioner
	existingErr := db.Collection("jawaban_kuesioner").FindOne(context.Background(), bson.M{
		"kuesioner_id": kObjID,
		"npm":          jawabanInput.NPM,
	}).Decode(&existing)

	if existingErr == nil {
		// Update jawaban yang ada
		update := bson.M{
			"$set": bson.M{
				"nama":         jawabanInput.Nama,
				"answers":      jawabanInput.Answers,
				"submitted_at": primitive.NewDateTimeFromTime(time.Now()),
			},
		}
		_, err = db.Collection("jawaban_kuesioner").UpdateOne(context.Background(), bson.M{"_id": existing.ID}, update)
		if err != nil {
			return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal memperbarui jawaban: "+err.Error())
		}
		existing.Nama = jawabanInput.Nama
		existing.Answers = jawabanInput.Answers
		existing.SubmittedAt = primitive.NewDateTimeFromTime(time.Now())
		return helper.SuccessResponse(c, existing)
	}

	// Insert baru
	jawaban := model.JawabanKuesioner{
		ID:          primitive.NewObjectID(),
		KuesionerID: kObjID,
		NPM:         jawabanInput.NPM,
		Nama:        jawabanInput.Nama,
		Answers:     jawabanInput.Answers,
		SubmittedAt: primitive.NewDateTimeFromTime(time.Now()),
	}

	_, err = helper.InsertOneDoc(db, "jawaban_kuesioner", jawaban)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal menyimpan jawaban: "+err.Error())
	}

	return helper.SuccessResponse(c, jawaban)
}

// GetJawabanByKuesionerID handles GET /kuesioner/jawaban/:id
func GetJawabanByKuesionerID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	kObjID, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "Format ID kuesioner tidak valid")
	}

	db := helper.GetDB()
	jawabans, err := helper.GetAllDoc[model.JawabanKuesioner](db, "jawaban_kuesioner", bson.M{"kuesioner_id": kObjID})
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal mengambil jawaban kuesioner: "+err.Error())
	}

	return helper.SuccessResponse(c, jawabans)
}

// GetStatusByNPM handles GET /kuesioner/status/:npm
func GetStatusByNPM(c *fiber.Ctx) error {
	npm := c.Params("npm")
	if npm == "" {
		return helper.ErrorResponse(c, fiber.StatusBadRequest, "NPM wajib diisi")
	}

	db := helper.GetDB()
	kuesioners, err := helper.GetAllDoc[model.Kuesioner](db, "kuesioner", bson.M{})
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal mengambil kuesioner: "+err.Error())
	}

	jawabans, err := helper.GetAllDoc[model.JawabanKuesioner](db, "jawaban_kuesioner", bson.M{"npm": npm})
	if err != nil {
		return helper.ErrorResponse(c, fiber.StatusInternalServerError, "Gagal mengambil jawaban: "+err.Error())
	}

	sudahDijawab := make(map[string]bool)
	for _, j := range jawabans {
		sudahDijawab[j.KuesionerID.Hex()] = true
	}

	var statusList []model.KuesionerStatus
	for _, k := range kuesioners {
		status := "belum"
		if sudahDijawab[k.ID.Hex()] {
			status = "sudah"
		}
		statusList = append(statusList, model.KuesionerStatus{
			KuesionerID: k.ID,
			Title:       k.Title,
			Description: k.Description,
			Status:      status,
		})
	}

	return helper.SuccessResponse(c, statusList)
}

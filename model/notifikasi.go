package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Notifikasi menyimpan pesan kampus yang dikirim ke mahasiswa.
type Notifikasi struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	NPM         string             `bson:"npm" json:"npm"`
	Judul       string             `bson:"judul" json:"judul"`
	Pesan       string             `bson:"pesan" json:"pesan"`
	Tipe        string             `bson:"tipe" json:"tipe"`
	Prioritas   string             `bson:"prioritas" json:"prioritas"`
	Dibaca      bool               `bson:"dibaca" json:"dibaca"`
	DibuatPada  time.Time          `bson:"dibuat_pada" json:"dibuat_pada"`
	DibacaPada  *time.Time         `bson:"dibaca_pada,omitempty" json:"dibaca_pada,omitempty"`
	DikirimOleh string             `bson:"dikirim_oleh" json:"dikirim_oleh"`
}

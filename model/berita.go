package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Berita represents a campus news article.
type Berita struct {
	ID      primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Judul   string             `json:"judul" bson:"judul"`
	Isi     string             `json:"isi" bson:"isi"`
	Penulis string             `json:"penulis" bson:"penulis"`
	Tanggal time.Time          `json:"tanggal" bson:"tanggal"`
}

// Komentar represents a comment on a news article.
type Komentar struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	BeritaID primitive.ObjectID `json:"berita_id" bson:"berita_id"`
	Nama     string             `json:"nama" bson:"nama"`
	Pesan    string             `json:"pesan" bson:"pesan"`
	Tanggal  time.Time          `json:"tanggal" bson:"tanggal"`
}

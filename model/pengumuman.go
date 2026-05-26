package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Pengumuman represents academic announcements
type Pengumuman struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Judul     string             `bson:"judul" json:"judul"`
	Isi       string             `bson:"isi" json:"isi"`
	Kategori  string             `bson:"kategori" json:"kategori"`
	Tanggal   time.Time          `bson:"tanggal" json:"tanggal"`
	Penulis   string             `bson:"penulis,omitempty" json:"penulis,omitempty"`
	FileUrl   string             `bson:"file_url,omitempty" json:"file_url,omitempty"`
	FileType  string             `bson:"file_type,omitempty" json:"file_type,omitempty"`
}

// Kategori represents announcement categories
type Kategori struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Nama string             `bson:"nama" json:"nama"`
}

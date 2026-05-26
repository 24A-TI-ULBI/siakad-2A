package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Buku struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Judul      string             `json:"judul" bson:"judul"`
	Penulis    string             `json:"penulis" bson:"penulis"`
	Penerbit   string             `json:"penerbit" bson:"penerbit"`
	Tahun      int                `json:"tahun" bson:"tahun"`
	Stok       int                `json:"stok" bson:"stok"`
	Tersedia   bool               `json:"tersedia" bson:"tersedia"`
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type Peminjaman struct {
	ID             primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	BukuID         primitive.ObjectID `json:"buku_id" bson:"buku_id"`
	NPM            string             `json:"npm" bson:"npm"`
	NamaMahasiswa  string             `json:"nama_mahasiswa" bson:"nama_mahasiswa"`
	TanggalPinjam  time.Time          `json:"tanggal_pinjam" bson:"tanggal_pinjam"`
	TanggalKembali *time.Time         `json:"tanggal_kembali,omitempty" bson:"tanggal_kembali,omitempty"`
	Status         string             `json:"status" bson:"status"` // aktif / kembali
	CreatedAt      time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt      time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
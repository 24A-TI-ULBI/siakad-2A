package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Dosen struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	NIDN        string             `bson:"nidn" json:"nidn"`
	Nama        string             `bson:"nama" json:"nama"`
	Email       string             `bson:"email" json:"email"`
	Telepon     string             `bson:"telepon" json:"telepon"`
	Prodi       string             `bson:"prodi" json:"prodi"`
	Pendidikan  string             `bson:"pendidikan" json:"pendidikan"`
	Status      string             `bson:"status" json:"status"`
	JabatanID   string             `bson:"jabatan_id" json:"jabatan_id"`
	JabatanNama string             `bson:"jabatan_nama" json:"jabatan_nama"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}

type Jabatan struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Nama      string             `bson:"nama" json:"nama"`
	Jenis     string             `bson:"jenis" json:"jenis"`
	Deskripsi string             `bson:"deskripsi" json:"deskripsi"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

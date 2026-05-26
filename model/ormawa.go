package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Ormawa struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Nama      string             `json:"nama" bson:"nama"`
	Deskripsi string             `json:"deskripsi" bson:"deskripsi"`
	Ketua     string             `json:"ketua" bson:"ketua"`
	Pengurus  []string           `json:"pengurus" bson:"pengurus"`
}

type Kegiatan struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	OrmawaID  primitive.ObjectID `json:"ormawa_id" bson:"ormawa_id"`
	Nama      string             `json:"nama" bson:"nama"`
	Deskripsi string             `json:"deskripsi" bson:"deskripsi"`
	Tanggal   string             `json:"tanggal" bson:"tanggal"`
	Lokasi    string             `json:"lokasi" bson:"lokasi"`
}

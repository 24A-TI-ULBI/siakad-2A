package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Prestasi struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	NPM       string             `bson:"npm" json:"npm"`
	NamaEvent string             `bson:"nama_event" json:"nama_event"`
	Tingkat   string             `bson:"tingkat" json:"tingkat"` // e.g. Lokal, Regional, Nasional, Internasional
	Juara     string             `bson:"juara" json:"juara"`     // e.g. Juara 1, Harapan 1, dll
	Tanggal   string             `bson:"tanggal" json:"tanggal"` // format: YYYY-MM-DD
	Kategori  string             `bson:"kategori" json:"kategori"` // nama kategori relasi
}

type KategoriPrestasi struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Nama string             `bson:"nama" json:"nama"` // e.g. Akademik, Non-Akademik, Olahraga, Seni
}

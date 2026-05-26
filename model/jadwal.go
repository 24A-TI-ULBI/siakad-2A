package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Jadwal merepresentasikan jadwal perkuliahan
type Jadwal struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	MataKuliah  string             `bson:"mata_kuliah,omitempty" json:"mata_kuliah,omitempty"`
	Dosen       string             `bson:"dosen,omitempty" json:"dosen,omitempty"`
	RuanganKode string             `bson:"ruangan_kode,omitempty" json:"ruangan_kode,omitempty"`
	Hari        string             `bson:"hari,omitempty" json:"hari,omitempty"`
	JamMulai    string             `bson:"jam_mulai,omitempty" json:"jam_mulai,omitempty"`
	JamSelesai  string             `bson:"jam_selesai,omitempty" json:"jam_selesai,omitempty"`
	Prodi       string             `bson:"prodi,omitempty" json:"prodi,omitempty"`
}

// Ruangan merepresentasikan data ruangan kelas
type Ruangan struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Kode      string             `bson:"kode,omitempty" json:"kode,omitempty"`
	Nama      string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Kapasitas int                `bson:"kapasitas,omitempty" json:"kapasitas,omitempty"`
	Fasilitas []string           `bson:"fasilitas,omitempty" json:"fasilitas,omitempty"`
	Gedung    string             `bson:"gedung,omitempty" json:"gedung,omitempty"`
}
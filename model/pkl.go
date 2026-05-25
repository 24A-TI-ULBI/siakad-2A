package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// PKL melambangkan data pendaftaran PKL mahasiswa
type PKL struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	NPM        string             `bson:"npm" json:"npm"`
	Perusahaan string             `bson:"perusahaan" json:"perusahaan"`
	Periode    string             `bson:"periode" json:"periode"`
	Pembimbing string             `bson:"pembimbing" json:"pembimbing"`
	Status     string             `bson:"status" json:"status"` // Contoh: "Pending", "Disetujui"
}

// LaporanPKL melambangkan submission laporan dan nilainya
type LaporanPKL struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	NPM        string             `bson:"npm" json:"npm"`
	URLLaporan string             `bson:"url_laporan" json:"url_laporan"`
	Nilai      int                `bson:"nilai" json:"nilai"`
	Keterangan string             `bson:"keterangan" json:"keterangan"` // Contoh: "Lulus", "Revisi"
}
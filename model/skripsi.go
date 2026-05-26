package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Skripsi represents a student's thesis/final project data
type Skripsi struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	NPM              string             `bson:"npm" json:"npm"`
	NamaMahasiswa    string             `bson:"nama_mahasiswa" json:"nama_mahasiswa"`
	Judul            string             `bson:"judul" json:"judul"`
	DosenPembimbing1 string             `bson:"dosen_pembimbing_1" json:"dosen_pembimbing_1"`
	NIDNPembimbing1  string             `bson:"nidn_pembimbing_1" json:"nidn_pembimbing_1"`
	DosenPembimbing2 string             `bson:"dosen_pembimbing_2,omitempty" json:"dosen_pembimbing_2,omitempty"`
	NIDNPembimbing2  string             `bson:"nidn_pembimbing_2,omitempty" json:"nidn_pembimbing_2,omitempty"`
	Status           string             `bson:"status" json:"status"` // pengajuan, berjalan, revisi, selesai
	TanggalDaftar    string             `bson:"tanggal_daftar" json:"tanggal_daftar"`
	TanggalSelesai   string             `bson:"tanggal_selesai,omitempty" json:"tanggal_selesai,omitempty"`
}

// Bimbingan represents a thesis consultation/guidance session record
type Bimbingan struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	NPM         string             `bson:"npm" json:"npm"`
	NIDN        string             `bson:"nidn" json:"nidn"`
	NamaDosen   string             `bson:"nama_dosen" json:"nama_dosen"`
	Tanggal     string             `bson:"tanggal" json:"tanggal"`
	Catatan     string             `bson:"catatan" json:"catatan"`
	Status      string             `bson:"status" json:"status"` // disetujui, revisi, menunggu
	ProgressBab string             `bson:"progress_bab,omitempty" json:"progress_bab,omitempty"`
}

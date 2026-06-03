package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SuratKeterangan represents a letter template (e.g. Surat Keterangan Mahasiswa Aktif)
type SuratKeterangan struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Kode        string             `bson:"kode" json:"kode"`
	Nama        string             `bson:"nama" json:"nama"`
	Deskripsi   string             `bson:"deskripsi" json:"deskripsi"`
	Persyaratan []string           `bson:"persyaratan" json:"persyaratan"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}

// PengajuanSurat represents a student request for a letter
type PengajuanSurat struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	SuratID          string             `bson:"surat_id" json:"surat_id"` // Matches SuratKeterangan hex ID
	SuratNama        string             `bson:"surat_nama" json:"surat_nama"` // Cached template name for ease of display
	NPM              string             `bson:"npm" json:"npm"`
	NamaMahasiswa    string             `bson:"nama_mahasiswa" json:"nama_mahasiswa"`
	Prodi            string             `bson:"prodi" json:"prodi"`
	Keterangan       string             `bson:"keterangan" json:"keterangan"` // Purpose of request
	Status           string             `bson:"status" json:"status"` // "proses", "selesai", "ditolak"
	Catatan          string             `bson:"catatan" json:"catatan"` // Admin notes (e.g. reason for rejection)
	TanggalPengajuan time.Time          `bson:"tanggal_pengajuan" json:"tanggal_pengajuan"`
	TanggalUpdate    time.Time          `bson:"tanggal_update" json:"tanggal_update"`
}

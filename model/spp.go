package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type PembayaranSPP struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	NPM              string             `bson:"npm" json:"npm"`
	NamaMahasiswa    string             `bson:"nama_mahasiswa" json:"nama_mahasiswa"`
	Semester         string             `bson:"semester" json:"semester"`
	TahunAjaran      string             `bson:"tahun_ajaran" json:"tahun_ajaran"`
	JumlahTagihan    float64            `bson:"jumlah_tagihan" json:"jumlah_tagihan"`
	StatusPembayaran string             `bson:"status_pembayaran" json:"status_pembayaran"`
	TanggalTagihan   string             `bson:"tanggal_tagihan" json:"tanggal_tagihan"`
	TanggalBayar     string             `bson:"tanggal_bayar,omitempty" json:"tanggal_bayar,omitempty"`
	MetodePembayaran string             `bson:"metode_pembayaran,omitempty" json:"metode_pembayaran,omitempty"`
}

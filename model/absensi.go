package model

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

type AbsensiID string

func (id *AbsensiID) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	raw := bson.RawValue{Type: t, Value: data}

	if value, ok := raw.StringValueOK(); ok {
		*id = AbsensiID(value)
		return nil
	}

	if value, ok := raw.ObjectIDOK(); ok {
		*id = AbsensiID(value.Hex())
		return nil
	}

	return fmt.Errorf("tipe _id absensi tidak didukung: %s", t)
}

type Absensi struct {
	ID         AbsensiID `json:"id" bson:"_id"`
	NPM        string    `json:"npm" bson:"npm"`
	MatkulCode string    `json:"matkulCode" bson:"matkulCode"`
	Tanggal    string    `json:"tanggal" bson:"tanggal"`
	Status     string    `json:"status" bson:"status"`
	Timestamp  string    `json:"timestamp" bson:"timestamp"`
}

type RekapMatkul struct {
	MatkulCode string  `json:"matkulCode"`
	Hadir      int     `json:"hadir"`
	Izin       int     `json:"izin"`
	Alpha      int     `json:"alpha"`
	TotalSesi  int     `json:"totalSesi"`
	Persentase float64 `json:"persentase"`
}

type StudentRekap struct {
	NPM                 string        `json:"npm"`
	TotalSesi           int           `json:"totalSesi"`
	Hadir               int           `json:"hadir"`
	Izin                int           `json:"izin"`
	Alpha               int           `json:"alpha"`
	PersentaseKehadiran float64       `json:"persentaseKehadiran"`
	RekapPerMatkul      []RekapMatkul `json:"rekapPerMatkul"`
}

type StudentMatkulRekap struct {
	NPM        string  `json:"npm"`
	Hadir      int     `json:"hadir"`
	Izin       int     `json:"izin"`
	Alpha      int     `json:"alpha"`
	TotalSesi  int     `json:"totalSesi"`
	Persentase float64 `json:"persentase"`
}

type CourseRekap struct {
	MatkulCode string               `json:"matkulCode"`
	Pendaftar  []StudentMatkulRekap `json:"pendaftar"`
}

package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Mahasiswa struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	NPM      string             `bson:"npm"      json:"npm"`
	Nama     string             `bson:"nama"     json:"nama"`
	Email    string             `bson:"email"    json:"email"`
	Phone    string             `bson:"phone"    json:"phone"`
	Prodi    string             `bson:"prodi"    json:"prodi"`
	Angkatan int                `bson:"angkatan" json:"angkatan"`
}

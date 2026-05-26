package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Pendaftaran struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	NamaMahasiswa string             `bson:"nama"        json:"nama"`
	NPM           string             `bson:"npm"         json:"npm"`
	Email         string             `bson:"email"       json:"email"`
	Semester      string             `bson:"semester"    json:"semester"`
	Prodi         string             `bson:"prodi"       json:"prodi"`
	IPK           string             `bson:"ipk"         json:"ipk"`
	BeasiswaID    string             `bson:"beasiswa_id" json:"beasiswa_id"`
	Beasiswa      string             `bson:"beasiswa"    json:"beasiswa"`
	Status        string             `bson:"status"      json:"status"`
}

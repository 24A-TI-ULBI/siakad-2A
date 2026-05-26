package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Beasiswa struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama     string             `bson:"nama"     json:"nama"`
	Syarat   string             `bson:"syarat"   json:"syarat"`
	Deadline string             `bson:"deadline" json:"deadline"`
}

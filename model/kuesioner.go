package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Question struct {
	ID          string   `bson:"id" json:"id"`
	Text        string   `bson:"text" json:"text"`
	Type        string   `bson:"type" json:"type"` // "text", "rating", "choice"
	Options     []string `bson:"options,omitempty" json:"options,omitempty"`
	IsRequired  bool     `bson:"is_required" json:"is_required"`
}

type Kuesioner struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	Questions   []Question         `bson:"questions" json:"questions"`
	CreatedAt   primitive.DateTime `bson:"created_at" json:"created_at"`
}

type AnswerItem struct {
	QuestionID string `bson:"question_id" json:"question_id"`
	Value      string `bson:"value" json:"value"`
}

type JawabanKuesioner struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	KuesionerID primitive.ObjectID `bson:"kuesioner_id" json:"kuesioner_id"`
	NPM         string             `bson:"npm" json:"npm"`
	Nama        string             `bson:"nama" json:"nama"`
	Answers     []AnswerItem       `bson:"answers" json:"answers"`
	SubmittedAt primitive.DateTime `bson:"submitted_at" json:"submitted_at"`
}

type KuesionerStatus struct {
	KuesionerID primitive.ObjectID `json:"kuesioner_id"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Status      string             `json:"status"` // "sudah" atau "belum"
}

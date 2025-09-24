package models

import "go.mongodb.org/mongo-driver/bson/primitive"


type Opinion struct {
	UserID    primitive.ObjectID `bson:"user_id,omitempty" json:"user_id,omitempty"`
	Text      string             `bson:"text" json:"text"`
	CreatedAt int64              `bson:"created_at" json:"created_at"`
}
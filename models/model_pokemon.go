package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type PokemonStats struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name  string             `bson:"name" json:"name"`
	Likes int                `bson:"likes" json:"likes"`
	Opinions []Opinion          `bson:"opinions" json:"opinions"`
}


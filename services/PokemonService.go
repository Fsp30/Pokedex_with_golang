package services

import (
	"context"
	"time"

	"github.com/Fsp30/Pokedex_with_golang/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PokemonService struct{
	Collection *mongo.Collection
}

func (s *PokemonService) InsertPokemon(pokemon models.PokemonStats) (*mongo.InsertOneResult, error){
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return s.Collection.InsertOne(ctx, pokemon)
}

func (s *PokemonService) GetPokemon(name string) (*models.PokemonStats, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var pokemon models.PokemonStats
	err := s.Collection.FindOne(ctx, bson.M{"name": name}).Decode(&pokemon)
	if err != nil {
		return nil, err
	}

	return &pokemon, nil
}

func (s *PokemonService) AddLikesPokemon(pokemonID primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{
		"$inc": bson.M{"likes": 1},
	}

	_, err := s.Collection.UpdateByID(ctx, pokemonID, update)
	return  err
}

func (s *PokemonService) AddOpinion(pokemonID primitive.ObjectID, opinion models.Opinion) error{
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	opinion.CreatedAt = time.Now().Unix()

	filter := bson.M{"_id": pokemonID}
	update := bson.M{
		"$push" : bson.M{"opinions": opinion},
	}

	_, err := s.Collection.UpdateOne(ctx, filter, update)
	return err
}

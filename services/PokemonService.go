package services

import (
	"context"
	"time"
	
	"github.com/Fsp30/Pokedex_with_golang/models"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"



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

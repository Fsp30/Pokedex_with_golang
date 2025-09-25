package services

import (
	"context"
	"time"
	
	"github.com/Fsp30/Pokedex_with_golang/models"

	"go.mongodb.org/mongo-driver/mongo"

)

type UserService struct {
	Collection *mongo.Collection
}

func (s *UserService) CreateUser(user models.User) (*mongo.InsertOneResult, error){
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return s.Collection.InsertOne(ctx, user)
}
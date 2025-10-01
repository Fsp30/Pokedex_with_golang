package main

import (
	"fmt"
	"log"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/gin-contrib/cors"

	"github.com/Fsp30/Pokedex_with_golang/handlers"
	"github.com/Fsp30/Pokedex_with_golang/services"
	"github.com/Fsp30/Pokedex_with_golang/config"
)

func main() {
	fmt.Println("Nike")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, client := config.SetupDB()
 
	pokemonService := &services.PokemonService{
		Collection: db.Collection("pokemons"),
	}

	userService := &services.UserService{
		Collection: db.Collection("users"),
	}

	pokemonHandler := &handlers.PokemonHandler{
		Service: pokemonService,
	}

	userHandler := &handlers.UserHandler{
		Service: userService,
	}


	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.GET("/pokemon/:name", pokemonHandler.GetPokemonInfo)
	r.POST("/pokemon/:name/opinion", pokemonHandler.AddOpinion)
	r.POST("/pokemon/:name/like", pokemonHandler.AddLike)
	r.GET("/pokemon/:name/opinions", pokemonHandler.ListOpinions)

	r.POST("/users", userHandler.CreateUser)

	defer func() {
	if err := client.Disconnect(context.Background()); err != nil {
		log.Printf("Error disconnect MongoDB: %v", err)
	} else {
		log.Println("MongoDB disconnect successfully")
	}
	}()
	
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

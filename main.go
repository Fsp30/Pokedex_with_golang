package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/Fsp30/Pokedex_with_golang/handlers"
)


func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}


	r := gin.Default()

	r.GET("/pokemon/:name", handlers.GetPokemonInfo)
	r.POST("/pokemon/:name/opinion", handlers.AddOpinion)
	r.POST("/pokemon/:name/like", handlers.AddLike)
	r.GET("/pokemon/:name/opinions", handlers.ListOpinions)

	r.Run(":8080")
}

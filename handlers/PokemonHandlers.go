package handlers

import (
	"strings"

	"net/http"

	"github.com/Fsp30/Pokedex_with_golang/models"
	"github.com/Fsp30/Pokedex_with_golang/services"
	"github.com/Fsp30/Pokedex_with_golang/storage"
	poke_go "github.com/JoshGuarino/PokeGo/pkg"
	"github.com/gin-gonic/gin"
)

type PokemonHandler struct {
	Service *services.PokemonService
}

var client = poke_go.NewClient()

func (h *PokemonHandler) GetPokemonInfo(c *gin.Context) {
	name := c.Param("name")

	if name == ""{
		c.JSON(http.StatusBadRequest, gin.H{"error" : "Pokémon name not provided"})
		return
	}

	name = strings.ToLower(name)

	apiPokemon, err := client.Pokemon.GetPokemon(name)

	if err != nil{
		c.JSON(http.StatusNotFound, gin.H{"error": "Pokemon Not Found in API PokeApi"})
	}

	pokemon, err := h.Service.GetPokemon(name)

	if err != nil{
		pokemon = &models.PokemonStats{
			Name: name,
			Likes: 0,
			Opinions: []models.Opinion{},
		}
	}

	response := gin.H{
		"api_data": apiPokemon,
		"likes": pokemon.Likes,
		"opinions": pokemon.Opinions,
	}

	c.JSON(http.StatusOK, response)

}

func  (h *PokemonHandler) AddLike(c *gin.Context) {
	name := c.Param("name")
	if name == ""{
		c.JSON(http.StatusBadRequest, gin.H{"error" : "Pokémon name not provided"})
		return
	}

	name = strings.ToLower(name)

	pokemon, err :=  h.Service.GetPokemon(name)

	if err != nil{
		pokemon = &models.PokemonStats{
			Name:     name,
			Likes:    0,
			Opinions: []models.Opinion{},
		}
		_, err := h.Service.InsertPokemon(*pokemon)
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create Pokémon record"})
			return
		}
	}
	
	err = h.Service.AddLikesPokemon(pokemon.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to register like"})
		return

	}


	c.JSON(http.StatusOK, gin.H{"message": "Like registrado"})
}

func AddOpinion(c *gin.Context) {
	name := c.Param("name")
	var body struct {
		User    string `json:"user"`
		Opinion string `json:"opinion"`
	}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}
	storage.SaveOpinion(name, body.User, body.Opinion)
	c.JSON(http.StatusOK, gin.H{"message": "Opinião registrada"})
}


func ListOpinions(c *gin.Context) {
	name := c.Param("name")
	opinions := storage.GetOpinions(name)
	likes := storage.GetLikes(name)
	c.JSON(http.StatusOK, gin.H{
		"opinions": opinions,
		"likes":    likes,
	})
}

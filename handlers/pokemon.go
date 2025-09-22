package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	poke_go "github.com/JoshGuarino/PokeGo/pkg"
	"github.com/Fsp30/Pokedex_with_golang/storage"
)

var client = poke_go.NewClient()

func GetPokemonInfo(c *gin.Context) {
	name := c.Param("name")
	pokemon, err := client.Pokemon.GetPokemon(name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pokémon não encontrado"})
		return
	}
	c.JSON(http.StatusOK, pokemon)
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

func AddLike(c *gin.Context) {
	name := c.Param("name")
	storage.AddLike(name)
	c.JSON(http.StatusOK, gin.H{"message": "Like registrado"})
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

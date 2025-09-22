package main 

import "fmt"
import(
	poke_go "github.com/JoshGuarino/PokeGo/pkg"
)

func main(){
	client := poke_go.NewClient()

	pokemon, err := client.Pokemon.GetPokemon("amaura")
	
	if err != nil{
		fmt.Println("Failed search pokemon: ", err)
		return
	}

	fmt.Println("Nome:", pokemon.Name)
	fmt.Println("Altura:", pokemon.Height)
	fmt.Println("Peso:", pokemon.Weight)

}
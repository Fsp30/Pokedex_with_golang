package main 

import "fmt"
import(
	poke_go "github.com/JoshGuarino/PokeGo/pkg"
)

func main(){
	client := poke_go.NewClient()
	fmt.Println("Teste go")
}
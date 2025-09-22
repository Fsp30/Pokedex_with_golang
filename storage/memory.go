package storage

var pokemonOpinions = make(map[string][]string)
var pokemonLikes = make(map[string]int)

func SaveOpinion(pokemon, user, opinion string) {
	entry := user + ": " + opinion
	pokemonOpinions[pokemon] = append(pokemonOpinions[pokemon], entry)
}

func AddLike(pokemon string) {
	pokemonLikes[pokemon]++
}

func GetOpinions(pokemon string) []string {
	return pokemonOpinions[pokemon]
}

func GetLikes(pokemon string) int {
	return pokemonLikes[pokemon]
}

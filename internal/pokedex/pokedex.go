package pokedex

import "github.com/FerMeza/pokedexcli/internal/pokeapi"

var Pokedex map[string]Pokemon = map[string]Pokemon{}

type Pokemon struct {
	Name           string
	BaseExperience int
}

func MapPokemonApiToDomain(pokemon *pokeapi.Pokemon) Pokemon {
	return Pokemon{
		Name:           pokemon.Name,
		BaseExperience: pokemon.BaseExperience,
	}
}

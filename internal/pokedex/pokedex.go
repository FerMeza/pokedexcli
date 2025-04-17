package pokedex

import (
	"github.com/FerMeza/pokedexcli/internal/pokeapi"
)

var Pokedex map[string]Pokemon = map[string]Pokemon{}

type Pokemon struct {
	Name   string
	Height int
	Weight int
	Stats  []struct {
		Name     string
		BaseStat int
	}
	Types          []string
	BaseExperience int
}

func MapPokemonApiToDomain(pokemon *pokeapi.Pokemon) Pokemon {
	poke := Pokemon{
		Name:   pokemon.Name,
		Height: pokemon.Height,
		Weight: pokemon.Weight,
		Stats: make([]struct {
			Name     string
			BaseStat int
		}, 0),
		Types:          make([]string, 0),
		BaseExperience: pokemon.BaseExperience,
	}

	for _, stat := range pokemon.Stats {
		poke.Stats = append(poke.Stats, struct {
			Name     string
			BaseStat int
		}{
			Name:     stat.Stat.Name,
			BaseStat: stat.BaseStat,
		})
	}

	for _, t := range pokemon.Types {
		poke.Types = append(poke.Types, t.Type.Name)
	}

	return poke
}

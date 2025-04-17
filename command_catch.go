package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/FerMeza/pokedexcli/internal/pokedex"
)

func commandCatch(config *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("missing pokemon argument")
	}

	pokemonName := args[0]
	pokemon, ok := pokedex.Pokedex[pokemonName]
	if !ok {
		pokemonDetail, err := config.pokeAPIClient.GetPokemon(pokemonName)
		if err != nil {
			return err
		}
		pokemon = pokedex.MapPokemonApiToDomain(&pokemonDetail)
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	baseExp := pokemon.BaseExperience
	const currMaxBaseExperience = 608

	if baseExp > 608 {
		baseExp = 608
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	catchPercentVal := r.Float64()

	supExp := float64(currMaxBaseExperience) * (1 / 0.99)

	percentNonCatch := float64(baseExp) / supExp

	caught := false

	if catchPercentVal > percentNonCatch {
		caught = true
	}

	if caught {
		if !ok {
			pokedex.Pokedex[pokemonName] = pokemon
		}
		fmt.Printf("%s was caught!\n", pokemonName)
		fmt.Println("You may now inspect it with the inspect command.")
		return nil
	}

	fmt.Printf("%s escaped!\n", pokemonName)
	return nil
}

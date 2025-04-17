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
	pokemonDetail, err := config.pokeAPIClient.GetPokemon(pokemonName)

	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	baseExp := pokemonDetail.BaseExperience
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
		_, ok := pokedex.Pokedex[pokemonName]
		if !ok {
			pokemon := pokedex.Pokemon{
				Name:           pokemonName,
				BaseExperience: pokemonDetail.BaseExperience,
			}
			pokedex.Pokedex[pokemonName] = pokemon
		}
		fmt.Printf("%s was caught!\n", pokemonName)
		return nil
	}

	fmt.Printf("%s escaped!\n", pokemonName)
	return nil
}

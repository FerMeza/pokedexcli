package main

import (
	"errors"
	"fmt"

	"github.com/FerMeza/pokedexcli/internal/pokedex"
)

func commandInspect(config *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("missing pokemon argument")
	}

	pokemonName := args[0]
	pokemon, ok := pokedex.Pokedex[pokemonName]
	if !ok {
		fmt.Println("You have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", pokemonName)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)

	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t)
	}

	return nil
}

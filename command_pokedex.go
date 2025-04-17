package main

import (
	"fmt"

	"github.com/FerMeza/pokedexcli/internal/pokedex"
)

func commandPokedex(config *config, args ...string) error {
	fmt.Println("Your Pokedex:")

	for _, pokemon := range pokedex.Pokedex {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil
}

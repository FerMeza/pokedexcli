package main

import (
	"errors"
	"fmt"
)

func commandExplore(config *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("missing location argument")
	}

	location := args[0]
	locationDetail, err := config.pokeAPIClient.GetLocation(location)

	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", location)

	if len(locationDetail.PokemonEncounters) == 0 {
		fmt.Printf("No pokemons found")
		return nil
	}

	fmt.Println("Found Pokemon:")

	for _, pokemon := range locationDetail.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}

	return nil
}

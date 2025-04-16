package main

import "fmt"

func commandMap(config *config) error {
	if config == nil {
		panic("Invalid state on command map")
	}

	pokeAreas, err := config.pokeAPIClient.ListLocations(config.nextLocationsURL)
	if err != nil {
		return err
	}

	for _, area := range pokeAreas.Results {
		fmt.Println(area.Name)
	}

	config.nextLocationsURL = pokeAreas.Next
	config.previousLocationsURL = pokeAreas.Previous

	return nil
}

func commandMapB(config *config) error {
	if config == nil {
		panic("Invalid state on command mapb")
	}

	pokeAreas, err := config.pokeAPIClient.ListLocations(config.previousLocationsURL)
	if err != nil {
		return err
	}

	for _, area := range pokeAreas.Results {
		fmt.Println(area.Name)
	}

	config.nextLocationsURL = pokeAreas.Next
	config.previousLocationsURL = pokeAreas.Previous

	return nil
}

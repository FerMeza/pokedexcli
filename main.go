package main

import (
	"time"

	"github.com/FerMeza/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeAPIClient: pokeClient,
	}

	startRepl(cfg)
}

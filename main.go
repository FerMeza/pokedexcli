package main

import (
	"time"

	"github.com/FerMeza/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		pokeAPIClient: pokeClient,
	}

	startRepl(cfg)
}

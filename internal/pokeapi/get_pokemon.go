package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName

	var pokemon Pokemon

	val, ok := c.cache.Get(url)
	if !ok {
		req, err := http.NewRequest(http.MethodGet, url, nil)

		if err != nil {
			return Pokemon{}, fmt.Errorf("error creating request: %w", err)
		}

		res, err := c.httpClient.Do(req)

		if res.StatusCode == http.StatusNotFound {
			return Pokemon{}, fmt.Errorf("pokemon %s not found", pokemonName)
		}

		if err != nil {
			return Pokemon{}, fmt.Errorf("error sending request: %w", err)
		}

		defer res.Body.Close()

		val, err = io.ReadAll(res.Body)

		if err != nil {
			return Pokemon{}, err
		}

		c.cache.Add(url, val)
	}

	err := json.Unmarshal(val, &pokemon)

	if err != nil {
		return Pokemon{}, fmt.Errorf("error unmarshaling val: %w", err)
	}

	return pokemon, nil
}

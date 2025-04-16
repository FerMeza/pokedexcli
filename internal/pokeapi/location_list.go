package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area" + "/?offset=0&limit=20"
	if pageURL != nil {
		url = *pageURL
	}

	var pokeAreas RespShallowLocations

	val, ok := c.cache.Get(url)
	if !ok {
		req, err := http.NewRequest(http.MethodGet, url, nil)

		if err != nil {
			return RespShallowLocations{}, fmt.Errorf("error creating request: %w", err)
		}

		res, err := c.httpClient.Do(req)

		if err != nil {
			return RespShallowLocations{}, fmt.Errorf("error sending request: %w", err)
		}

		defer res.Body.Close()

		val, err = io.ReadAll(res.Body)
		if err != nil {
			return RespShallowLocations{}, err
		}

		c.cache.Add(url, val)
	}

	err := json.Unmarshal(val, &pokeAreas)

	if err != nil {
		return RespShallowLocations{}, fmt.Errorf("error unmarshaling val: %w", err)
	}

	return pokeAreas, nil
}

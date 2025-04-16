package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return RespShallowLocations{}, fmt.Errorf("error creating request: %v", err)
	}

	res, err := c.httpClient.Do(req)

	if err != nil {
		return RespShallowLocations{}, fmt.Errorf("error sending request: %v", err)
	}

	defer res.Body.Close()

	var pokeAreas RespShallowLocations

	err = json.NewDecoder(res.Body).Decode(&pokeAreas)
	if err != nil {
		return RespShallowLocations{}, fmt.Errorf("error decoding response: %v", err)
	}

	return pokeAreas, nil
}

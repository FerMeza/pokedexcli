package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationDetail(locationName string) (RespFullLocation, error) {
	url := baseURL + "/location-area/" + locationName

	var pokeLocation RespFullLocation

	val, ok := c.cache.Get(url)
	if !ok {
		req, err := http.NewRequest(http.MethodGet, url, nil)

		if err != nil {
			return RespFullLocation{}, fmt.Errorf("error creating request: %w", err)
		}

		res, err := c.httpClient.Do(req)

		if res.StatusCode == http.StatusNotFound {
			return RespFullLocation{}, fmt.Errorf("location %s not found", locationName)
		}

		if err != nil {
			return RespFullLocation{}, fmt.Errorf("error sending request: %w", err)
		}

		defer res.Body.Close()

		val, err = io.ReadAll(res.Body)

		if err != nil {
			return RespFullLocation{}, err
		}

		c.cache.Add(url, val)
	}

	err := json.Unmarshal(val, &pokeLocation)

	if err != nil {
		return RespFullLocation{}, fmt.Errorf("error unmarshaling val: %w", err)
	}

	return pokeLocation, nil
}

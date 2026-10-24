package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) ListPokemonFromArea(params string) (DeepLocationResponse, error) {
	if len(params) <= 0 {
		return DeepLocationResponse{}, errors.New("No location provided")
	}
	url := baseURL + "/location-area/" + params
	response, exists := c.pokeCache.Get(url)
	if exists {
		locationResponse := DeepLocationResponse{}
		err := json.Unmarshal(response, &locationResponse)
		if err != nil {
			return DeepLocationResponse{}, err
		}

		return locationResponse, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return DeepLocationResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return DeepLocationResponse{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return DeepLocationResponse{}, err
	}

	locationResponse := DeepLocationResponse{}
	err = json.Unmarshal(dat, &locationResponse)
	if err != nil {
		return DeepLocationResponse{}, err
	}
	c.pokeCache.Add(url, dat)
	return locationResponse, nil
}

package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) GetPokemonInfo(pageURL *string, params ...string) (PokemonEndpointResponse, error) {
	if len(params) <= 0 {
		return PokemonEndpointResponse{}, errors.New("No pokemon provided")
	}
	url := baseURL + "/pokemon/" + params[0]
	if pageURL != nil {
		url = *pageURL
	}
	response, exists := c.pokeCache.Get(url)
	if exists {
		pokemonResponse := PokemonEndpointResponse{}
		err := json.Unmarshal(response, &pokemonResponse)
		if err != nil {
			return PokemonEndpointResponse{}, err
		}

		return pokemonResponse, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonEndpointResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonEndpointResponse{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonEndpointResponse{}, err
	}

	pokemonResponse := PokemonEndpointResponse{}
	err = json.Unmarshal(dat, &pokemonResponse)
	if err != nil {
		return PokemonEndpointResponse{}, err
	}
	c.pokeCache.Add(url, dat)
	return pokemonResponse, nil
}

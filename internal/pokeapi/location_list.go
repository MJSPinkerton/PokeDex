package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) ListLocations(pageURL *string) (ShallowLocationResponse, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	response, exists := c.pokeCache.Get(url)
	if exists {
		locationsResp := ShallowLocationResponse{}
		err := json.Unmarshal(response, &locationsResp)
		if err != nil {
			return ShallowLocationResponse{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ShallowLocationResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return ShallowLocationResponse{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return ShallowLocationResponse{}, err
	}

	locationsResp := ShallowLocationResponse{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return ShallowLocationResponse{}, err
	}
	c.pokeCache.Add(url, dat)
	return locationsResp, nil
}

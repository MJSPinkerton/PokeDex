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

	return locationsResp, nil
}

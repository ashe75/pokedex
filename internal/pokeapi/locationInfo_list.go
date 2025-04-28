package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocationInfo(pageURL *string, locationName string) (LocationArea, error) {
	url := baseURL + "/location-area/" + locationName
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		locationAreaResp := LocationArea{}
		err := json.Unmarshal(val, &locationAreaResp)
		if err != nil {
			return LocationArea{}, err
		}

		return locationAreaResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	locationAreaResp := LocationArea{}
	err = json.Unmarshal(dat, &locationAreaResp)
	if err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(url, dat)
	return locationAreaResp, nil
}

package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListPokemonInfo(pageURL *string, pokemonName string) (PokemonInfo, error) {
	url := baseURL + "/pokemon/" + pokemonName
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		pokemonInfoResp := PokemonInfo{}
		err := json.Unmarshal(val, &pokemonInfoResp)
		if err != nil {
			return PokemonInfo{}, err
		}

		return pokemonInfoResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonInfo{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonInfo{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonInfo{}, err
	}

	pokemonInfoResp := PokemonInfo{}
	err = json.Unmarshal(dat, &pokemonInfoResp)
	if err != nil {
		return PokemonInfo{}, err
	}

	c.cache.Add(url, dat)
	return pokemonInfoResp, nil
}

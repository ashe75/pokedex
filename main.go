package main

import (
	"time"

	"github.com/ashe75/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)

	config := &Config{
		pokeapiClient: pokeClient,
		CaughtPokemon: make(map[string]pokeapi.PokemonInfo),
	}
	startRepl(config)

}

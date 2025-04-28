package main

import "fmt"

func commandExplore(config *Config, location string) error {
	fmt.Printf("Exploring %s...\n", location)
	locationRespInfo, err := config.pokeapiClient.ListLocationInfo(config.NextLocationAreaURL, location)
	if err != nil {
		return err
	}
	fmt.Println("Found Pokemon:")
	for _, encounter := range locationRespInfo.PokemonEncounters {
		pokemonName := encounter.Pokemon.Name
		fmt.Printf(" - %v\n", pokemonName)
	}
	return nil
}

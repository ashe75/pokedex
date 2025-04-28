package main

import (
	"fmt"
	"math/rand"
	"time"
)

func commandCatch(config *Config, pokemonName string) error {
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	pokemonInfo, err := config.pokeapiClient.ListPokemonInfo(config.NextLocationAreaURL, pokemonName)
	if err != nil {
		return err
	}
	baseExp := pokemonInfo.BaseExperience

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	maxValue := 1000
	minValue := 100
	catchChance := maxValue - baseExp
	if catchChance < minValue {
		catchChance = minValue
	}

	randNum := r.Intn(maxValue)

	caught := randNum < catchChance

	if !caught {
		fmt.Printf("%s escaped!\n", pokemonName)
	} else {
		fmt.Printf("%s was caught!\n", pokemonName)
		config.CaughtPokemon[pokemonName] = pokemonInfo
	}
	return nil
}

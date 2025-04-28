package main

import (
	"fmt"
)

func commandInspect(config *Config, pokemonName string) error {
	info, exists := config.CaughtPokemon[pokemonName]
	if exists {
		fmt.Printf("Name: %s\n", pokemonName)
		fmt.Printf("Height: %v\n", info.Height)
		fmt.Printf("Weight: %v\n", info.Weight)
		fmt.Printf("Stats: \n")
		for _, stat := range info.Stats {
			fmt.Printf(" -%v: %v\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Printf("Types:\n")
		for _, t := range info.Types {
			fmt.Printf(" - %v\n", t.Type.Name)
		}
	} else {
		fmt.Printf("have not caught %s yet\n", pokemonName)
	}

	return nil
}

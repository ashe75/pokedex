package main

import (
	"fmt"
)

func commandPokedex(config *Config, param string) error {
	fmt.Println("Your Pokedex:")
	for n := range config.CaughtPokemon {
		fmt.Printf(" - %v\n", n)
	}
	return nil
}

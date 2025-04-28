package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ashe75/pokedexcli/internal/pokeapi"
)

type Config struct {
	pokeapiClient       pokeapi.Client
	NextLocationAreaURL *string
	PrevLocationAreaURL *string
	CaughtPokemon       map[string]pokeapi.PokemonInfo
}

func startRepl(config *Config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		commandParameter := ""
		if len(words) > 1 {
			commandParameter = words[1]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(config, commandParameter)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	var res []string
	text = strings.ToLower(text)
	text = strings.TrimSpace(text)
	res = strings.Fields(text)
	return res
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Get the list of all the Pokemon on set location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "try to catch pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "inspect pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "show list of caught pokemons",
			callback:    commandPokedex,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

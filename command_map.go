package main

import (
	"errors"
	"fmt"
)

func commandMap(config *Config, parameter string) error {
	locationsResp, err := config.pokeapiClient.ListLocations(config.NextLocationAreaURL)
	if err != nil {
		return err
	}

	config.NextLocationAreaURL = locationsResp.Next
	config.PrevLocationAreaURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(config *Config, parameter string) error {
	if config.PrevLocationAreaURL == nil {
		return errors.New("you're on the first page")
	}

	locationResp, err := config.pokeapiClient.ListLocations(config.PrevLocationAreaURL)
	if err != nil {
		return err
	}

	config.NextLocationAreaURL = locationResp.Next
	config.PrevLocationAreaURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

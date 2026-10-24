package main

import (
	"errors"
	"fmt"
)

func commandMapf(config *config) error {
	locations, err := config.PokeApiClient.ListLocations(config.NextLocationURL)
	if err != nil {
		return err
	}

	config.NextLocationURL = locations.Next
	config.PreviousLocationURL = locations.Previous

	for _, area := range locations.Results {
		fmt.Println(area.Name)
	}

	return nil
}

func commandMapBack(config *config) error {
	if config.PreviousLocationURL == nil {
		return errors.New("you're on the first page")
	}
	locations, err := config.PokeApiClient.ListLocations(config.PreviousLocationURL)
	if err != nil {
		return err
	}

	config.NextLocationURL = locations.Next
	config.PreviousLocationURL = locations.Previous

	for _, area := range locations.Results {
		fmt.Println(area.Name)
	}

	return nil
}

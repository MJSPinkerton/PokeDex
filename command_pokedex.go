package main

import (
	"fmt"
)

func commandPokeDex(config *config, param ...string) error {
	if len(config.PokeDex.CaughtPokemon) <= 0 {
		fmt.Println("No entries in the pokedex")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for key := range config.PokeDex.CaughtPokemon {
		fmt.Printf("  - %s\n", key)
	}
	return nil
}

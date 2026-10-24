package main

import (
	"errors"
	"fmt"
)

func commandInspect(config *config, param ...string) error {
	if len(param) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := param[0]
	pokemon, ok := config.PokeDex.CaughtPokemon[name]
	if !ok {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typeInfo := range pokemon.Types {
		fmt.Println("  -", typeInfo.Type.Name)
	}
	return nil
}

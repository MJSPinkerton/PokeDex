package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func commandCatch(config *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}
	name := args[0]
	pokemon, err := config.PokeApiClient.GetPokemonInfo(nil, name)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %v... \n", name)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if pokemon.BaseExperience >= r.Intn(pokemon.BaseExperience*2) { //highest is 608 i think
		fmt.Printf("%v escaped!\n", pokemon.Name)
	} else {
		fmt.Printf("%v was caught!\n", pokemon.Name)
		fmt.Println("You may now inspect it with the inspect command.")
		config.PokeDex.Add(pokemon)
	}

	return nil
}

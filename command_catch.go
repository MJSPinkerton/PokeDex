package main

import (
	"fmt"
	"math/rand"
	"time"
)

func commandCatch(config *config, param *param) error {
	pokemon, err := config.PokeApiClient.GetPokemonInfo(nil, param.params)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %v... \n", param.params[0])
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if pokemon.BaseExperience >= r.Intn(pokemon.BaseExperience*2) { //highest is 608 i think
		fmt.Printf("%v escaped!\n", pokemon.Name)
	} else {
		fmt.Printf("%v was caught!\n", pokemon.Name)
		config.PokeDex.Add(pokemon)
	}

	return nil
}

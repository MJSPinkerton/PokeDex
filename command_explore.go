package main

import "fmt"

func commandExplore(config *config, param *param) error {
	pokemon, err := config.PokeApiClient.ListPokemonFromArea(nil, param.params)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %v... \n", param.params[0])
	fmt.Println("Found Pokemon:")
	for _, pokemon := range pokemon.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}

	return nil
}

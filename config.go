package main

import "github.com/MJSPinkerton/PokeDex/internal/pokeapi"

type config struct {
	commands            map[string]cliCommand
	PokeApiClient       pokeapi.Client
	NextLocationURL     *string
	PreviousLocationURL *string
}

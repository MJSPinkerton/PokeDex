package main

import (
	"time"

	"github.com/MJSPinkerton/PokeDex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	pokeDex := NewDex()
	config := &config{
		commands:      getCommands(),
		PokeApiClient: pokeClient,
		PokeDex:       pokeDex,
	}
	startRepl(config)
}

package main

import (
	"sync"

	"github.com/MJSPinkerton/PokeDex/internal/pokeapi"
)

type PokeDex struct {
	CaughtPokemon map[string]pokeapi.PokemonEndpointResponse
	mtx           *sync.Mutex
}

func NewDex() PokeDex {

	newDex := PokeDex{
		CaughtPokemon: map[string]pokeapi.PokemonEndpointResponse{},
		mtx:           &sync.Mutex{},
	}
	return newDex
}

func (p *PokeDex) Add(newPokemon pokeapi.PokemonEndpointResponse) {
	p.mtx.Lock()
	defer p.mtx.Unlock()
	p.CaughtPokemon[newPokemon.Name] = newPokemon
}

func (p *PokeDex) Get(key string) (pokeapi.PokemonEndpointResponse, bool) {
	p.mtx.Lock()
	defer p.mtx.Unlock()
	entry, exists := p.CaughtPokemon[key]
	return entry, exists
}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Encounter struct {
	MinLevel        int                `json:"min_level"`
	MaxLevel        int                `json:"max_level"`
	ConditionValues []NamedAPIResource `json:"condition_values"`
	Chance          int                `json:"chance"`
	Method          NamedAPIResource   `json:"method"`
}

type VersionEncounterDetail struct {
	Version          NamedAPIResource `json:"version"`
	MaxChance        int              `json:"max_chance"`
	EncounterDetails []Encounter      `json:"encounter_details"`
}

type PokemonEncounter struct {
	Pokemon        NamedAPIResource         `json:"pokemon"`
	VersionDetails []VersionEncounterDetail `json:"version_detail"`
}

type Name struct {
	Name     string           `json:"name"`
	Language NamedAPIResource `json:"language"`
}

type EncounterVersionDetails struct {
	Rate    int              `json:"rate"`
	Version NamedAPIResource `json:"version"`
}

type NamedAPIResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type EncounterMethodRate struct {
	EncounterMethod NamedAPIResource          `json:"encounter_method"`
	VersionDetails  []EncounterVersionDetails `json:"version_details"`
}

type LocationArea struct {
	ID                   int                   `json:"id"`
	Name                 string                `json:"name"`
	GameIndex            int                   `json:"game_index"`
	EncounterMethodRates []EncounterMethodRate `json:"encounter_method_rates"`
	Location             NamedAPIResource      `json:"location"`
	Names                []Name                `json:"names"`
	PokemonEncounters    []PokemonEncounter    `json:"pokemon_encounters"`
}

type APIResponse struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(config *config) error {
	res, err := http.Get(config.Next)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	var apiResponse APIResponse
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&apiResponse)
	if err != nil {
		return err
	}
	config.Next = apiResponse.Next
	if apiResponse.Previous != nil {
		config.Previous = apiResponse.Previous
	}
	fmt.Printf("Next: %v \n", config.Next)
	fmt.Printf("Previous: %v\n", config.Previous)
	for _, area := range apiResponse.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func commandMapBack(config *config) error {
	if config.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	res, err := http.Get(*config.Previous)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	var apiResponse APIResponse
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&apiResponse)
	if err != nil {
		return err
	}
	config.Next = apiResponse.Next
	if apiResponse.Previous != nil {
		config.Previous = apiResponse.Previous
	} else {
		config.Previous = nil
	}
	for _, area := range apiResponse.Results {
		fmt.Println(area.Name)
	}
	return nil
}

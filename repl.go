package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(config *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}
		commandName := words[0]
		param := &param{params: []string{}}
		if len(words) > 1 {
			param.params = append(param.params, words[1])
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(config, param)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, *param) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "List the next 20 location areas",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "map back",
			description: "List the previous 20 location areas",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore area",
			description: "List all pokemon found in the area",
			callback:    commandExplore,
		},
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(config *config, args []string) error
}

var commands map[string]cliCommand

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
			description: "Displays the map and then the next map",
			callback:    command_map,
		},
		"mapb": {
			name:        "mapb",
			description: "goes to the previous map",
			callback:    command_mapb,
		},
		"explore": {
			name:        "explore",
			description: "Explore a map",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Try and catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a caught pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "look at the pokemon you caught",
			callback:    commandPokedex,
		},
	}
}

func cleanInput(text string) []string {
	listString := strings.Fields(strings.ToLower(strings.Trim(text, " ")))
	return listString
}

func startREPL(config *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		input := scanner.Text()

		clean := cleanInput(input)
		commands = getCommands()

		value, ok := commands[clean[0]]
		if !ok {
			fmt.Print("Unknown command\n")
		} else {
			value.callback(config, clean)
		}
	}
}

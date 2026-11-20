package main

import "fmt"

func commandHelp(cfg *config , _args []string) error {
	fmt.Print("Welcome to the Pokedex!\nUsage: \n \n")
	for _, value := range cfg.commands {
		fmt.Printf("%s: %s\n", value.name, value.description)
	}
	return nil
}

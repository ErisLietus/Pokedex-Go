package main

import "fmt"

func commandPokedex(cfg *config, args []string) error {

	if len(cfg.pokemon) == 0 {
		fmt.Println("Catch some pokemon first")
	}
	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.pokemon {
		fmt.Printf("- %s\n", pokemon.Name)
	}
	return nil

}

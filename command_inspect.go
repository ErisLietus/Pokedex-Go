package main

import "fmt"

func commandInspect(cfg *config, args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("Unknown Pokemon")
	}
	pokemon := args[1]

	if _, ok := cfg.pokemon[pokemon]; !ok {
		fmt.Println("You have not caught this pokemon")
		return fmt.Errorf("you have not caught this pokemon")
	}
	caughtPokemon := cfg.pokemon[pokemon]

	fmt.Printf("Name: %s\n", caughtPokemon.Name)
	fmt.Printf("Height: %d\n", caughtPokemon.Height)
	fmt.Printf("Weight: %d\n", caughtPokemon.Weight)
	fmt.Printf("Stats:\n")
	for _ ,stats := range caughtPokemon.Stats {
		fmt.Printf("-%s: %d\n", stats.Stat.Name, stats.BaseStat)
	}
	fmt.Printf("Types:")
	for _, pokemonType := range caughtPokemon.Types{
		fmt.Printf("-%s\n", pokemonType.Type.Name)
	}
	return nil
	
}

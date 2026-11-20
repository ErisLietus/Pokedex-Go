package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type pokemon struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func commandExplore(cfg *config, args []string) error {
	var encouters pokemon
	if len(args) != 2 {
		return fmt.Errorf("Unknown Area")
	}
	areaName := args[1]
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", areaName)

	data, found := cfg.cache.Get(url)
	if !found {
		res, err := http.Get(url)

		if err != nil {
			return err
		}
		if res.Status != "200 OK" {
			return fmt.Errorf("Not found")
		}

		defer res.Body.Close()
		data, err = io.ReadAll(res.Body)

		if err != nil {
			return err
		}
		cfg.cache.Add(url, data)
	}
	if err := json.Unmarshal(data, &encouters); err != nil {
		return err
	}

	fmt.Printf("Exploring %s\n", areaName)
	fmt.Println("Found Pokemon:")
	for _, mon := range encouters.PokemonEncounters {
		fmt.Printf("- %s\n" , mon.Pokemon.Name)
	}

	return nil

}

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)


func commandCatch(cfg *config, args []string) error {
	var poke Pokemon
	if len(args) != 2 {
		return fmt.Errorf("Unknown Pokemon")
	}
	pokemon := args[1]
	url := fmt.Sprintf("%s/%s", cfg.pokemonApi, pokemon)

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

	if err := json.Unmarshal(data, &poke); err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)
	difficulty := 50 
	catchRate := rand.Intn(poke.BaseExperience)

	if catchRate <= difficulty {
		fmt.Printf("%s was caught!\n" ,poke.Name)
		cfg.pokemon[poke.Name] = poke
	}else{
		fmt.Printf("%s escaped!\n", poke.Name)
	}
	return nil

}

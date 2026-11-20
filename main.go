package main

import (
	"time"

	pokecache "github.com/ErisLietus/pokedex_go/internal"
)

type config struct {
	commands   map[string]cliCommand
	pokeApi    string
	pokemonApi string
	next       string
	previous   string
	cache      *pokecache.Cache
	pokemon    map[string]Pokemon
}

var cfg = config{
	commands:   getCommands(),
	pokeApi:    "https://pokeapi.co/api/v2/location-area",
	pokemonApi: "https://pokeapi.co/api/v2/pokemon",
	next:       "https://pokeapi.co/api/v2/location-area",
	previous:   "",
	cache:      pokecache.NewCache(5 * time.Minute),
	pokemon:    make(map[string]Pokemon),
}

func main() {
	startREPL(&cfg)
}

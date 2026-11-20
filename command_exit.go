package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config, _args []string) error {
	fmt.Print("\nClosing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}
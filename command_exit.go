package main

import (
	"fmt"
	"os"
	"pokedexcli/internal/pokeapi"
)

func commandExit(_config *pokeapi.Config, _name string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")

	os.Exit(0)
	return nil
}

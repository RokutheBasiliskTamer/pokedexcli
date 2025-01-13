package main

import (
	"fmt"
	"os"

	"github.com/RokutheBasiliskTamer/pokedexcli/internal/pokeapi"
)

func commandExit(_cfg *pokeapi.Config, _args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")

	os.Exit(0)
	return nil
}

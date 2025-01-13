package main

import (
	"fmt"

	"github.com/RokutheBasiliskTamer/pokedexcli/internal/pokeapi"
)

func commandMapB(cfg *pokeapi.Config, _args ...string) error {
	if cfg.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	res, err := cfg.Client.GetLocationArea(cfg.Previous)
	if err != nil {
		return fmt.Errorf("error grabbing map: %w", err)
	}

	var page pokeapi.PaginationResponse
	if err = page.UnByteify(res); err != nil {
		return fmt.Errorf("error unbyteifying: %w", err)
	}

	for _, result := range page.Results {

		fmt.Printf("%s", result.Name)
		fmt.Println()
	}

	cfg.Next = page.Next
	cfg.Previous = page.Previous
	return nil
}

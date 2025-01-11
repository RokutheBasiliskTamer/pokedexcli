package main

import (
	"fmt"
	"pokedexcli/internal/pokeapi"
)

func commandMap(cfg *pokeapi.Config, _args ...string) error {
	res, err := cfg.Client.GetLocationArea(cfg.Next)
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

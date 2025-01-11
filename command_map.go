package main

import (
	"fmt"
	"pokedexcli/internal/pokeapi"
)

func commandMap(config *pokeapi.Config, _name string) error {
	res, err := config.Client.GetLocationArea(config.Next)
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
	config.Next = page.Next
	config.Previous = page.Previous
	return nil
}

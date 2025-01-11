package main

import (
	"fmt"
	"pokedexcli/internal/pokeapi"
)

func commandMap(config *pokeapi.Config, _name string) error {
	res, err := config.Client.GetLocationAreasPage(config.Next)
	if err != nil {
		return fmt.Errorf("error grabbing map: %w", err)
	}

	for _, result := range res.Results {
		fmt.Printf("%s", result.Name)
		fmt.Println()
	}
	config.Next = res.Next
	config.Previous = res.Previous
	return nil
}

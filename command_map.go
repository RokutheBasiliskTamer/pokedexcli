package main

import (
	"fmt"
	"pokedexcli/internal/pokeapi"
)

func commandMap(config *pokeapi.Config) error {
	res, err := config.Client.GetLocationAreas(config.Next)
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

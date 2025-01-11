package main

import (
	"fmt"
	"pokedexcli/internal/pokeapi"
)

func commandMapB(config *pokeapi.Config) error {
	if config.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	res, err := config.Client.GetLocationAreas(config.Previous)
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

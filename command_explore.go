package main

import (
	"fmt"
	"pokedexcli/internal/pokeapi"
)

func commandExplore(cfg *pokeapi.Config, args ...string) error {
	if len(args) < 1 {
		fmt.Println("Location name required!")
		return nil
	}
	name := args[0]
	fmt.Printf("Exploring %s...\n", name)
	byteData, err := cfg.Client.GetLocationArea(&name)
	if err != nil {
		return fmt.Errorf("error exploring: %w", err)
	}
	var location pokeapi.LocationArea
	if err = location.UnByteify(byteData); err != nil {
		return fmt.Errorf("error unbyteifying: %w", err)
	}
	for _, pokemon_encounter := range location.Pokemon_encounters {
		fmt.Printf(" - %s\n", pokemon_encounter.Pokemon.Name)
	}
	return nil

}

package main

import (
	"fmt"
	"pokedexcli/internal/pokeapi"
)

func commandExplore(config *pokeapi.Config, name string) error {
	fmt.Printf("Exploring %s...\n", name)
	byteData, err := config.Client.GetLocationArea(&name)
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

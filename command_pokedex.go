package main

import (
	"fmt"
	"pokedexcli/internal/pokeapi"
)

func commandPokedex(cfg *pokeapi.Config, _args ...string) error {
	if len(cfg.CaughtPokemon) < 1 {
		fmt.Println("You've yet to catch any pokemon! Use the catch command to catch some!")
		return nil
	}
	println("Your Pokedex:")
	for key := range cfg.CaughtPokemon {
		fmt.Printf(" - %s\n", key)
	}
	return nil
}

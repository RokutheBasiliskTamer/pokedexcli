package main

import (
	"fmt"
	"pokedexcli/internal/pokeapi"
)

func commandHelp(_cfg *pokeapi.Config, _args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

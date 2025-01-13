package main

import (
	"fmt"

	"github.com/RokutheBasiliskTamer/pokedexcli/internal/pokeapi"
)

func commandInspect(cfg *pokeapi.Config, args ...string) error {
	if len(args) < 1 {
		fmt.Println("Pokemon name required!")
		return nil
	}
	name := args[0]

	pokeData, ok := cfg.CaughtPokemon[name]
	if !ok {
		fmt.Println("You have not caught that pokemon!")
		return nil
	}
	stats := pokeData.Stats
	types := pokeData.Types
	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Height: %d\n", pokeData.Height)
	fmt.Printf("Weight: %d\n", pokeData.Weight)
	fmt.Println("Stats:")
	for _, stat := range stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.Base_stat)
	}
	fmt.Println("Types:")
	for _, typ := range types {
		fmt.Printf("  -%s\n", typ.Type.Name)
	}

	return nil
}

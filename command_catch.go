package main

import (
	"fmt"
	"math/rand"
	"pokedexcli/internal/pokeapi"
)

func commandCatch(cfg *pokeapi.Config, args ...string) error {
	if len(args) < 1 {
		fmt.Println("Pokemon name required!")
		return nil
	}
	name := args[0]

	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	if _, ok := cfg.CaughtPokemon[name]; ok {
		fmt.Printf("%s was caught!\n", name)
		return nil
	}

	byteData, err := cfg.Client.GetPokemon(&name)

	if err != nil {
		return fmt.Errorf("error fetching pokemon: %w", err)
	}
	var pokemon pokeapi.Pokemon
	if err = pokemon.UnByteify(byteData); err != nil {
		return fmt.Errorf("error unbyteifying pokemon: %w", err)
	}

	exp := pokemon.Base_experience
	if catch(exp) {
		cfg.CaughtPokemon[name] = pokemon
		fmt.Printf("%s was caught!\n", name)
		return nil
	}
	fmt.Printf("%s got away!\n", name)
	return nil
}

func catch(baseExp int) bool {
	catch_roll := rand.Intn(baseExp)
	//fmt.Printf("%d", catch_roll)
	return catch_roll > 40
}

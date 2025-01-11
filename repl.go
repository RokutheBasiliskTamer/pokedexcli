package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedexcli/internal/pokeapi"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(config *pokeapi.Config) error
}

func cleanInput(text string) []string {
	var cleanedInput []string
	words := strings.Fields(strings.ToLower(text))

	cleanedInput = append(cleanedInput, words...)

	return cleanedInput
}

func getCommands() (commands map[string]cliCommand) {
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays next page of 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous page of 20 locations",
			callback:    commandMapB,
		},
	}
	return commands
}

func startRepl() {

	var config pokeapi.Config
	config.Client = pokeapi.NewClient()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		input := scanner.Text()
		cleanedInput := cleanInput(input)
		if len(cleanedInput) == 0 {
			continue
		}
		cmd := cleanedInput[0]
		commands := getCommands()
		command, ok := commands[cmd]
		if ok {
			if err := command.callback(&config); err != nil {
				fmt.Println()
				fmt.Printf("%v", err)
			}
			continue
		} else {
			fmt.Println("Invalid Command!")
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}

	}
}

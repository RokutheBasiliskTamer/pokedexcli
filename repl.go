package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedexcli/internal/pokeapi"
	"strings"
	"time"
)

type cliCommand struct {
	name        string
	description string
	callback    func(config *pokeapi.Config) error
}

func cleanInput(text string) []string {

	return strings.Fields(strings.ToLower(text))

}

func getCommands() (commands map[string]cliCommand) {
	return map[string]cliCommand{
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

}

func startRepl() {
	// initialize the empty config struct and set the client
	var config pokeapi.Config
	config.Client = pokeapi.NewClient(5*time.Second, time.Minute*5)

	// initialize the scanner and enter and endless loop
	scanner := bufio.NewScanner(os.Stdin)
	for {
		//print the repl name and scan for input
		//clean the input when recieved and make sure it wasnt empty
		fmt.Print("Pokedex >")
		scanner.Scan()
		input := scanner.Text()
		cleanedInput := cleanInput(input)
		if len(cleanedInput) == 0 {
			continue
		}

		//set the command to the first word of the input and check if its a valid command
		cmd := cleanedInput[0]
		commands := getCommands()
		command, ok := commands[cmd]
		if ok {
			//if its valid call the commands callback
			fmt.Println()
			if err := command.callback(&config); err != nil {
				fmt.Println()
				fmt.Printf("%v", err)
			}
			continue
		} else {
			//tell the user it was invalid and start loop again
			fmt.Println("Invalid Command!")
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}

	}
}

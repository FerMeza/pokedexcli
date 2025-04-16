package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/FerMeza/pokedexcli/internal/pokeapi"
)

var commands map[string]cliCommand

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		moreTokens := scanner.Scan()
		if !moreTokens {
			err := scanner.Err()
			if err != nil {
				fmt.Printf("There was an error on input: %v", err)
				return
			}
			fmt.Printf("Goodbye!")
			return
		}
		userInput := scanner.Text()
		cleanInput := cleanInput(userInput)
		if len(cleanInput) == 0 {
			continue
		}

		if len(cleanInput) >= 0 {
			command, ok := commands[cleanInput[0]]
			if !ok {
				fmt.Println("Unknown command")
				continue
			} else {
				err := command.callback(cfg)
				if err != nil {
					fmt.Printf("error executing command: %v", err)
				}
			}
		}
	}
}

type config struct {
	pokeAPIClient        pokeapi.Client
	nextLocationsURL     *string
	previousLocationsURL *string
}

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	words := strings.Fields(lowerText)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func init() {
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
			description: "Displays names of next 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays names of previous 20 locations areas",
			callback:    commandMapB,
		},
	}
}

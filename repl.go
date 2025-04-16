package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var commands = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
}

func startRepl() {
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
				err := command.callback()
				if err != nil {
					fmt.Printf("error executing command: %v", err)
				}
			}
		}
	}
}

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	words := strings.Fields(lowerText)
	return words
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")
	for _, val := range commands {
		fmt.Printf("%s: %s\n", val.name, val.description)
	}
	return nil
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func init() {
	commands["help"] = struct {
		name        string
		description string
		callback    func() error
	}{
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	}
}

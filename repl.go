package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

var commands map[string]cliCommand

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	config := config{}
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
				err := command.callback(&config)
				if err != nil {
					fmt.Printf("error executing command: %v", err)
				}
			}
		}
	}
}

type config struct {
	Next     *string
	Previous *string
}

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	words := strings.Fields(lowerText)
	return words
}

func commandExit(config *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")
	for _, val := range commands {
		fmt.Printf("%s: %s\n", val.name, val.description)
	}
	return nil
}

type pokeArea struct {
	Name string `json:"name"`
}

type responsePokeArea struct {
	Next     *string    `json:"next"`
	Previous *string    `json:"previous"`
	Results  []pokeArea `json:"results"`
}

func commandMap(config *config) error {
	if config == nil {
		panic("Invalid state on command map")
	}
	url := config.Next
	if url == nil {
		defURL := "https://pokeapi.co/api/v2/location-area/"
		url = &defURL
	}

	req, err := http.NewRequest(http.MethodGet, *url, nil)

	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	client := http.DefaultClient
	res, err := client.Do(req)

	if err != nil {
		return fmt.Errorf("error sending request: %v", err)
	}

	defer res.Body.Close()

	var pokeAreas responsePokeArea

	err = json.NewDecoder(res.Body).Decode(&pokeAreas)
	if err != nil {
		return fmt.Errorf("error decoding response: %v", err)
	}

	for _, area := range pokeAreas.Results {
		fmt.Println(area.Name)
	}

	config.Next = pokeAreas.Next
	config.Previous = pokeAreas.Previous

	return nil
}

func commandMapB(config *config) error {
	if config == nil {
		panic("Invalid state on command mapb")
	}
	url := config.Previous
	if url == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	req, err := http.NewRequest(http.MethodGet, *url, nil)

	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	client := http.DefaultClient
	res, err := client.Do(req)

	if err != nil {
		return fmt.Errorf("error sending request: %v", err)
	}

	defer res.Body.Close()

	var pokeAreas responsePokeArea

	err = json.NewDecoder(res.Body).Decode(&pokeAreas)
	if err != nil {
		return fmt.Errorf("error decoding response: %v", err)
	}

	for _, area := range pokeAreas.Results {
		fmt.Println(area.Name)
	}

	config.Next = pokeAreas.Next
	config.Previous = pokeAreas.Previous

	return nil
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

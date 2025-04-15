package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
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
			fmt.Println("There is no command!")
		} else {
			fmt.Printf("Your command was: %s\n", cleanInput[0])
		}
	}
}

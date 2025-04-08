package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commands = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"help": {
		name:        "help",
		description: "Exit the Pokedex",
		callback:    commandHelp,
	},
}

func main() {
	scanner := bufio.NewScanner(os.Stdin) // Create a scanner for standard input
	for {
		fmt.Print("Pokedex > ")
		scanned := scanner.Scan() // Wait for input
		if !scanned {             // Break if there's an error or EOF
			break
		}
		text := scanner.Text() // Get the input as a string
		cleanedText := cleanInput((text))

		// Lookup the command in the registry
		if len(cleanedText) > 0 {
			command, exists := commands[cleanedText[0]]
			if exists {
				err := command.callback()
				if err != nil {
					fmt.Println("Error:", err)
				}
			} else {
				fmt.Println("Unknown command")
			}
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(strings.TrimSpace(text)))
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	return nil
}

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func main() {
	cfg := &config{}
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
			command, exists := getCommands()[cleanedText[0]]
			if exists {
				err := command.callback(cfg)
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

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

type config struct {
	NextURL     string
	PreviousURL string
}

type LocationAreaResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMapf(cfg *config) error {
	// 1. Determine which URL to use (first request or using the NextURL)
	url := "https://pokeapi.co/api/v2/location-area"
	if cfg.NextURL != "" {
		url = cfg.NextURL
	}

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("An error occurred trying to find the locations")
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body")
		return err
	}

	var locationResp LocationAreaResponse
	err = json.Unmarshal(body, &locationResp)
	if err != nil {
		fmt.Println("Error parsing location data")
		return err
	}

	// Print location names
	for _, location := range locationResp.Results {
		fmt.Println(location.Name)
	}

	// Update config for pagination
	cfg.NextURL = locationResp.Next
	cfg.PreviousURL = locationResp.Previous

	return nil
}

func commandMapb(cfg *config) error {
	// Check if we're on the first page
	if cfg.PreviousURL == "" {
		fmt.Println("You're on the first page")
		return nil
	}

	// Make the HTTP request using the previous URL
	resp, err := http.Get(cfg.PreviousURL)
	if err != nil {
		fmt.Println("An error occurred trying to find the previous locations")
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body")
		return err
	}

	var locationResp LocationAreaResponse
	err = json.Unmarshal(body, &locationResp)
	if err != nil {
		fmt.Println("Error parsing location data")
		return err
	}

	// Print location names
	for _, location := range locationResp.Results {
		fmt.Println(location.Name)
	}

	// Update config for pagination
	cfg.NextURL = locationResp.Next
	cfg.PreviousURL = locationResp.Previous

	return nil
}

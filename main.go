package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
		fmt.Println("Your command was:", cleanedText[0]) // Print back the input
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(strings.TrimSpace(text)))
}

package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	// Inside your command handler for "catch"
	pokemonName := args[0] // Assuming args is the command arguments

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	// Get pokemon data from API
	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil
	}

	// Now use the BaseExperience to calculate catch chance
	catchChance := 100 - (pokemon.BaseExperience / 4)
	if catchChance < 5 {
		catchChance = 5 // Minimum catch chance
	}

	// Use rand to determine if caught
	randomValue := rand.Intn(100)
	caught := randomValue < catchChance

	if caught {
		fmt.Printf("%s was caught!\n", pokemonName)
		// Add to your map of caught Pokemon
		//caughtPokemon[pokemonName] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}

	return nil
}

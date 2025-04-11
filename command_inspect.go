package main

import (
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) == 0 {
		fmt.Println("Please enter a pokemon name!")
		return nil
	}
	pokemon := args[0]
	if len(pokemon) == 0 {
		fmt.Println("Please enter a valid pokemon!")
	} else {
		// Check if the Pokémon exists in the map
		pokemonData, exists := cfg.caughtPokemon[pokemon]

		if exists {
			// The Pokémon is in the map
			fmt.Printf("Name: %s\n", pokemonData.Name)
			fmt.Printf("Height: %d\n", pokemonData.Height)
			fmt.Printf("Weight: %d\n", pokemonData.Weight)
			fmt.Println("Stats:")
			for _, data := range pokemonData.Stats {
				fmt.Printf("  - %s: %d\n", data.Stat.Name, data.BaseStat)
			}
			fmt.Println("Types:")
			for _, data := range pokemonData.Types {
				fmt.Printf("  - %s\n", data.Type.Name)
			}
		} else {
			// The Pokémon is not in the map
			fmt.Println("You have not caught that pokemon")
		}
	}
	return nil
}

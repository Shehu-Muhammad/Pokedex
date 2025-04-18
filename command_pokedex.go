package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.caughtPokemon) == 0 {
		fmt.Println("You haven't caught any pokemon, yet")
	} else {
		fmt.Println("Your Pokedex:")
		for pokemon := range cfg.caughtPokemon {
			fmt.Printf("- %s\n", pokemon)
		}
	}

	return nil
}

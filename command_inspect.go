package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a Pokemon name")
	}

	name := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}
	fmt.Printf("Inspecting %s...\n", pokemon.Name)

	// Check if the user has caught this Pokemon
	caughtPokemon, exists := cfg.Pokedex[pokemon.Name]
	if !exists {
		// User has not caught this Pokemon
		return fmt.Errorf("you have not caught a %s yet. No information available.", pokemon.Name)
	}

	// If the user has caught the Pokemon, print various details
	fmt.Printf("Name: %s\n", caughtPokemon.Name)
	fmt.Printf("Height: %d decimetres\n", caughtPokemon.Height)
	fmt.Printf("Weight: %d hectograms\n", caughtPokemon.Weight)
	
	// Print stats
	fmt.Println("Stats:")
	for _, stat := range caughtPokemon.Stats {
		fmt.Printf("- %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	
	// Print abilities
	fmt.Println("Abilities:")
	for _, ability := range caughtPokemon.Abilities {
		fmt.Printf("- %s (Hidden: %t)\n", ability.Ability.Name, ability.IsHidden)
	}
	
	// Print species
	fmt.Printf("Species: %s\n", caughtPokemon.Species.Name)

	// Print types
	fmt.Printf("Types:\n")
	for _, typeInfo := range caughtPokemon.Types {
		fmt.Printf("- %s\n", typeInfo.Type.Name)
	}

	return nil
}
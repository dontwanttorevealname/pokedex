package main

import (
	"errors"
	"fmt"
	"math/rand"
	"math"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}
	fmt.Printf("Found %s...\n", pokemon.Name)

	// Check if there are any types to avoid panic
	if len(pokemon.Types) > 0 {
		// Iterate over each type
		fmt.Printf("%s is the following types:\n", pokemon.Name)
		for _, typeInfo := range pokemon.Types {
			fmt.Printf("- %s type\n", typeInfo.Type.Name)
		}
	} else {
		fmt.Printf("%s has no types listed.\n", pokemon.Name)
	}

	// Calculate catch probability based on base experience
	baseXP := pokemon.BaseExperience
	
	// Adjust the catch chance calculation to make it harder to catch higher base XP Pokémon
	catchChance := 1.0 - math.Pow(float64(baseXP)/1500.0, 2) // Squaring makes high values drop catch chances more significantly
	
	// Ensure catchChance does not go below 0 or above 1
	if catchChance < 0 {
		catchChance = 0
	}
	if catchChance > 1 {
		catchChance = 1
	}

	// Generate random float between 0 and 1
	if rand.Float64() < catchChance {
		fmt.Printf("You've successfully caught %s!\n", pokemon.Name)

		// Add to Pokedex
		cfg.Pokedex[pokemon.Name] = pokemon // Assuming `Pokedex` is a map[string]*Pokemon
		fmt.Printf("%s has been added to your Pokedex!\n", pokemon.Name)
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}
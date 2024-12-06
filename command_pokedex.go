package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
    if len(cfg.Pokedex) == 0 {
        fmt.Println("Your Pokedex is empty.")
        return nil // Return nil to indicate no error occurred
    }

    fmt.Println("Pokemon in your Pokedex:")
    for _, pokemon := range cfg.Pokedex {
        fmt.Println(pokemon.Name)
    }

    return nil // Return nil to indicate completion without an error
}
package main

import (
	"time"

	"github.com/dontwanttorevealname/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		pokeapiClient: pokeClient,
            Pokedex:       make(map[string]pokeapi.Pokemon),
	}

	startRepl(cfg)
}

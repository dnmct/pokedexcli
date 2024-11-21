package main

import (
	"time"

	"github.com/dnmct/pokedexcli/internal/config"
	"github.com/dnmct/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config.Config{
		PokeapiClient: pokeClient,
		CaughtPokemon: map[string]pokeapi.Pokemon{},
	}

	startRepl(cfg)
}

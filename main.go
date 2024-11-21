package main

import (
	"time"
)

func main() {
	pokeClient := NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
		caughtPokemon: map[string]Pokemon{},
	}

	startRepl(cfg)
}

package commands

import (
	"fmt"

	"github.com/dnmct/pokedexcli/internal/config"
)

func commandPokedex(cfg *config.Config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for _, p := range cfg.CaughtPokemon {
		fmt.Printf(" - %s\n", p.Name)
	}
	return nil
}

package commands

import (
	"os"

	"github.com/dnmct/pokedexcli/internal/config"
)

func commandExit(cfg *config.Config, args ...string) error {
	os.Exit(0)
	return nil
}

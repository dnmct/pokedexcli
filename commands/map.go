package commands

import (
	"errors"
	"fmt"

	"github.com/dnmct/pokedexcli/internal/config"
)

func commandMapf(cfg *config.Config, args ...string) error {
	locationsResp, err := cfg.PokeapiClient.ListLocations(cfg.NextLocationsURL)
	if err != nil {
		return err
	}

	cfg.NextLocationsURL = locationsResp.Next
	cfg.PrevLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config.Config, args ...string) error {
	if cfg.PrevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locationResp, err := cfg.PokeapiClient.ListLocations(cfg.PrevLocationsURL)
	if err != nil {
		return err
	}

	cfg.NextLocationsURL = locationResp.Next
	cfg.PrevLocationsURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

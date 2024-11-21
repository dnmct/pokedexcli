package commands

import "github.com/dnmct/pokedexcli/internal/config"

type CliCommand struct {
	name        string
	description string
	Callback    func(*config.Config, ...string) error
}

func GetCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			Callback:    commandHelp,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Attempt to catch a pokemon",
			Callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "View details about a caught Pokemon",
			Callback:    commandInspect,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Explore a location",
			Callback:    commandExplore,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			Callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			Callback:    commandMapb,
		},
		"pokedex": {
			name:        "pokedex",
			description: "See all the pokemon you've caught",
			Callback:    commandPokedex,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			Callback:    commandExit,
		},
	}
}

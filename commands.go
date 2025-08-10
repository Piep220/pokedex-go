package main

import (
	"fmt"
	"os"
	"pokedex-go/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
} 

type config struct {
	pokeApiClient  pokeapi.Client
	areaMapNext    string
	areaMapPrevous string
	pokedex		   pokedex
}

func commandExit(cfg *config, args ...string) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	cfg.pokeApiClient.Stop()
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, args ...string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
            description: "Exit the Pokedex",
            callback:    commandExit,
        },
		"map": {
			name:        "map",
			description: "Displays up to 20 of the next map locations. Call again for more.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays up to 20 of the previous map locations. Call again for more.",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Displays a list of the pokemon in a given region",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a pokemon, by name.",
			callback:    commandCatch,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Shows the contence of the pokedex.",
			callback:    commandPokedex,
		},
	}
}
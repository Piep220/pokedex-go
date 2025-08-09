package main

import (
	"fmt"
	"os"
	"pokedex-go/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
} 

type config struct{
	pokeApiClient  pokeapi.Client
	areaMapNext    string
	areaMapPrevous string
}

func commandExit(cfg *config) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	cfg.pokeApiClient.Stop()
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap(cfg *config) error {
	next_url := pokeapi.BaseURL + "/location-area?offset=0&limit=20"
	switch cfg.areaMapNext {
	case "default":
		//do nothing
	case "":
		fmt.Println("End of map list, try mapb for reverse.")
		return nil
	default:
		next_url = cfg.areaMapNext
	}

	areaData, err := cfg.pokeApiClient.GetAreaPage(next_url)
  		if err != nil {
       	fmt.Println("Error fetching area data:", err)
       	return err
   	}
	
	cfg.areaMapNext = areaData.Next
	cfg.areaMapPrevous = areaData.Previous

	for i := range areaData.Results {
		fmt.Println(areaData.Results[i].Name)
	}

	return nil
}

func commandMapb(cfg *config) error {
	previous_url := ""
	switch cfg.areaMapPrevous {
	case "default":
		fmt.Println("No previous page, try map for forward list.")
		return nil
	case "":
		fmt.Println("No previous page, try map for forward list.")
		return nil
	default:
		previous_url = cfg.areaMapPrevous
	}

	areaData, err := cfg.pokeApiClient.GetAreaPage(previous_url)
  		if err != nil {
       	fmt.Println("Error fetching area data:", err)
       	return err
   	}
	
	cfg.areaMapNext = areaData.Next
	cfg.areaMapPrevous = areaData.Previous

	for i := range areaData.Results {
		fmt.Println(areaData.Results[i].Name)
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
	}
}
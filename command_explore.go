package main

import (
	"fmt"
)


func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("you must provide a location name")
	}

	areaName := args[0]
	pokemonNames, err := cfg.pokeApiClient.GetPokemonInArea(areaName)
	if err != nil {
		fmt.Printf("Area name invalid. Use map to get full area names. eg pastoria-city-area\n")
		return fmt.Errorf("error fetching pokemon from %s, returned error: %s", areaName, err)
	}

	fmt.Println("Found Pokemon:")
	for _, pokemon := range pokemonNames {
		fmt.Printf(" - %s\n", pokemon)
	}

	return nil
}
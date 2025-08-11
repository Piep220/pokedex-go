package main

import (
	"fmt"
)

//Print hgiht, weight, stats and types of a pokemon in the pokedex.
func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("you must provide a pokemon name")
	}
	pokemonName := args[0]

	//Check if valid name
	ok, err := cfg.pokeApiClient.VerifyPokemonName(pokemonName)
	if !ok {
		fmt.Println("Pokemon, not found. Check name.")
		return nil
	}
	if err != nil {
		fmt.Println(err)
		return err
	}

	//Print details if in pokedex
	for _, pokemon := range cfg.pokedex.pokemon{
		if pokemon.name == pokemonName {
			fmt.Printf("Name: %s\n", pokemonName)
			fmt.Printf("Height: %d\n", pokemon.details.Height)
			fmt.Printf("Weight: %d\n", pokemon.details.Weight)
			
			fmt.Println("Stats:")
			for _, stat := range pokemon.details.Stats {
				fmt.Printf("  -%s: %d\n",stat.Stat.Name,stat.BaseStat)
			}
			
			fmt.Println("Types:")
			for _, types := range pokemon.details.Types {
				fmt.Printf("  -%s\n", types.Type.Name)
			}
			return nil
		}
	}

	fmt.Println("You still need to catch that Pokemon!")
	return nil
}
package main

import (
	"fmt"
	"pokedex-go/internal/pokeapi"
)

type pokedex struct {
	pokemon 	[]pokemon
}

type pokemon struct {
	name 	string
	details pokeapi.PokemonDetails
}

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your PokeDex contains:")

	fmt.Println("Caught Pokemon:")
	for _, pokemon := range cfg.pokedex.pokemon {
		fmt.Printf(" - %s\n", pokemon.name)
	}
	return nil
}
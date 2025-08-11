package main

import (
	"fmt"
	"math/rand"
	"math"
)


func expToProb(exp float64, mu, sigma float64) float64 {
    // Logistic curve
    raw := 1 / (1 + math.Exp(-(exp-mu)/sigma))
    if raw < 0 { return 0 }
    if raw > 1 { return 1 }
    return raw
}

func tryCatch(baseExp float64, ballMod, statusMod float64) bool {
	var mu, sigma float64 = 125.0, 60.0
	mu *= statusMod
	sigma *= ballMod
    p := expToProb(baseExp, mu, sigma)

    // Clamp to [0,1]
    if p < 0 { p = 0 }
    if p > 1 { p = 1 }

    return rand.Float64() >= p
}

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("you must provide a pokemon name")
	}
	pokemonName := args[0]

	ok, err := cfg.pokeApiClient.VerifyPokemonName(pokemonName)
	if !ok {
		fmt.Println("Pokemon, not found. Check name.")
		return nil
	}
	if err != nil {
		fmt.Println(err)
		return err
	}

	
	fmt.Printf("Throwing a Pokeball at %s... ", pokemonName)

	//Get pokemon details
	pokemonDetails, err := cfg.pokeApiClient.GetPokemonDetails(pokemonName)
	if err != nil {
		fmt.Println(err)
		return err
	}

	//Catch fail
	if !tryCatch(float64(pokemonDetails.BaseExperience), 1.0, 1.0){
		fmt.Printf("%s escaped!\n", pokemonName)
		return nil
	}

	//Catch sucess
	fmt.Printf("%s was caught!\n", pokemonName)

	new_pokemon := pokemon{
		name: pokemonName,
		details: *pokemonDetails,
	}
	cfg.pokedex.pokemon = append(cfg.pokedex.pokemon, new_pokemon)

	return nil
}
package pokeapi

import (
	"fmt"
)

//get page of pokemon data and return *PokemonDetails struct
func (c *Client)GetPokemonDetails(name string) (*PokemonDetails, error) {
	url := BaseURL + "/pokemon/" + name
	details, err := getJSON[PokemonDetails](c, url)
	if err != nil {
        return nil, fmt.Errorf("error getting JSON, GetPokemonDetails %s: %w", url, err)
    }

	return details, nil
}

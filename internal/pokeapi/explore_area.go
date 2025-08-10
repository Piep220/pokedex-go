package pokeapi

import (
	"fmt"
	"strings"
)

//get page of location-area details and return *LocationAreaDetails struct
func (c *Client)GetAreaDetails(url string) (*LocationAreaDetails, error) {
	if !strings.HasPrefix(url, BaseURL + "/location-area/") {
		return nil, fmt.Errorf("location area URL, %s is not a valid area url", url)
	}

	areaDetails, err := getJSON[LocationAreaDetails](c, url)
	if err != nil {
        return nil, fmt.Errorf("error getting JSON %s: %w", url, err)
    }

	return areaDetails, nil
}

func (c Client)GetPokemonInArea(areaName string) ([]string, error) {
	url := BaseURL + "/location-area/" + areaName
	areaDetails, err := c.GetAreaDetails(url)
	if err != nil {
		return nil, fmt.Errorf("error getting area details: %w", err)
	}

	pokemon := []string{}
	for _, encounter := range areaDetails.PokemonEncounters {
		pokemon = append(pokemon, encounter.Pokemon.Name)
	}
	return pokemon, nil
}
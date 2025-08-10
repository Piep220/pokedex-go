package pokeapi

import (
	"encoding/json"
	"fmt"
	"strconv"
	"slices"
)
	
type PokemonNamesPage struct {
	Count    int       `json:"count"`
	Next     any       `json:"next"`
	Previous any       `json:"previous"`
	Results  []Results `json:"results"`
}
type Results struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

//get page of location-area data and return *LocationAreaPage struct
func (c *Client)GetPokemonNames() ([]string, error) {

	//Check for cached version
	v, ok := c.pokeCache.Get("pokemonNames")
	if ok {
    	var vals []string
		err := json.Unmarshal(v, &vals);
    	if  err != nil {
    	    return nil, err
    	}
    	return vals, nil
    }

	//Setup
	url := BaseURL + "/pokemon"

	//Get count
	urlSingle := url + "?limit=1"
	type count struct{ 
		Count int `json:"count"` 
	}
	countResult, err := getJSON[count](c, urlSingle)
	if err != nil {
        return nil, fmt.Errorf("getJSON error in GetPokemonNames %s: %w", url, err)
    }

	//Get full list
	urlAll := url + "?limit=" + strconv.Itoa(countResult.Count)
	namesPage, err := getJSON[PokemonNamesPage](c, urlAll)
	if err != nil {
        return nil, fmt.Errorf("getJSON error in GetPokemonNames %s: %w", url, err)
    }

	//Unpack
	var namesList []string
	for _, pokemon := range namesPage.Results {
		namesList = append(namesList, pokemon.Name)
	}

	return namesList, nil
}

//Checks if name is valid, true if found, fasle not found.
func (c *Client)VerifyPokemonName(pokemonName string) (bool, error) {
	pokemonList, err := c.GetPokemonNames()
	if err != nil {
		return false, fmt.Errorf("get pokemon list error in veryifyPokemonName: %w", err)
	}

	if slices.Contains(pokemonList,pokemonName) {
		return true, nil
	}

	return false, nil
}
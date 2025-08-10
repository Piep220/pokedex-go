package main

import (
	"fmt"
	"pokedex-go/internal/pokeapi"
)

func commandMap(cfg *config, args ...string) error {
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

func commandMapb(cfg *config, args ...string) error {
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
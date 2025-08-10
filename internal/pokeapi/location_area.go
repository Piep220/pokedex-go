package pokeapi

import (
	"fmt"
	"strings"
)

type LocationAreaPage struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

//get page of location-area data and return *LocationAreaPage struct
func (c *Client)GetAreaPage(url string) (*LocationAreaPage, error) {
	if !strings.HasPrefix(url, BaseURL + "/location-area") {
		return nil, fmt.Errorf("location area URL, %s is not a valid area url", url)
	}

	areaPage, err := getJSON[LocationAreaPage](c, url)
	if err != nil {
        return nil, fmt.Errorf("cache unmarshal %s: %w", url, err)
    }

	return areaPage, nil
}

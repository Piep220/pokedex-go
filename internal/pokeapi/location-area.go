package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

func (c *Client)GetAreaPage(url string) (*LocationAreaPage, error) {
	if !strings.HasPrefix(url, BaseURL + "/location-area") {
		return nil, fmt.Errorf("location Area URL, %s is not a valid url", url)
	}

	//resp, err := http.Get(url)
	//if err != nil {
	//	return &areaPage, fmt.Errorf("error getting url: %s", err)
	//}
	//defer resp.Body.Close()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("get request error, %s", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http do error, %s", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading JSON, %s", err)
	}

	areaPage := LocationAreaPage{}
	err = json.Unmarshal(data, &areaPage)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON, %s", err)
	}

	return &areaPage, nil
}

package pokeapi

import (
	"io"
	"fmt"
	"net/http"
	"encoding/json"
	"strings"
)

// getJSON fetches JSON from url, caches the raw bytes,
// unmarshals into T and returns a pointer to T.
func getJSON[T any](c *Client, url string) (*T, error) {
	if !strings.HasPrefix(url, BaseURL) {
		return nil, fmt.Errorf("base URL, %s is not valid", url)
	}

	v, ok := c.pokeCache.Get(url);
    if ok {
        var out T
		err := json.Unmarshal(v, &out)
        if err != nil {
            return nil, fmt.Errorf("cache unmarshal %s: %w", url, err)
        }
        return &out, nil
    }

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, fmt.Errorf("new request: %w", err)
    }

    resp, err := c.httpClient.Do(req)
    if err != nil {
        return nil, fmt.Errorf("http do: %w", err)
    }
    defer resp.Body.Close()

    data, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("read body: %w", err)
    }

    c.pokeCache.Add(url, data)

	var out T
	err = json.Unmarshal(data, &out)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON, %s", err)
	}

    return &out, nil
}
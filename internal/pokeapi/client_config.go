package pokeapi

import (
	"net/http"
	"pokedex-go/internal/pokecache"
	"time"
)

const (
	BaseURL = "https://pokeapi.co/api/v2"
)


// Client -
type Client struct {
	httpClient http.Client
	pokeCache  pokecache.Cache
}

// NewClient -
func NewClient(timeout time.Duration) Client {
	c := Client{
		httpClient: http.Client{
			Timeout: timeout,
			
		},
		pokeCache: pokecache.NewCache(timeout, timeout),
	}
	return c
}

//Stop terminates background connections and goroutine reapLoop.
func (c *Client) Stop() {
	c.httpClient.CloseIdleConnections()
    c.pokeCache.Stop()
}

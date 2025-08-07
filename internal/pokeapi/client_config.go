package pokeapi

import (
	"net/http"
	"time"
)

const (
	BaseURL = "https://pokeapi.co/api/v2"
)


// Client -
type Client struct {
	httpClient http.Client
}

// NewClient -
func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
			
		},
	}
}

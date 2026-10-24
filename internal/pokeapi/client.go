package pokeapi

import (
	"net/http"
	"time"

	"github.com/MJSPinkerton/PokeDex/internal/pokecache"
)

// Client -
type Client struct {
	httpClient http.Client
	pokeCache  pokecache.Cache
}

// NewClient -
func NewClient(timeout, cacheInterval time.Duration) Client {
	pokeCache := pokecache.NewCache(cacheInterval)
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		pokeCache: pokeCache,
	}
}

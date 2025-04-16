package pokeapi

import (
	"net/http"
	"time"

	"github.com/FerMeza/pokedexcli/internal/pokecache"
)

type Client struct {
	cache      *pokecache.Cache
	httpClient http.Client
}

func NewClient(timeout time.Duration, interval time.Duration) Client {
	cache := pokecache.NewCache(interval)
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: cache,
	}
}

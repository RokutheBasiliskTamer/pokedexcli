package pokeapi

import "pokedexcli/internal/pokecache"

type PaginationResponse struct {
	Count    int
	Next     *string
	Previous *string
	Results  []struct {
		Name string
		URL  string
	}
}

type Config struct {
	Next     *string
	Previous *string
	Client   *Client
	Cache    *pokecache.Cache
}

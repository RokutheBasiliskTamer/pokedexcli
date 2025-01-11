package pokeapi

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
}

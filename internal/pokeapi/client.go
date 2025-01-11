package pokeapi

import (
	"fmt"
	"io"
	"net/http"
	"pokedexcli/internal/pokecache"
	"strings"
	"time"
)

type Endpoint interface {
	UnByteify([]byte) error
}

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
	baseURL    string
}

func NewClient(timeout, cacheInterval time.Duration) *Client {
	return &Client{
		cache: *pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
		baseURL: "https://pokeapi.co/api/v2",
	}
}

func (c *Client) makeRequest(url string) ([]byte, error) {
	//create request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	//send request
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return nil, fmt.Errorf("bad status code: %d", res.StatusCode)
	}

	//read data from request
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	//return data
	return data, nil

}

func (c *Client) getData(pageURL string) ([]byte, error) {

	var byteData []byte

	//see if weve cached that entry already
	res, ok := c.cache.Get(pageURL)
	if ok {

		byteData = res
	} else {

		//make request and get byte data
		res, err := c.makeRequest(pageURL)
		if err != nil {
			return byteData, err
		}
		byteData = res
		c.cache.Add(pageURL, byteData)
	}

	return byteData, nil
}

func (c *Client) GetLocationArea(location *string) ([]byte, error) {
	fullUrl := c.baseURL + "/location-area/"
	if location != nil {
		if strings.Contains(*location, fullUrl) {
			fullUrl = *location
		} else {
			fullUrl += *location
		}

	}

	byteData, err := c.getData(fullUrl)
	if err != nil {
		return nil, fmt.Errorf("error getting location: %w", err)
	}
	return byteData, nil
}

func (c *Client) GetPokemon(location *string) ([]byte, error) {
	fullUrl := c.baseURL + "/pokemon/"
	if location != nil {
		if strings.Contains(*location, fullUrl) {
			fullUrl = *location
		} else {
			fullUrl += *location
		}

	}

	byteData, err := c.getData(fullUrl)
	if err != nil {
		return nil, fmt.Errorf("error getting location: %w", err)
	}
	return byteData, nil
}

package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	httpClient http.Client
	baseURL    string
}

func NewClient() *Client {
	return &Client{
		httpClient: http.Client{},
		baseURL:    "https://pokeapi.co/api/v2",
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

func (c *Client) GetLocationAreas(pageURL *string) (PaginationResponse, error) {

	fullUrl := ""
	//if no url passed, use the first location-area
	if pageURL != nil {
		fullUrl = *pageURL
	} else {
		fullUrl = c.baseURL + "/location-area"
	}
	var response PaginationResponse
	//make request and get byte data
	res, err := c.makeRequest(fullUrl)
	if err != nil {
		return response, err
	}
	//unmarshal byte data into struct for pagination requests
	if err := json.Unmarshal(res, &response); err != nil {
		return response, err
	}
	return response, nil

}

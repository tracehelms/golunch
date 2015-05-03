package yelp

import (
	"encoding/json"
)

type Business struct {
	Name     string   `json:"name"`
	Rating   float32  `json:"rating"`
	Url      string   `json:"url"`
	Location Location `json:"location"`
}

type Location struct {
	Address []string `json:"display_address"`
}

type SearchResult struct {
	Businesses []Business `json:"businesses"`
}

func (client *Client) Search(query string, location string) SearchResult {
	params := map[string]string{
		"term":     query,
		"location": location,
		"limit":    "20",
	}

	resp := MakeRequest(client, params)

	defer resp.Body.Close()
	var r SearchResult

	err := json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		panic(err)
	}

	return r
}

package yelp

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
)

type Restaurant struct {
	name    string `json: "name"`
	rating  uint32 `json: "rating"`
	url     string `json: "url"`
	address string `json: "display_address"`
}

type Restaurants struct {
	businesses []Restaurant `json: "businesses"`
}

func (client *Client) Search(query string, location string, limit int) Restaurants {
	params := map[string]string{
		"term":     query,
		"location": location,
		"limit":    strconv.Itoa(limit),
	}

	resp := MakeRequest(client, params)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var r Restaurants
	err = json.Unmarshal(body, &r)
	if err != nil {
		panic(err)
	}

	return r
}

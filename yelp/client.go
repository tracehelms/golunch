// credit for authentication goes to https://github.com/JustinBeckwith/go-yelp
package yelp

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/mrjones/oauth"
)

const url string = "http://api.yelp.com/v2/search"

type Credentials struct {
	ConsumerKey       string // Consumer Key from the yelp API access site.
	ConsumerSecret    string // Consumer Secret from the yelp API access site.
	AccessToken       string // Token from the yelp API access site.
	AccessTokenSecret string // Token Secret from the yelp API access site.
}

type Client struct {
	Options *Credentials
	Client  *http.Client
}

func MakeRequest(client *Client, params map[string]string) *http.Response {
	c := oauth.NewCustomHttpClientConsumer(
		client.Options.ConsumerKey,
		client.Options.ConsumerSecret,
		oauth.ServiceProvider{
			RequestTokenUrl:   "",
			AuthorizeTokenUrl: "",
			AccessTokenUrl:    "",
		},
		client.Client,
	)

	token := &oauth.AccessToken{
		client.Options.AccessToken,
		client.Options.AccessTokenSecret,
		make(map[string]string),
	}

	resp, err := c.Get(url, params, token)
	if err != nil {
		panic(err)
	}
	return resp
}

func New() *Client {
	var creds Credentials
	data, err := ioutil.ReadFile("../keys.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &creds)
	if err != nil {
		panic(err)
	}
	return &Client{
		Client:  http.DefaultClient,
		Options: &creds, // comma, wtf?!
	}
}

package moonshot

import (
	"net/http"
)

type Client struct {
	client *http.Client

	apiKey string
}

func New(apiKey string) *Client {

	if apiKey == "" {
		panic("moonshot: apiKey is required")
	}

	return &Client{
		client: http.DefaultClient,
		apiKey: apiKey,
	}
}

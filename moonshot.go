package moonshot

import (
	"net/http"
)

const (
	moonshotBaseURL = "https://api.moonshot.cn"

	ModelV18K  = "moonshot-v1-8k"
	ModelV32K  = "moonshot-v1-32k"
	ModelV128K = "moonshot-v1-128k"
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

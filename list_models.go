package moonshot

import (
	"context"
	"encoding/json"
	"net/http"
)

type ModelDetail struct {
	Created int64  `json:"created,omitempty"`
	ModelID string `json:"id,omitempty"`
	Object  string `json:"object,omitempty"`
	OwnedBy string `json:"owned_by,omitempty"`
	Root    string `json:"root,omitempty"`
	Parent  string `json:"parent,omitempty"`
}

// ResponseListModels is the response from the `/v1/models`.
type ResponseListModels struct {
	ModelDetails []ModelDetail `json:"data"`
}

func (c *Client) ListModels(ctx context.Context) (*ResponseListModels, error) {

	httpReq, err := c.newHTTPRequest(ctx)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result ResponseListModels
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) newHTTPRequest(ctx context.Context) (*http.Request, error) {

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, moonshotBaseURL+apiListModels, nil)
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+c.apiKey)

	return httpReq, nil
}
